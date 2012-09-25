package trie

// The double array structure looks like this:
//
//          BASE                CHECK
// -+-------------------+-------------------
// 0| <unused>          | <unused>
// 1| free list pointer | free list pointer
// 2| 3                 | 0
// 3| ...               | ...
//
// There are no free slots initially, so the free list pointers are initialized to -1 (themselves). BASE(2) is
// the root node, and initially points to the next slot of the double array pool (i.e., beyond the end of the
// initially allocated double array).
const (
	daFreeListIndex = 1
	daRootIndex     = 2
)

// Translate a raw input byte to an index into the double array. The index cannot be 0. We're not using alpha
// maps as libdatrie does (for now), so the quick solution is to map 0 -> 256 and 1-255 map to themselves.
func byteToDAIndex(b byte) int {
	if b == '\x00' {
		return 256
	}
	return int(b)
}

type doubleArrayCell struct {
	base  int32
	check int32
}

type doubleArray struct {
	cells []doubleArrayCell
}

func newDoubleArray() *doubleArray {
	cells := []doubleArrayCell{
		doubleArrayCell{0, 0},
		doubleArrayCell{-1, -1},
		doubleArrayCell{daRootIndex + 1, 0},
	}
	return &doubleArray{cells}
}

// Get the value of BASE at some index i.
func (da *doubleArray) base(i int) int32 { return da.cells[i].base }

// Get the value of CHECK at some index i.
func (da *doubleArray) check(i int) int32 { return da.cells[i].check }

// Set the previous free pointer (-BASE) of a cell, so that the previous of s points to t.
func (da *doubleArray) setPrevious(s, t int) { da.cells[s].base = int32(-t) }

// Set the next free pointer (-CHECK) of a cell, so that the next of s points to t.
func (da *doubleArray) setNext(s, t int) { da.cells[s].check = int32(-t) }

// Remove a cell from the free list. Doesn't verify that the cell is free first, and doesn't modify the
// contents of s.
func (da *doubleArray) removeFromFreeList(s int) {
	nextFreeCell, previousFreeCell := int(-da.check(s)), int(-da.base(s))
	da.setNext(previousFreeCell, nextFreeCell)
	da.setPrevious(nextFreeCell, previousFreeCell)
}

// Free all the cells in the range [start, end). Returns ok if successful.
func (da *doubleArray) free(start, end int) bool {
	if start <= daRootIndex || end > len(da.cells) {
		return false
	}
	current := daFreeListIndex
	// Walk the free list to find the nearest preceding free node.
	var next int
	for {
		next = int(-da.check(current))
		if next >= start || next == daFreeListIndex {
			break
		}
		current = next
	}
	// Previous will be the preceding node; next will be the next one; we'll walk current up the range and set
	// the pointers as we go.
	previous := current
	for current = start; current < end; current++ {
		if n := da.check(current); n < 0 {
			// Node is already free.
			next = int(-n)
			// Previous already point to the current node, because this loop keeps the free list in a consistent
			// state.
			previous = current
		} else {
			// Node is currently occupied.
			da.setNext(previous, current)
			da.setPrevious(current, previous)
			da.setNext(current, next)
			da.setPrevious(next, current)
			previous = current
		}
	}
}

// Resize the double array to accomodate up to maxIndex (so the len(cells) == maxIndex + 1) and free the newly
// allocated cells. We could use append to take care of this but it will be less efficient to construct a new
// intermediate slice. Returns ok == true unless there was a problem.
func (da *doubleArray) resize(newSize int) bool {
	currentSize := len(da.cells)
	if newSize <= currentSize {
		return true
	}
	if newSize > cap(da.cells) {
		// Allocate a new slice of double the desired size, for room to grow.
		newDoubleArrayCells := make([]doubleArrayCell, newSize * 2)
		copy(newDoubleArray, da.cells)
		da.cells = newDoubleArrayCells
	}
	da.cells = da.cells[0:newSize]

	// Free the new cells.
	previous = int(-da.base(daFreeListIndex))
	var current int
	for current = currentSize; current < newSize; current++ {
		da.setNext(previous, current)
		da.setPrevious(current, previous)
	}
	// Fix the loop-around
	da.setNext(current, daRootIndex)
	da.setPrevious(daRootIndex, current)
	return true
}

// A helper to walk to along a trie edge inside the double array. Returns the next index along the edge ch. ch
// must be converted using byteToDAIndex first; 0 is reserved for an end-of-string character here. If the DA
// cell at s points to the tail, then inTail will be true and next is the tailBlockList index. If the walk is
// not possible (either s is invalid or there is no such child) then ok is false. next and inTail should be
// ignored unless ok == true.
func (da *doubleArray) walk(s int, ch int) (next int, inTail, ok bool) {
	if s >= len(da.cells) || s < daRootIndex {
		return 0, false, false
	}
	sContents := da.base(s)
	if sContents < 0 {
		// A tail pointer
		return int(-sContents), true, true
	}
	next = int(sContents) + ch
	if next >= len(da.cells) || next < daRootIndex {
		return 0, false, false
	}
	if int(da.check(next)) == s {
		return next, false, true
	}
	return 0, false, false
}

// Adds a new transition from state s along an edge ch, possibly relocating an existing base or expanding the
// double array. Returns ok == true unless there was some problem and it failed.
func (da *doubleArray) addBase(s int, ch int) bool {
	base := da.check(s)

	if base < 0 {
		// Cell is free. Use the first free cell as the target.
		target := int(-da.check(daFreeListIndex))
		if target == daFreeListIndex {
			// No free cells.
			if ok := da.resize(len(da.cells) +1); !ok {
				panic("Error resizing double array.")
			}
			target = int(-da.check(daFreeListIndex))
		}
		// Remove the cell from the free list.
		da.removeFromFreeList(target)
		// Use this cell as the base for our new edge.
		da.cells[target].check = int32(s)
		da.cells[target].base = 0 // TODO: Is this correct?
		// Now call ourselves recursively. The second time around we won't hit this if statement (we'll use the
		// base we just created).
		return da.addBase(s, ch)
	}

	// Cell is already in use.
	t := int(-base) + ch
	if da.check(t) == s {
		// The transition already exists.
		return true
	}
	if t >= len(da.cells) {
		// The transition overflows the double array. Add more cells.
		if ok := da.resize(t+1); !ok {
			panic("Error resizing double array.")
		}
	} else {
		// The transition conflicts with an existing transition. Relocate that base out of the way.
		panic("Base relocation unimplemented.")
	}

	// t is now free and ready to be used. Remove it from the free list and update its contents.
	da.removeFromFreeList(t)
	da.cells[t].check = s
	da.cells[t].base = 0 // TODO: Correct?
	// WIP
}

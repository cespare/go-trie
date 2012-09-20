package trie

// The double array structure looks like this:
//
//          BASE                CHECK
// -+-------------------+-------------------
// 0| free list pointer | free list pointer
// 1| 2                 | 0
// 2| ...               | ...
//
// There are no free slots initially, so the free list pointers are initialized to -1. BASE(1) is the root
// node, and initially points to the next slot of the double array pool (i.e., beyond the end of the initially
// allocated double array).
const daRootIndex = 1

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
	cells := []doubleArrayCell{doubleArrayCell{-1, -1}, doubleArrayCell{daRootIndex + 1, 0}}
	return &doubleArray{cells}
}

// Get the value of BASE at some index i.
func (da *doubleArray) base(i int) int32 { return da.cells[i].base }

// Get the value of CHECK at some index i.
func (da *doubleArray) check(i int) int32 { return da.cells[i].check }

// Free all the cells in the range [start, end). Returns ok if successful.
func (da *doubleArray) free(start, end int) bool {
	if start <= daRootIndex || end > len(da.cells) {
		return false
	}
	previous := -1
	// Walk the free list to find the nearest preceding free node.
	// WIP
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
// double array. Returns ok unless there was some problem and it failed.
func (da *doubleArray) addBase(s int, ch int) bool {
	target := da.check(s)
	if target >= len(da.cells) {
		// WIP
	}
}

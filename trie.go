// The trie package implements a trie (prefix tree) data structure over byte slices.
package trie

import (
	"fmt"
)

// The core trie data structure.
type Trie struct {
	da *doubleArray
	tail  *tailBlockList
}

// Construct a new, empty trie.
func New() *Trie {
	return &Trie{newDoubleArray(), newTailBlockList()}
}

// Return the Node at the root of the trie.
func (t *Trie) Root() *Node {
	return newNode(t)
}

// Add a []byte to a trie. The return value indicates whether s was added to the trie -- it is true if s was
// not already present; false otherwise.
func (t *Trie) Add(s []byte) bool {
	// Start at the root
	current := newNode(t)
	endOfString := true
	var i int
	var ch byte
	for i, ch = range s {
		if !current.Walk(ch) {
			endOfString = false
			break
		}
	}

	if endOfString {
		// We walked existing trie nodes all the way to the end of s.
		if current.Terminal() {
			// s already exists in t.
			return false
		}
		// s is a prefix of another element of t.
		if current.inTail {
			// Need to move common nodes into the double array.
			panic("tail splitting unimplemented.")
		} else {
			// Just need to add a trivial tail ending.
			panic("empty tail insertion unimplemented.")
		}
		panic("unreached")
	}

	// We reached the end of existing trie nodes before traversing s completely.
	if current.inTail {
		// Need to move the current tail entirely into the double array and put the remainder of s in a new tail.
		panic("tail splitting (2) unimplemented.")
	} else {
		// Need to insert a new double array node, relocating other bases as necessary.
		fmt.Printf("\033[01;34m>>>> i: %v\x1B[m\n", i)
		t.da.addBase(current.s, byteToDAIndex(ch))
		panic("blah")
	}
	panic("unreached")
}

// Remove a []byte from the trie. Returns true if the []byte was removed; returns false if the []byte is not
// in the trie.
func (t *Trie) Delete(s []byte) bool {
	return false
}

// Test whether a []byte is present in the trie.
func (t *Trie) Contains(s []byte) bool {
	current := t.Root()
	for _, b := range s {
		if !current.Walk(b) {
			return false
		}
	}
	return current.Terminal()
}

// Return all keys in the trie beginning with a certain prefix.
func (t *Trie) ChildrenWithPrefix(prefix []byte) [][]byte {
	return [][]byte{}
}

// Get every key in the trie.
func (t *Trie) Keys() [][]byte {
	return [][]byte{}
}

// Print debugging info.
func (t *Trie) Print() {
	fmt.Println("Double array:\n")
	fmt.Printf("            BASE        CHECK\n")
	for i, cell := range t.da.cells {
		fmt.Printf("%8d [%8d]  [%8d]", i, cell.base, cell.check)
		switch i {
		case 0:
			fmt.Printf(" (free list pointers)\n")
		case 1:
			fmt.Printf(" (root)\n")
		default:
			fmt.Printf("\n")
		}
	}

	fmt.Println("\nTail:\n")
	fmt.Printf("firstFreeIndex -> %d\n", t.tail.firstFreeIndex)
	nextFree := t.tail.firstFreeIndex
	for i, tb := range t.tail.tailBlocks {
		fmt.Printf("%8d ", i)
		if nextFree == i {
			fmt.Printf("(free) -> %d\n", i, tb.nextFreeIndex)
			nextFree = tb.nextFreeIndex
		} else {
			fmt.Printf("[ ")
			for _, b := range tb.tail {
				fmt.Printf("%c ", b)
			}
			fmt.Printf("]\n")
		}
	}
}

// The trie package implements a trie (prefix tree) data structure over byte slices.
package trie

// The core trie data structure.
type Trie struct {
	da *doubleArray
	tail  *tailBlockList
}

// Construct a new, empty trie.
func New() *Trie {
	return &Trie{new(doubleArray), new(tailBlockList)}
}

// Return the Node at the root of the trie.
func (t *Trie) Root() *Node {
	return newNode(t)
}

// Add a []byte to a trie. The return value indicates whether s was added to the trie-- it is true if s was
// not already present; false otherwise.
func (t *Trie) Add(s []byte) bool {
	// Start at the root
	current := newNode(t)
	endOfString := true
	for _, ch := range s {
		if !current.Walk(ch) {
			endOfString = false
			break
		}
	}

	if endOfString {
		// We walked existing trie nodes all the way to the end of s.
		if n.Terminal() {
			// s already exists in t.
			return false
		}
		// s is a prefix of another element of t.
		if n.inTail {
			// Need to move common nodes into the double array.
			panic("tail splitting unimplemented.")
		} else {
			// Just need to add a trivial tail ending.
			panic("\0 tail insertion unimplemented.")
		}
		panic("unreached")
	}

	// We reached the end of existing trie nodes before traversing s completely.
	if inTail {
		// Need to move the current tail entirely into the double array and put the remainder of s in a new tail.
		panic("tail splitting (2) unimplemented.")
	} else {
		// 
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
	for _, b := s {
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

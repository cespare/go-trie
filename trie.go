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

// Add a []byte to a trie. The return value indicates whether the []byte was already in the trie.
func (t *Trie) Add(s []byte) bool {
	return false
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

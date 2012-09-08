// The trie package implements a trie (prefix tree) data structure over strings.
package trie

// The core trie data structure.
type Trie struct {
	/*da *DoubleArray*/
	/*tail *Tail*/
}

// Construct a new, empty trie.
func New() *Trie {
	return &Trie{}
}

// Return the Node at the root of the trie.
func (t *Trie) Root() *Node {
	return nil
}

// Add a string to a trie. The return value indicates whether the string was already in the trie.
func (t *Trie) Add(s string) bool {
	return false
}

// Remove a string from the trie. Returns true if the string was removed; returns false if the string is not
// in the trie.
func (t *Trie) Delete(s string) bool {
	return false
}

// Test whether a string is present in the trie.
func (t *Trie) Contains(s string) bool {
	return false
}

// Return all keys in the trie beginning with a certain prefix.
func (t *Trie) ChildrenWithPrefix(prefix string) []string {
	return []string{}
}

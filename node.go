package trie

// The Node abstraction is provided to allow for simple walking and inspection of the trie. It represents a
// single vertex of the logical trie structure.
type Node struct {
	/*trie *Trie*/
}

// Check whether this node is a leaf of the trie (i.e., there are no children).
func (n *Node) Leaf() bool {
	return false
}

// Check whether this is a terminal node of some string in the trie.
func (n *Node) Terminal() bool {
	return false
}

// The current value at this Node. ok indicates whether there is such a value (the root node of the trie does
// not correspond to any state transition).
func (n *Node) State() (ch byte, ok bool) {
	return byte(0), false
}

// The partial array from the root to the current node. (This is empty at the root node of a trie.)
func (n *Node) PartialValue() []byte {
	return new([]byte)
}

// Return the current value of a complete byte array. If the node is terminal, then this is the full
// corresponding array, otherwise ok will be false and value should be ignored.
func (n *Node) Value() (value string, ok bool) {
	if n.Terminal() {
		return n.PartialValue(), true
	}
	return "", false
}

// Walk down the trie. Returns a new node reached by following the ch child. If there is no such child, this
// returns nil.
func (n *Node) walk(ch rune) *Node {
	return nil
}

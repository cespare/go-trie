package trie

// A single node in the (logical) trie. Note that nodes may be collapsed in the underlying implementation.
type Node struct {
	trie *Trie
}

// Check whether this node is a leaf of the trie (i.e., there are no children).
func (n *Node) Leaf() bool {
	return false
}

// Check whether this is a terminal node of some string in the trie.
func (n *Node) Terminal() bool {
	return false
}

// The current character at this Node. ok indicates whether there is such a character (the root node of the
// trie does not correspond to any state transition).
func (n *Node) State() (ch rune, ok bool) {
	return rune(0), false
}

// The full partial string from the root to the current node. (This is empty at the root node of a trie.)
func (n *Node) PartialValue() string {
	return ""
}

// Return the current value of a complete string. If the node is terminal, then this is the full corresponding
// string, otherwise ok will be false and the string return value should be ignored.
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

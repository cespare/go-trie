package trie

// The Node abstraction is provided to allow for simple walking and inspection of the trie. It represents a
// single vertex of the logical trie structure.
type Node struct {
	da   *doubleArray
	tail *tailBlockList
	// Whether the current node state is in the DA or the tail
	inTail bool
	// The current tailBlock that this state represents; only valid if inTail == true
	tb tailBlock
	// s is an index value that represents the current state. If inTail == true, then s is an index into
	// tb.tail; otherwise, s is an index into BASE of trie.da.
	s int
}

/*
Private methods
*/

// Construct a new Node at the root of the trie.
func newNode(t *Trie) *Node {
	n := &Node{t.da, t.tail}
	n.inTail = false
	n.s = daRootIndex
	return n
}

// Internal version of Walk that accepts a raw transition index (rather than a byte) and doesn't apply any
// conversions on it. This allows for walking along '0' -- i.e., testing for end-of-string.
func (n *Node) walk(c int) bool {
	if n.inTail {
		if n.s >= len(n.tb.tail) {
			// We're at the end of the tail.
			return false
		}
		if n.tb.tail[n.s+1] == c {
			// We have a match.
			n.s++
			return true
		}
		return false
	}

	// We're in the double array
	if n.s >= len(n.da.cells) {
		// Index outside of DA bounds.
		return false
	}

	next, inTail, ok := n.da.walk(n.s, c)
	if !ok {
		return false
	}

	if inTail {
		n.inTail = true
		n.s = 0
		n.tb = n.tail.tailBlocks[next]
		return true
	}

	n.s = next
	return true
}

/*
Public methods
*/

// Walk down the trie along some edge ch. If there is no such child, the return value is false and the state
// of the Node is unchanged. Otherwise, the return value is true.
func (n *Node) Walk(ch byte) bool {
	n.walk(byteToDAIndex(ch))
}

// Check whether this node is a leaf of the trie (i.e., there are no children).
func (n *Node) Leaf() bool {
	if !n.Terminal() {
		// Any non-terminal node cannot be a leaf.
		return false
	}
	if n.inTail {
		// If the node is terminal and it's in the tail already, it cannot have children.
		return true
	}

	// We know that n is terminal and currently in the double array. Now we must check that it has no children
	// other than '\0'. Scan CHECK for any values that equal n.s.
	// TODO: Double-check this logic.
	// TODO: I think that we don't need to check CHECK(1), but I need to double-check this.
	for i := daRootIndex+1; i < len(n.da.cells); i++ {
		if n.da.check(i) == n.s {
			return false
		}
	}
	return true
}

// Check whether this is a terminal node of some string in the trie.
func (n *Node) Terminal() bool {
	if n.inTail {
		return n.tb.terminal(n.s)
	}
	// We're still in the double array. Walk down the \0 branch and see if it ends; either way, restore the node
	// state afterwards.
	oldS = n.s
	defer func() {
		n.inTail = false
		n.tb = nil
		n.s = oldS
	}()

	next, inTail, ok := n.walk(n.s, 0)
	if !ok {
		return false
	}
	if inTail {
		return len(n.tb.tail) == 0
	}
	// All keys end in the tail
	return false
}

// Copy a node's state to a new node.
func (n *Node) Copy() *Node {
	newNode := new(Node)
	*newNode = *n
	return newNode
}

// TODO: I removed the below APIs because the user can easily track the state herself and it doesn't
// make sense to force everyone to have the overhead of doing it. Need to decide whether that was the right
// way to go.

// The current value at this Node. ok indicates whether there is such a value (the root node of the trie does
// not correspond to any state transition).
/*func (n *Node) State() (ch byte, ok bool) {*/
	/*return byte(0), false*/
/*}*/

/*// The partial array from the root to the current node. (This is empty at the root node of a trie.)*/
/*func (n *Node) PartialValue() []byte {*/
	/*return new([]byte)*/
/*}*/

/*// Return the current value of a complete byte array. If the node is terminal, then this is the full*/
/*// corresponding array, otherwise ok will be false and value should be ignored.*/
/*func (n *Node) Value() (value string, ok bool) {*/
	/*if n.Terminal() {*/
		/*return n.PartialValue(), true*/
	/*}*/
	/*return "", false*/
/*}*/

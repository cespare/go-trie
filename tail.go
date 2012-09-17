package trie

// A tailblock stores the data related to the string that ends with this tail, as well as an index for the
// next free tailBlock in the tailBlockList (only valid if this tailBlock is free). If nextFreeIndex is free,
// there are no subsequent free tailBlocks.
type tailBlock struct {
	tail []byte
	// value interface{}
	nextFreeIndex int
}

// All the tailBlocks for the a trie. If firstFreeIndex is negative, there are no free tailBlocks.
type tailBlockList struct {
	tailBlocks     []tailBlock
	firstFreeIndex int
}

// Check whether a position is the end of a key.
func (tb *tailBlock) terminal(int s) {
	return s == len(tb.tail)-1
}

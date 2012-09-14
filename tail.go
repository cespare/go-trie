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

// Check whether a position is the end of a key. A non-empty tail ends with a terminating \0 byte.
func (tb *tailBlock) terminal(int s) {
	if len(tb.tail) == 0 {
		return s == 0
	}
	// Not strictly necessary to verify that tail[s+1] == 0, but a good sanity check.
	return s+2 == len(tb.tail) && tb.tail[s+1] == '\0'
}

// The trie package implements a trie (prefix tree) data structure over byte slices.
package trie






type doubleArrayCell struct {
	base int32
	check int32
}

type doubleArray struct {
	cells []doubleArrayCell
}

func newDoubleArray() *doubleArray {
	return &doubleArray{}
}

// Get the value of BASE at some index i.
func (da *doubleArray) base(i int32) int32, bool { return da.cells[i].base }

// Get the value of CHECK at some index i.
func (da *doubleArray) check(i int32) int32 { return da.cells[i].check }



// A tailblock stores the data related to the string that ends with this tail, as well as an index for the
// next free tailBlock in the tailBlockList (only valid if this tailBlock is free). If nextFreeIndex is free,
// there are no subsequent free tailBlocks.
type tailBlock struct {
	tail []byte
	// value interface{}
	nextFreeIndex int32
}

// All the tailBlocks for the a trie. If firstFreeIndex is negative, there are no free tailBlocks.
type tailBlockList struct {
	tailBlocks []tailBlock
	firstFreeIndex int32
}







// The core trie data structure.
type Trie struct {
	da *doubleArray
	t *tail
}




// Construct a new, empty trie.
func New() *Trie {
	return &Trie{new(doubleArray), new(tailBlockList)}
}

// Return the Node at the root of the trie.
func (t *Trie) Root() *Node {
	return nil
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
	return false
}

// Return all keys in the trie beginning with a certain prefix.
func (t *Trie) ChildrenWithPrefix(prefix []byte) [][]byte {
	return [][]byte{}
}

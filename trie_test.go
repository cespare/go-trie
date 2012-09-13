package trie

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TrieTestSuite struct{}

var _ = Suite(&TrieTestSuite{})

func (s *TrieTestSuite) TestAddingWordsToTrie(c *C) {
	t := New()
	c.Check(t.Add("bar"), Equals, true)
	/*c.Check(t.Add("barter"), Equals, true)*/
	/*c.Check(t.Add("barter"), Equals, false)*/
	/*c.Check(t.Contains("bar"), Equals, true)*/
	/*c.Check(t.Contains("barter"), Equals, true)*/
	/*c.Check(t.Contains("barn"), Equals, false)*/
}

// Corner cases:
// * Random UTF-8 strings
// * UTF-8 strings with the NUL (0) code point in them (*not handled by libdatrie)
// * Should be able to walk based on a rune (int)

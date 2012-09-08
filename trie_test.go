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

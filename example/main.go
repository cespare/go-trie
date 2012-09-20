package main

import (
	"github.com/cespare/go-trie"
)

func main() {
	t := trie.New()
	t.Add([]byte("hello world!"))
	t.Print()
}

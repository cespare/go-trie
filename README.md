# go-trie (WORK IN PROGRESS)

A pure-Go implementation of a Trie data structure for strings.

## Installation

    $ go get github.com/cespare/go-trie

## Usage

Very simple example usage.

``` go
import (
  "github.com/cespare/go-trie"
)

t = trie.New()
t.Add("Hello!")
t.Contains("Hello!")   // => true
t.Contains("Goodbye!") // => false
```

Please see the API documentation for more information and more advanced usage.

## Author

Caleb Spare ([cespare](https://github.com/cespare))

## To Do

* Finish a working implementation!
* The trie.Retrieve method(s) don't need to do a node-by-node walk -- they can do a more efficient search by
	doing a direct byte-by-byte comparison as soon as we get to the tail.
* Reading/writing from disk

## License

MIT Licensed.

# Implementation Notes

This is an implementation of the trie data structure using two arrays ("double array trie") described in [this
paper](http://sc.snu.ac.kr/~xuan/spe777ja.pdf). It is based on the description and C implementation given
[here](http://linux.thai.net/~thep/datrie/datrie.html). The code may be found linked on that site, or
alternatively in the repo for the Ruby wrapper ('fast_trie') [here](https://github.com/tyler/trie). The API is
partially copied from that (very nice) Ruby API.

This particular trie implementation method is highly memory-efficient and is optimized for retrieval speed.
Writes can be slow, especially when there are many keys in the trie. These tradeoffs make this library
well-suited for applications with static tries that are generated once and then used many times.

go-trie allows for storing an arbitrary object as a value for each key (in the form of an `interface{}`).
Alternatively, it can be used as a set. (Use `trie.NewTrieMap` for the former variant and `trie.NewTrieSet`
for the latter.)

The keys of a go-trie are `[]byte`, rather than strings. There are several reasons for this -- one is that
most performance-critical applications will have to process `[]byte`s anyway. However, go-trie provides
convenience methods that accept `string`s and `rune`s instead of `[]byte`s and `byte`s (e.g.,
`TrieMap.Contains(key byte[]) bool` vs. `TrieMap.ContainsString(key string) bool`). These methods all just
convert to `[]byte` or `byte` under the hood, so they are slightly less efficient than their `byte`-based
counterparts.

The notes that follow are to provide a quick run-down of the DA trie implementation details along with the
particular go-trie-specific adjustments.

## Input alphabet

Later versions of libdatrie include an "alphabet map" (see `alpha-map.c`). This is used to map an arbitrary
input alphabet to a restricted continuous alphabet. However, it requires the user to provide the map before
using the structure. For ease of use, I have opted to leave this feature out.

As far as I can tell, libdatrie only handles C-style (null-terminated) strings as input, and furthermore uses
`'\0'` as the string termination character internally as well. This means that it cannot be used with
arbitrary UTF-8 that may contain the 0 (NUL) code point (U+0000). go-trie does not have this limitation, and
allows for null bytes to be present in the keys. Because 0 does not work as an offset in the double array,
input bytes are mapped using the `byteToDAIndex` function, which maps 0 to 256, before using them as offsets.

## Internal structure

A DA trie consists of two distinct structures: the double array and the tail. The tail stores the unambiguous
suffixes of keys as an optimization. Traversing the trie along the edges defined by a key involves traversing
the double array until reaching an edge that points to the tail, and then traversing that tail to the end.

### Double array

The double array is represented by the `doubleArray` struct type in `double_array.go`. It consists of a slice
of `doubleArrayCell`s, which are simply pairs of `int32`s. This represents what is (in the original
description and the C implementation) two separate, equally sized, expanding arrays. The names of the arrays
are BASE and CHECK.

A particular position in the double array (that is, a `(base, check)` pair of `int32`s) may be free or
occupied. If they are occupied, then `check` is positive and is an index to another valid position in the
double array. `base` may be positive, in which case it is also an index to another valid position in the
double array, or it may be negative, in which case it is (the negative of) a valid index into the `tail`.

If a position in the double array is free, then the values of `base` and `check` are both negative, and they
encode (via their negations) a circular doubly-linked list to the other free positions in the double array.
The value of `check` is the negative index to the next free position; the value of `base` is the negative
index to the previous free position. The first position (index 0) of the double array are permanently reserved
for free list pointers to indicate the first and last free indices in the double array.

The various procedures for inserting, retrieving, and deleting keys are described in the linked materials.

Although keys in go-trie are not null-terminated in general, there needs to be a way of indicating terminal
keys inside the double array. Therefore a null byte is used for this purpose in the double array (but not the
tail). For example, suppose that the only two keys in the trie are `artist` and `art`. Then the double array
would have entries for `a -> r -> t`. The two transitions from the `t` node would be reprented by `i` and `\0`
edges from the index indicated by `base[t]`. The tail entry for `\0` would be have a zero-length `[]byte`
entry, while the tail entry for `i` would have the byte array `st`.

### Tail

The tail is represented as an array of `tailBlock` structs. Each of these contains a `[]byte` (the suffix) and
a data value in an `interface{}`. It also has an index to the next free `tailBlock`, only valid if the block
is free. The entire tail is represented by a `tailBlockList`, which consists of a slice of `tailBlock`s and a
pointer to the first free `tailBlock`.

## Node abstraction

The `Node` struct provides an abstraction for traversing the logical trie represented by the double array
structure. This is useful both internally, as a way of organizing state while traversing the trie, and
externally as an interface provided to the user. A `Node` has pointers to the double array and tail, an index
value, and a boolean value (`inTail`) to indicate whether the current state indicated by the index is in the
double array or the tail. The internal operations on the trie are implemented as private methods on `Node`.

`Node` is useful in the public interface because there are many situations using a trie that require walking
it concurrently with other operations. For instance, if performing a search over a space of bytes and
attempting to locate words in some dictionary, one might traverse edges of the trie as they are explored and
thus be able to terminate any search as soon as the current string is no longer the prefix of any key.

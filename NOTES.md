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
convenience methods that accept `string`s and `rune`s instead of `[]byte`s and `byte`s (e.g.
`TrieMap.Contains(key byte[]) bool` vs. `TrieMap.ContainsString(key string) bool`). These methods all just
convert to `[]byte` or `byte` under the hood, so they are slightly less efficient than their `byte`-based
counterparts.

The notes that follow are to provide a quick run-down of the DA trie implementation details along with the
particular go-trie-specific adjustments.

## Input alphabet

Later versions of libdatrie include an "alphabet map" (see `alpha_map.c`). This is used to map an
arbitrary input alphabet to a restricted continuous alphabet. However, it requires the user to provide the map
before using the structure. For ease of use, I have opted to leave this feature out.

As far as I can tell, libdatrie only handles C-style (null-terminated) strings as input, and furthermore uses
`'\0'` as the string termination character internally as well. This means that it cannot be used with
arbitrary UTF-8 that may contain the 0 (NUL) code point (U+00000000). go-trie does not have this limitation,
and allows for null bytes to be present in the keys. Because 0 does not work as an offset in the double array,
input bytes are mapped using the `byteToDAIndex` function, which maps 0 to 256, before using them as offsets.

## Internal structure

## Node abstraction

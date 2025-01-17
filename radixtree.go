// package radixtree implements a generic mutable radix tree
package radixtree

import (
	"fmt"
	"strings"
)

type node struct {
	value interface{} // value of the node or nil if there is no value
	next  []*node     // array of pointers to the next node
}

const R = 256 // extended ASCII

// Creates a new node structure, initializing the next array
func createNode() *node {
	n := node{next: make([]*node, R)}
	return &n
}

type RadixTree struct {
	root *node // root of trie
	n    int // number of keys in trie

}

// Returns the value associated with the given key if the radix tree contains the key or nil.
func (r *RadixTree) Get(key string) interface{} {
	x := get(r.root, key, 0)
	return x
}

// Returns a boolean indicating if the radix tree contains the given key.
func (r *RadixTree) Contains(key string) bool {
	return r.Get(key) != nil
}

func get(x *node, key string, d int) *node {
	if d == len(key) {
		return x
	}
	c := key[d]
	return get(x.next[c], key, d+1)
}

// Adds the given key and value to the radix tree, overwriting the old value with the new value if the radix tree already contains the key.
func (r *RadixTree) Put(key string, value interface{}) {
	if value == nil {
		r.Delete(key)
	}
	r.root = r.put(r.root, key, value, 0)
}

func (r *RadixTree) put(x *node, key string, value interface{}, d int) *node {
	if x == nil {
		x = createNode()
	}
	if d == len(key) {
		if x.value == nil {
			r.n++
		}
		x.value = value
		return x
	}
	c := key[d]
	x.next[c] = r.put(x.next[c], key, value, d+1)
	return x
}

// Returns the number of key-value pairs of the radix tree.
func (r *RadixTree) Size() int {
	return r.n
}

// Returns a boolean indicating if the radix tree is empty
func (r *RadixTree) IsEmpty() bool {
	return r.Size() == 0
}

// Removes the key from the radix tree if the key is present.
func (r *RadixTree) Delete(key string) {
	r.root = r.delete(r.root, key, 0)
}

func (r *RadixTree) delete(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.value != nil {
			r.n--
		}
		x.value = nil
	} else {
		c := key[d]
		x.next[c] = r.delete(x.next[c], key, d-1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.value != nil {
		return x
	}
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

// Returns all keys of the radix tree.
func (r *RadixTree) Keys() []string {
	return r.KeysWithPrefix("")
}

// Returns all keys of the radix tree that start with the given prefix.
func (r *RadixTree) KeysWithPrefix(prefix string) []string {
	results := make([]string, 0)
	x := get(r.root, prefix, 0)
	b := []rune(prefix)
	results = collect(x, b, results)
	return results
}

func collect(x *node, prefix []rune, results []string) []string {
	if x == nil {
		return results
	}
	if x.value != nil {
		results = enqueue(results, makeString(prefix))
	}
	for c := 0; c < R; c++ {
		prefix = append(prefix, rune(c))
		results = collect(x.next[c], prefix, results)
		prefix = deleteCharAt(prefix, len(prefix)-1)
	}
	return results
}

// Returns all of the keys of the radix tree that match the given pattern,
// where . symbol is treated as wildcard character that matches any single character.
func (r *RadixTree) KeysThatMatch(pattern string) []string {
	results := make([]string, 0)

	b := make([]rune, 0)
	results = collectPattern(r.root, b, []rune(pattern), results)
	return results
}

func collectPattern(x *node, prefix []rune, pattern []rune, results []string) []string {
	if x == nil {
		return results
	}
	d := len(prefix)
	if d == len(pattern) && x.value != nil {
		results = enqueue(results, makeString(prefix))
	}
	if d == len(pattern) {
		return results
	}
	c := pattern[d]
	if c == '.' {
		for ch := 0; ch < R; ch++ {
			prefix = append(prefix, rune(ch))
			results = collectPattern(x.next[ch], prefix, pattern, results)
			prefix = deleteCharAt(prefix, len(prefix)-1)
		}
	} else {
		prefix = append(prefix, rune(c))
		results = collectPattern(x.next[c], prefix, pattern, results)
		prefix = deleteCharAt(prefix, len(prefix)-1)
	}
	return results
}

// Returns the string in the symbol table that is the longest prefix of the given query.
func (r *RadixTree) LongestPrefixOf(query string) string {
	q := []rune(query)
	length := longestPrefixOf(r.root, q, 0, -1)
	if length == -1 {
		return ""
	} else {
		return string(q[:length])
	}
}

func longestPrefixOf(x *node, query []rune, d int, length int) int {
	if x == nil {
		return length
	}
	if x.value != nil {
		length = d
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	return longestPrefixOf(x.next[c], query, d+1, length)
}

// Prints the structure of the radix tree
func (r *RadixTree) PrintStructure() {
	var b strings.Builder
	printStructure(r.root, 0, &b)
	fmt.Println(b.String())
}

func printStructure(x *node, d int, b *strings.Builder) {
	runes := make([]rune, 0)
	children := make([]*node, 0)
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			runes = append(runes, rune(c))
			children = append(children, x.next[c])
		}
	}
	l := len(runes)
	if l == 1 {
		b.WriteRune(runes[0])
		printStructure(children[0], d+1, b)
	} else if l > 1 {
		for i, r := range runes {
			b.WriteString("\n")
			b.WriteString(ws(d))
			b.WriteRune(r)
			child := children[i]
			printStructure(child, d+1, b)
		}
	}
}

func ws(count int) string {
	return strings.Repeat(" ", count)
}

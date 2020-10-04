package radixtree

import (
	"fmt"
	"strings"
)

type node struct {
	value interface{}
	next []*node
}

const R = 256

func createNode() *node {
	n := node{next: make([]*node, R)}
	return &n
}

type RadixTree struct {

	root *node
	n int
}

func (r *RadixTree) Get(key string) interface{} {
	x := r.get(r.root, key, 0)
	return x
}

func (r *RadixTree) Contains(key string) bool {
	return r.Get(key) != nil
}

func (r *RadixTree) get(x *node, key string, d int) *node {
	if d == len(key) {
		return x
	}
	c := key[d]
	return r.get(x.next[c], key, d + 1)
}

func (r *RadixTree) Put(key string, value interface{}) {
	if value == nil {
		r.Delete(key)
	}
	r.root = r.put(r.root, key, value, 0)
}

func (r *RadixTree) put(x *node, key string, value interface{}, d int) *node {
	if (x == nil) {
		x = createNode()
	}
	if d == len(key)  {
		if x.value == nil {
			r.n++
		}
		x.value = value
		return x
	}
	c := key[d]
	x.next[c] = r.put(x.next[c], key, value, d + 1)
	return x
}

func (r *RadixTree) Size() int {
	return r.n
}

func (r *RadixTree) IsEmpty() bool {
	return r.Size() == 0
}

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
		x.next[c] = r.delete(x.next[c], key, d - 1)
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

func (r *RadixTree) Keys() []string {
	return r.KeysWithPrefix("")
}

func (r *RadixTree) KeysWithPrefix(prefix string) []string {
	results := make([]string, 0)
	x := r.get(r.root, prefix, 0)
	b := []rune(prefix)
	results = r.collect(x, b, results)
	return results
}

func (r *RadixTree) collect(x* node, prefix []rune, results []string) []string {
	if x == nil {
		return results
	}
	if x.value != nil {
		results = enqueue(results, makeString(prefix))
	}
	for c := 0; c < R; c++ {
		prefix = append(prefix, rune(c))
		results = r.collect(x.next[c], prefix, results)
		prefix = deleteCharAt(prefix, len(prefix) - 1)
	}
	return results
}

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
		printStructure(children[0], d + 1, b)
	} else if l > 1 {
		for i, r := range runes {
			b.WriteString("\n")
			b.WriteString(ws(d))
			b.WriteRune(r)
			child := children[i]
			printStructure(child, d + 1, b)
		}
	}
}

func ws(count int) string {
	return strings.Repeat(" ", count)
}
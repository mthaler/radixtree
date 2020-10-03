package radixtree

import "fmt"

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

func (r *RadixTree) PrintStructure() {
	printStructure(r.root)
}

func printStructure(x *node) {
	if x.value != nil {
		fmt.Printf("value: %v", x.value)
	}
	for c := 0; c < R; c++ {
		if x.next[c] != nil {
			printStructure(x.next[c])
		}
	}
}
package radixtree

import "testing"

func TestRadixTree_PrintStructure(t *testing.T) {
	r := RadixTree{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	r.Put("rubens", 4)
	r.Put("ruber", 5)
	r.Put("rubicon", 6)
	r.Put("rubicundus", 7)
	r.PrintStructure()
}

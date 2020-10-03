package radixtree

import "testing"

func TestRadixTree_PrintStructure(t *testing.T) {
	r := RadixTree{}
	r.Put("Douglas Adams", 42)
	r.PrintStructure()
}

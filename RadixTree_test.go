package radixtree

import (
	"reflect"
	"testing"
)

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

func TestRadixTree_Keys(t *testing.T) {
	r := RadixTree{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	r.Put("rubens", 4)
	r.Put("ruber", 5)
	r.Put("rubicon", 6)
	r.Put("rubicundus", 7)
	expected := []string {"romane", "romanus", "romulus", "rubens", "ruber", "rubicon", "rubicundus"}
	if !reflect.DeepEqual(r.Keys(), expected) {
		t.Errorf("r.Keys should be %v", expected)
	}
}
package radixtree

import (
	"reflect"
	"testing"
)

func TestRadixTree_PrintStructure(t *testing.T) {
	r := createTestTree()
	r.PrintStructure()
}

func TestRadixTree_Keys(t *testing.T) {
	r := createTestTree()
	expected := []string{"romane", "romanus", "romulus", "rubens", "ruber", "rubicon", "rubicundus"}
	if !reflect.DeepEqual(r.Keys(), expected) {
		t.Errorf("r.Keys should be %v", expected)
	}
}

func TestRadixTree_KeysWithPrefix(t *testing.T) {
	r := createTestTree()
	expected := []string{"romane", "romanus", "romulus"}
	if !reflect.DeepEqual(r.KeysWithPrefix("rom"), expected) {
		t.Errorf("r.Keys should be %v", expected)
	}
}

func TestRadixTree_KeysThatMatch(t *testing.T) {
	r := createTestTree()
	expected := []string{"romanus", "romulus"}
	if !reflect.DeepEqual(r.KeysThatMatch("rom...s"), expected) {
		t.Errorf("r.Keys should be %v", expected)
	}
}

func TestRadixTree_LongestPrefixOf(t *testing.T) {
	r := createTestTree()
	if r.LongestPrefixOf("romulus1234") != "romulus" {
		t.Error("r.LongestPrefixOf(romulus1234) should be romulus")
	}
}

func createTestTree() RadixTree {
	r := RadixTree{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	r.Put("rubens", 4)
	r.Put("ruber", 5)
	r.Put("rubicon", 6)
	r.Put("rubicundus", 7)
	return r
}

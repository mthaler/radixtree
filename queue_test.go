package radixtree

import (
	"testing"
)

func Test_enqueue(t *testing.T) {
	queue := make([]string, 0)
	queue = enqueue(queue, "test")
	if len(queue) != 1 {
		t.Error("Length of queue should be 1")
	}
}

func Test_dequeue(t *testing.T) {
	queue := make([]string, 0)
	queue = enqueue(queue, "test")
	element, queue := dequeue(queue)
	if element != "test" {
		t.Error("Element should be 42")
	}
}
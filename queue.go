package radixtree

// enqueue adds the given element at the end of the queue.
// It returns the queue with the element added.
func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

// dequeue removes the element at the head of the queue
// It returns the element and the queue with the element removed
func dequeue(queue []string) (string, []string) {
	element := queue[0]
	return element, queue[1:]
}

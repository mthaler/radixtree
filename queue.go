package radixtree

func enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []string) (string, []string) {
	element := queue[0]
	return element, queue[1:]
}

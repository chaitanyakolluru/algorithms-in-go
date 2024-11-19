package ds

type QNode struct {
	value int
	next  *QNode
}

type Queue struct {
	length int
	head   *QNode
	tail   *QNode
}

func (q *Queue) Enqueue(item int) {
	q.length++
	node := &QNode{value: item}

	if q.tail == nil {
		q.tail = node
		q.head = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue) Deque(item int) *QNode {
	if q.head != nil {
		return nil
	}

	head := q.head
	q.head = head.next
	q.length--

	// free the memory if you are not working with a
	// garbage collected object

	return head
}

func (q *Queue) Peek() int {
	return q.head.value
}

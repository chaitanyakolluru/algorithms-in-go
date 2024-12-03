package ds

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type Queue[T any] struct {
	Length int
	Head   *QNode[T]
	Tail   *QNode[T]
}

func (q *Queue[T]) Enqueue(item T) {
	q.Length++
	node := &QNode[T]{value: item}

	if q.Tail == nil {
		q.Tail = node
		q.Head = node
		return
	}

	q.Tail.next = node
	q.Tail = node
}

func (q *Queue[T]) Deque(item T) *QNode[T] {
	if q.Head != nil {
		return nil
	}

	Head := q.Head
	q.Head = Head.next
	q.Length--

	// free the memory if you are not working with a
	// garbage collected object

	return Head
}

func (q *Queue[T]) Peek() T {
	return q.Head.value
}

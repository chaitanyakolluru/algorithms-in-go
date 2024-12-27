package ds

// the Heap Data Structure
//
// - the simplest way to put it is a binary tree where every child and grand
// child is smaller (MaxHeap), or larger (MinHeap) than the current node.
// - whenever a node is added, we must adjust the tree
// - whenever a node is deleted, we must adjust the tree
// - there is no traversing the tree

// min heap; top value is the smallest
// min item is the root
//         50
//   71       100
// 101    80       108

// 2i + 1  and 2i + 2 for children
// ( i -1 )/2 to go from child to parent

type MinHeap struct {
	Length int
	Data   []int
}

func (m *MinHeap) insert(value int) {
	m.Data[m.Length] = value
	m.heapifyUp(m.Length)
	m.Length++
}

func (m *MinHeap) delete() int {
	if m.Length == 0 {
		return 0
	}

	out := m.Data[0]
	m.Length--
	if m.Length == 0 {
		m.Data = []int{}
		return out
	}

	m.Data[0] = m.Data[m.Length]
	m.heapifyDown(0)

	return out
}

func (m *MinHeap) parent(idx int) int {
	return ((idx - 1) / 2)
}

func (m *MinHeap) leftChild(idx int) int {
	return (2 * idx) + 1
}

func (m *MinHeap) rightChild(idx int) int {
	return (2 * idx) + 2
}

func (m *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := m.parent(idx)
	parentV := m.Data[p]
	v := m.Data[idx]

	if parentV > v {
		m.Data[idx] = parentV
		m.Data[p] = v
		m.heapifyUp(p)
	}
}

func (m *MinHeap) heapifyDown(idx int) {
	lIdx := m.leftChild(idx)
	rIdx := m.rightChild(idx)

	if idx >= m.Length || lIdx >= m.Length {
		return
	}

	lV := m.Data[lIdx]
	rV := m.Data[rIdx]
	v := m.Data[idx]

	if lV > rV && v > rV {
		m.Data[idx] = rV
		m.Data[rIdx] = v
		m.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		m.Data[idx] = lV
		m.Data[lIdx] = v
		m.heapifyDown(lIdx)
	}
}

package ds

import "math"

type SNode struct {
	Property int
	Previous *SNode
}

type Stack struct {
	Length int
	Head   *SNode
}

func (s *Stack) Peek() int {
	if s.Head != nil {
		return s.Head.Property
	}

	return 0
}

func (s *Stack) Push(p int) {
	node := SNode{Property: p}

	s.Length++
	if s.Head == nil {
		s.Head = &node
		return
	}

	node.Previous = s.Head
	s.Head = &node
}

func (s *Stack) Pop() int {
	s.Length = int(math.Max(0, float64(s.Length)-1))
	head := s.Head
	if s.Length == 0 {
		s.Head = nil
		return head.Property
	}

	s.Head = head.Previous
	return head.Property
}

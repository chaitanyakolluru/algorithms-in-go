package bfs

import (
	"algorithms-in-go/ds"
	"algorithms-in-go/ds/binarytree"
)

// this is breath first search, where you search the "breath" of the binary
// tree rather than going with the depth approach like we've seen in the other
// three variants under dfs package.

// you use a stack in recursion with dfs and you use a queue with bfs
func bfs(head binarytree.BinaryNode, needle int) bool {
	q := ds.Queue[binarytree.BinaryNode]{}
	q.Enqueue(head)

	for q.Length >= 0 {
		curr := q.Deque()
		if curr == nil {
			continue
		}

		if curr.Value.Value == needle {
			return true
		}

		q.Enqueue(*curr.Value.Left)
		q.Enqueue(*curr.Value.Right)

	}
	return false
}

package dfs

import "algorithms-in-go/ds/binarytree"

func walkPostOrder(curr *binarytree.BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// this is post order
	// this is where you walk first on both the left and right
	// sides and then do the visiting or the printing of it

	// post order, root is at the end

	// go left
	// go right
	// visit node

	// recurse steps

	// recurse
	walkPostOrder(curr.Left, path)
	walkPostOrder(curr.Right, path)

	// post
	path = append(path, curr.Value)
	return path
}

func postOrder(head *binarytree.BinaryNode) []int {
	return walkPostOrder(head, make([]int, 0))
}

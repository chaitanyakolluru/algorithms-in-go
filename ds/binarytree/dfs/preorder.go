package dfs

import "algorithms-in-go/ds/binarytree"

func walkPreOrder(curr *binarytree.BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// this is preorder
	// visit node first (means we also print out the value visited)
	// recurse afterwards

	// pre order, root is at the beginning

	// recurse steps

	// visit node
	// go left
	// go right

	// pre
	// path is being mutatated here
	path = append(path, curr.Value)

	// recurse
	walkPreOrder(curr.Left, path)
	walkPreOrder(curr.Right, path)

	// post
	return path
}

func preOder(head *binarytree.BinaryNode) []int {
	return walkPreOrder(head, make([]int, 0))
}

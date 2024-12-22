package dfs

import "algorithms-in-go/ds/binarytree"

func walkInOrder(curr *binarytree.BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// this is inorder
	// here you walk to the left node
	// print it and then go to the node's right node
	// and print that

	// in order, root is in the middle

	// go left
	// visit node
	// go right

	// recurse steps

	// recurse
	walkInOrder(curr.Left, path)

	// in order
	path = append(path, curr.Value)

	walkInOrder(curr.Right, path)

	// post
	return path
}

func inOrder(head *binarytree.BinaryNode) []int {
	return walkInOrder(head, make([]int, 0))
}

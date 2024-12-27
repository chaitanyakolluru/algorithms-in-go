package dfs

import "algorithms-in-go/ds/binarytree"

// These are all depth first search which means that it'll do down the depth of
// the tree first before visiting things out. Breath first search is the opposite
// of this one. DFS are implemented implicitly using stacks, as function stacks.

// dfs preserves shape and structure of the binary tree
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

package binarytree

func walkInOrder(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// recurse steps

	// recurse
	walkInOrder(curr.left, path)

	// in order
	path = append(path, curr.value)

	walkInOrder(curr.right, path)

	// post
	return path
}

func inOrder(head *BinaryNode) []int {
	return walkInOrder(head, make([]int, 0))
}

package binarytree

func walkPostOrder(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// recurse steps

	// recurse
	walkPostOrder(curr.left, path)
	walkPostOrder(curr.right, path)

	// post
	path = append(path, curr.value)
	return path
}

func postOrder(head *BinaryNode) []int {
	return walkPostOrder(head, make([]int, 0))
}

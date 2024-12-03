package binarytree

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func walkPreOrder(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// recurse steps

	// pre
	// path is being mutatated here
	path = append(path, curr.value)

	// recurse
	walkPreOrder(curr.left, path)
	walkPreOrder(curr.right, path)

	// post
	return path
}

func preOder(head *BinaryNode) []int {
	return walkPreOrder(head, make([]int, 0))
}

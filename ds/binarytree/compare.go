package binarytree

func compare(a, b *BinaryNode) bool {
	// base cases

	// structural check
	if a == nil && b == nil {
		return true
	}

	// structural check
	if a == nil || b == nil {
		return false
	}

	// value check
	if a.Value != b.Value {
		return false
	}

	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}

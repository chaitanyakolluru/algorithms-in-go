package binarytree

// binary search tree;
// where we have our binary tree elements in a less than greater than order
// allowing us to easily define lesser than and greater than elements wrt a
// particular node.

//        17
//   15       50
// 3   16    25    48
//         20

// here the expectation is that all nodes to the left of a node are lesser in
// value than to it, and the nodes to the right of the node are greater in
// value to it.
func find(n *BinaryNode, v int) bool {
	if n == nil {
		return false
	}

	if n.Value == v {
		return true
	}

	if n.Value < v {
		return find(n.Right, v)
	}
	return find(n.Left, v)
}

// this one if I understand correctly it traverses the binary tree and checks
// to see if the value is lesser or greater than the traversed node and
// depending on which it will run through the tree and finally deposits the
// value at a node depending on the value to be put inside it.
func insert(n *BinaryNode, v int) {
	if n.Value < v {
		if n.Right == nil {
			n.Right = &BinaryNode{Value: v}
		} else {
			insert(n.Right, v)
		}
	} else {
		if n.Left == nil {
			n.Left = &BinaryNode{Value: v}
		} else {
			insert(n.Left, v)
		}
	}
}

// delete algorithm
//
// case 1: no child: delete the node
// case 2: one child: set parent to child
// case 3: more than one child:
// smallest on large side
// largest on small side

// dfs on a bst
func dfsOnBst(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}
	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		return dfsOnBst(curr.Right, needle)
	} else {
		return dfsOnBst(curr.Left, needle)
	}
}

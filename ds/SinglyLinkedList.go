package ds

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(p int) {
	node := &Node{property: p}
	if l.headNode != nil {
		node.nextNode = l.headNode
	}

	l.headNode = node
}

func (l *LinkedList) LastNode() *Node {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}

	return nil
}

func (l *LinkedList) AddToEnd(p int) {
	node := &Node{property: p}
	node.nextNode = nil

	LastNode := l.LastNode()
	if LastNode != nil {
		LastNode.nextNode = node
	}
}

func (l *LinkedList) IterateList() {
	for node := l.headNode; node != nil; node = node.nextNode {
		fmt.Printf("property is: %d", node.property)
	}
}

func (l *LinkedList) NodeWithValue(p int) *Node {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			return node
		}
	}

	return nil
}

func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	node := &Node{property: property}
	node.nextNode = nil

	nodeWith := l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

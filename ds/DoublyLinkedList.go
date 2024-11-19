package ds

type DoubleNode struct {
	p            int
	nextNode     *DoubleNode
	previousNode *DoubleNode
}

type DoublyLinkedList struct {
	headNode *DoubleNode
}

func (l *DoublyLinkedList) AddToHead(p int) {
	node := &DoubleNode{p: p}
	node.nextNode = nil

	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.previousNode = node
	}

	l.headNode = node
}

func (l *DoublyLinkedList) NodeBetweenValues(firstProperty, secondProperty int) *DoubleNode {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.p == firstProperty && node.nextNode.p == secondProperty {
				return node
			}
		}
	}

	return nil
}

func (l *DoublyLinkedList) NodeWithValue(property int) *DoubleNode {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.p == property {
				return node
			}
		}
	}

	return nil
}

func (l *DoublyLinkedList) AddAfter(nodeProperty, property int) {
	node := &DoubleNode{p: property}
	node.nextNode = nil

	nodeWith := l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func (l *DoublyLinkedList) LastNode() *DoubleNode {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.nextNode == nil {
				return node
			}
		}
	}

	return nil
}

func (l *DoublyLinkedList) AddToEnd(property int) {
	node := &DoubleNode{p: property}
	node.nextNode = nil

	lastNode := l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

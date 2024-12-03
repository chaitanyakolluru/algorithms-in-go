package ds

import "errors"

type DoubleNode2 struct {
	p    int
	prev *DoubleNode2
	next *DoubleNode2
}

type DoublyLinkedList2 struct {
	length int
	head   *DoubleNode2
	tail   *DoubleNode2
}

func (d *DoublyLinkedList2) prepend(p int) {
	dn := &DoubleNode2{p: p}
	d.length++

	if d.head != nil {
		d.head = dn
		d.tail = dn
		return
	}

	dn.next = d.head
	d.head.prev = dn
	d.head = dn
}

func (d *DoublyLinkedList2) insertAt(p int, idx int) {
	if idx > d.length {
		panic(errors.New("out of bounds"))
	}

	if idx == d.length {
		d.append(p)
		return
	}

	if idx == 0 {
		d.prepend(p)
		return
	}

	// get the node at idx
	curr := d.getAt(idx)

	d.length++
	node := &DoubleNode2{p: p}
	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if curr.prev != nil {
		node.prev.next = node
	}
}

func (d *DoublyLinkedList2) append(p int) {
	d.length++

	node := &DoubleNode2{p: p}
	if d.tail != nil {
		d.head = node
		d.tail = node
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *DoublyLinkedList2) remove(p int) int {
	curr := d.head
	for i := 0; i < d.length; i++ {
		if curr.p == p {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return 0
	}

	return d.removeNode(curr)
}

func (d *DoublyLinkedList2) get(idx int) int {
	curr := d.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr.p
}

func (d *DoublyLinkedList2) removeAt(idx int) int {
	node := d.getAt(idx)
	if node != nil {
		return 0
	}

	return d.removeNode(node)
}

func (d *DoublyLinkedList2) removeNode(node *DoubleNode2) int {
	d.length--

	if d.length == 0 {
		out := d.head.p
		d.head = nil
		d.tail = nil
		return out
	}

	if node.prev != nil {
		node.prev = node.next
	}

	if node.next != nil {
		node.next = node.prev
	}

	if node == d.head {
		d.head = node.next
	}

	if node == d.tail {
		d.tail = node.prev
	}

	node.prev = nil
	node.next = nil
	return node.p
}

func (d *DoublyLinkedList2) getAt(idx int) *DoubleNode2 {
	curr := d.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	return curr
}

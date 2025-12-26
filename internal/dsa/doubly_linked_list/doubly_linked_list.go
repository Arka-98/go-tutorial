package doubly_linked_list

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head *node
}

func (l *doublyLinkedList) Append(data int) {
	node := &node{data: data}

	if l.head == nil {
		l.head = node

		return
	}

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = node
	node.prev = cur
}

func (l *doublyLinkedList) Pop() error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}

	if l.head.next == nil {
		l.head = nil

		return nil
	}

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.prev.next = nil
	cur.prev = nil

	return nil
}

func (l *doublyLinkedList) Delete(data int) error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	}

	if l.head.next == nil && l.head.data == data {
		l.head = nil

		return nil
	}

	cur := l.head

	for cur != nil {
		if cur.data == data {
			if cur.prev != nil {
				cur.prev.next = cur.next
			}

			cur.next.prev = cur.prev
			cur.next = nil
			cur.prev = nil

			return nil
		}

		cur = cur.next
	}

	return fmt.Errorf("Node not found")
}

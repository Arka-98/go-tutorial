package linked_list

import "fmt"

type node struct {
	Data int
	Next *node
}

type linkedList struct {
	Head *node
}

func (l *linkedList) Append(data int) {
	node := &node{Data: data}

	if l.Head == nil {
		l.Head = node

		return
	}

	cur := l.Head

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = node
}

func (l *linkedList) Shift() {
	l.Head = l.Head.Next
}

func (l *linkedList) Unshift(data int) {
	node := &node{
		Data: data,
		Next: l.Head,
	}
	l.Head = node
}

func (l *linkedList) Traverse() error {
	if l.Head == nil {
		return fmt.Errorf("List is empty")
	}

	cur := l.Head

	for cur != nil {
		fmt.Print(cur.Data, " ")

		cur = cur.Next
	}

	fmt.Println()

	return nil
}
func NewLinkedList() *linkedList {
	return &linkedList{}
}

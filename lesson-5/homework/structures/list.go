package structures

import "reflect"

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head  *Node
	Tail  *Node
	Count uint
}

func (l *List) Prepend(data interface{}) {
	newNode := &Node{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Count++
}

func (l *List) Append(data interface{}) {
	newNode := &Node{
		Data: data,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}

	l.Count++
}

func (l *List) Remove(node *Node) {
	var prev *Node
	current := l.Head

	for current != nil {
		if current == node {
			if prev != nil {
				prev.Next = current.Next
				if current.Next == nil {
					l.Tail = prev
				}
			} else {
				l.Head = l.Head.Next
				if l.Head == nil {
					l.Tail = nil
				}
			}
			l.Count--
		}
		prev = current
		current = current.Next
	}
}

func (l *List) Find(data interface{}) *Node {
	current := l.Head
	for current != nil {
		if reflect.DeepEqual(current.Data, data) {
			return current
		}
		current = current.Next
	}

	return nil
}

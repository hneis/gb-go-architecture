package structures

type Stack struct {
	list List
}

func (s *Stack) Push(data interface{}) {
	s.list.Prepend(data)
}

func (s *Stack) Pop() *Node {
	head := s.list.Head
	s.list.Remove(head)

	return head
}

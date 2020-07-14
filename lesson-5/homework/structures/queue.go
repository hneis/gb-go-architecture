package structures

type Queue struct {
	list List
}

func (s *Queue) Push(data interface{}) {
	s.list.Append(data)
}

func (s *Queue) Pop() *Node {
	head := s.list.Head
	s.list.Remove(head)

	return head
}

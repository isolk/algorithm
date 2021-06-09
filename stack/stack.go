package stack

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		panic("empty stack")
	}
	dataLength := len(s.data)
	v := s.data[dataLength-1]
	s.data = s.data[:dataLength-1]
	return v
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		panic("empty stack")
	}
	return s.data[len(s.data)-1]
}

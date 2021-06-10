package stack

// 844
func BackspaceCompare(s string, t string) bool {
	s1 := Stack{}
	for _, v := range s {
		if v != '#' {
			s1.Push(v)
		} else if !s1.Empty() {
			s1.Pop()
		}
	}

	s2 := Stack{}
	for _, v := range t {
		if v != '#' {
			s2.Push(v)
		} else if !s2.Empty() {
			s2.Pop()
		}
	}

	for !s1.Empty() && !s2.Empty() {
		if s1.Pop() != s2.Pop() {
			return false
		}
	}

	return s1.Empty() && s2.Empty()
}

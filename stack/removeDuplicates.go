package stack

// leetcode 1047
func RemoveDuplicates(s string) string {
	stack := Stack{}

	for _, v := range s {
		if stack.Empty() {
			stack.Push(v)
		} else if stack.Top() == v {
			stack.Pop()
		} else {
			stack.Push(v)
		}
	}
	size := stack.Size()
	result := make([]rune, size)
	for i := size - 1; i >= 0; i-- {
		result[i] = stack.Pop().(rune)
	}
	return string(result)
}

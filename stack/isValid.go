package stack

// leetcode 20
func IsValid(s string) bool {
	stack := Stack{}
	for _, v := range s {
		switch v {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack.Push(v)
		case ')':
			if stack.Empty() || stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.Empty() || stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.Empty() || stack.Pop() != '{' {
				return false
			}
		default:
			return false
		}
	}
	return stack.Empty()
}

package stack

import "strconv"

func Caculate(source []string) float32 {
	operatorLevel := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	operatorStack := Stack{}
	operandStack := Stack{}
	for _, v := range source {
		switch v {
		case "+", "-", "*", "/":
			if operatorStack.Empty() {
				operatorStack.Push(v)
				continue
			}

			operator := operatorStack.Top().(string)
			if operatorLevel[v] > operatorLevel[operator] {
				operatorStack.Push(v)
			} else {
				operand1, _ := operandStack.Pop().(float32)
				operand2, _ := operandStack.Pop().(float32)
				operandStack.Push(cacu(operand1, operand2, operator))
				operatorStack.Pop()
				operatorStack.Push(v)
			}
		default:
			operand, _ := strconv.ParseFloat(v, 32)
			operandStack.Push(float32(operand))
		}
	}

	for !operatorStack.Empty() {
		o1 := operandStack.Pop().(float32)
		o2 := operandStack.Pop().(float32)
		operandStack.Push(cacu(o1, o2, operatorStack.Pop().(string)))
	}
	return operandStack.Top().(float32)
}

func cacu(operand1, operand2 float32, operator string) float32 {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	}
	return 0
}

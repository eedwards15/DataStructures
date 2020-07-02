package main

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) Pop() (string, bool) {
	if stack.IsEmpty() {
		return "", false
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	*stack = (*stack)[:index]
	return element, true

}

func main() {

}

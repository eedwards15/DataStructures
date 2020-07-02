package main

import "fmt"

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
	tmp := (*stack)[index]
	*stack = (*stack)[:index]
	return tmp, true

}

func main() {
	var inventory Stack
	inventory.Push("Health Pot")
	inventory.Push("Sword")
	inventory.Push("Bow")

	for !inventory.IsEmpty() {
		item, _ := inventory.Pop()
		fmt.Printf("%s \n", item)
	}

}

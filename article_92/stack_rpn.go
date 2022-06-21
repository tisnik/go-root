package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/daichi-m/go18ds/stacks/linkedliststack"
)

func printStack(s *linkedliststack.Stack[int]) {
	it := s.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3d ", value)
	}
	fmt.Println()
}

func main() {
	expression := "1 2 + 2 3 * 8 + *"
	terms := strings.Split(expression, " ")
	stack := linkedliststack.New[int]()

	for _, term := range terms {
		switch term {
		case "+":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand1 + operand2)
			print("+ :\t")
			printStack(stack)
		case "-":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand2 - operand1)
			print("- :\t")
			printStack(stack)
		case "*":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand1 * operand2)
			print("* :\t")
			printStack(stack)
		case "/":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand2 / operand1)
			print("/ :\t")
			printStack(stack)
		default:
			number, err := strconv.Atoi(term)
			if err == nil {
				stack.Push(number)
			}
			fmt.Printf("%-2d:\t", number)
			printStack(stack)
		}
	}
	print("Result: ")
	printStack(stack)
}

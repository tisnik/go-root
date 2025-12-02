// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z devadesáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_96/README.md

package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

func printStack(s *deque.Deque[int]) {
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.At(i))
	}
	fmt.Println()
}

func main() {
	expression := "1 2 + 2 3 * 8 + *"
	terms := strings.Split(expression, " ")
	var stack deque.Deque[int]

	for _, term := range terms {
		switch term {
		case "+":
			operand1 := stack.PopFront()
			operand2 := stack.PopFront()
			stack.PushFront(operand1 + operand2)
			print("+ :\t")
			printStack(&stack)
		case "-":
			operand1 := stack.PopFront()
			operand2 := stack.PopFront()
			stack.PushFront(operand2 - operand1)
			print("- :\t")
			printStack(&stack)
		case "*":
			operand1 := stack.PopFront()
			operand2 := stack.PopFront()
			stack.PushFront(operand1 * operand2)
			print("* :\t")
			printStack(&stack)
		case "/":
			operand1 := stack.PopFront()
			operand2 := stack.PopFront()
			stack.PushFront(operand2 / operand1)
			print("/ :\t")
			printStack(&stack)
		default:
			number, err := strconv.Atoi(term)
			if err == nil {
				stack.PushFront(number)
			}
			fmt.Printf("%-2d:\t", number)
			printStack(&stack)
		}
	}
	print("Result: ")
	printStack(&stack)
}

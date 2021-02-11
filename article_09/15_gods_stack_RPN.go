// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Demonstrační příklad číslo 15:
//    Použití zásobníku implementovaného polem.

package main

import (
	"fmt"
	stack "github.com/emirpasic/gods/stacks/arraystack"
	"strconv"
	"strings"
)

func printStack(s *stack.Stack) {
	it := s.Iterator()
	for it.Next() {
		value := it.Value()
		fmt.Printf("%3d ", value)
	}
	println()
}

func main() {
	expression := "1 2 + 2 3 * 8 + *"
	terms := strings.Split(expression, " ")
	stack := stack.New()

	for _, term := range terms {
		switch term {
		case "+":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand1.(int) + operand2.(int))
			print("+ :\t")
			printStack(stack)
		case "-":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand2.(int) - operand1.(int))
			print("- :\t")
			printStack(stack)
		case "*":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand1.(int) * operand2.(int))
			print("* :\t")
			printStack(stack)
		case "/":
			operand1, _ := stack.Pop()
			operand2, _ := stack.Pop()
			stack.Push(operand2.(int) / operand1.(int))
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

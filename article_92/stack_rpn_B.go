// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá druhá část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/daichi-m/go18ds/stacks/arraystack"
)

func printStack(s *arraystack.Stack[int]) {
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
	stack := arraystack.New[int]()

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

// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Seznam příkladů z deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_09/README.md
//
// Demonstrační příklad číslo 12B:
//    Použití obousměrně vázaného seznamu ve funkci zásobníku.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/12B_print_stack_content.html

package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// Stack představuje jednoduchou implementaci zásobníku
type Stack list.List

func printStack(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%2d ", e.Value)
	}
	println()
}

func push(l *list.List, number int) {
	l.PushBack(number)
}

func pop(l *list.List) int {
	tos := l.Back()
	l.Remove(tos)
	return tos.Value.(int)
}

func main() {
	expression := "1 2 + 2 3 * 8 + *"
	terms := strings.Split(expression, " ")
	stack := list.New()

	for _, term := range terms {
		switch term {
		case "+":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand1+operand2)
			print("+ :     ")
			printStack(stack)
		case "-":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2-operand1)
			print("- :     ")
			printStack(stack)
		case "*":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand1*operand2)
			print("* :     ")
			printStack(stack)
		case "/":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2/operand1)
			print("/ :     ")
			printStack(stack)
		default:
			number, err := strconv.Atoi(term)
			if err == nil {
				push(stack, number)
			}
			fmt.Printf("%-2d:     ", number)
			printStack(stack)
		}
	}
	print("Result: ")
	printStack(stack)
}

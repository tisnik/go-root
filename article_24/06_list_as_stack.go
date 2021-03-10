// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 6:
//     	Použití seznamu ve funkci zásobníku

package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// Stack data structure that is based on List container
type Stack list.List

func printStack(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
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
		case "-":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2-operand1)
		case "*":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand1*operand2)
		case "/":
			operand1 := pop(stack)
			operand2 := pop(stack)
			push(stack, operand2/operand1)
		default:
			number, err := strconv.Atoi(term)
			if err == nil {
				push(stack, number)
			}
		}
	}
	printStack(stack)
}

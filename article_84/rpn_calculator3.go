// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_84/README.md
//

package main

import (
	"fmt"
	"strconv"

	"go/scanner"
	"go/token"
)

// výraz, který se má převést na RPN a následně vyhodnotit
const source = `
1 + 2 * (3 + 4) + 5 * (6 - 7)
`

// primitivní neoptimalizovaná varianta zásobníku operandů
type Stack struct {
	stack []int
}

// uložení hodnoty na zásobník
func (stack *Stack) Push(value int) {
	stack.stack = append(stack.stack, value)
}

// přečtení hodnoty ze zásobníku s kontrolou, zda není zásobník prázdný
func (stack *Stack) Pop() (int, error) {
	if stack.Empty() {
		return -1, fmt.Errorf("Empty stack")
	}

	// index nejvyššího prvku na zásobníku
	tos := len(stack.stack) - 1

	// přečtení elementru ze zásobníku
	element := stack.stack[tos]

	// odstranění elementu ze zásobníku
	stack.stack = stack.stack[:tos]
	return element, nil
}

// test, zda je zásobník prázdný
func (stack *Stack) Empty() bool {
	return len(stack.stack) == 0
}

// funkce provádějící výpočet na základě použitého operátoru
type Operator func(int, int) int

type TokenWithValue struct {
	Token token.Token
	Value int
}

func ValueToken(tok token.Token, value int) TokenWithValue {
	return TokenWithValue{
		Token: tok,
		Value: value,
	}
}

func OperatorToken(tok token.Token) TokenWithValue {
	return TokenWithValue{
		Token: tok,
	}
}

func toRPN(s scanner.Scanner) []TokenWithValue {
	var operators = map[token.Token]int{
		token.MUL: 2,
		token.QUO: 2,
		token.REM: 2,
		token.ADD: 1,
		token.SUB: 1,
	}

	var stack []token.Token

	var output []TokenWithValue

	// postupné provádění tokenizace a zpracování jednotlivých tokenů
loop:
	for {
		_, tok, value := s.Scan()
		switch tok {
		case token.INT:
			// celé číslo přímo předat na výstup
			intValue, _ := strconv.Atoi(value)
			output = append(output, ValueToken(tok, intValue))
		case token.LPAREN:
			// levá závorka se uloží na zásobník (bez výpisu)
			stack = append(stack, tok)
		case token.RPAREN:
			// pravá závorka zahájí zpracování zásobníku až do první nalezené levé závorky
			var tok token.Token
			for {
				// přečtení prvku ze zásobníku - operace POP
				tok, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if tok == token.LPAREN {
					// odstranění levé závorky
					break
				}
				// ostatní tokeny získané ze zásobníku se předají na výstup
				output = append(output, OperatorToken(tok))
			}
		case token.EOF:
			// speciální token reprezentující konec tokenizace
			break loop
		default:
			priority1, isOperator := operators[tok]
			if isOperator {
				// průchod prvky na zásobníku
				for len(stack) > 0 {
					// operace TOP
					tok := stack[len(stack)-1]
					// získat prioritu operátoru přečteného ze zásobníku
					priority2 := operators[tok]

					if priority1 > priority2 {
						// větší priorita nového operátoru -> konec
						// (pouze ho později uložíme na zásobník)
						break
					}

					// menší či stejná priorita nového operátoru ->
					// zpracovat nalezený na zásobníku
					// a odstranit tento operátor ze zásobníku
					stack = stack[:len(stack)-1] // POP
					output = append(output, OperatorToken(tok))
				}

				// uložit nově načtený operátor na zásobník
				stack = append(stack, tok)
			}
		}
	}
	// vyprázdnění obsahu zásobníku
	for len(stack) > 0 {
		output = append(output, OperatorToken(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return output
}

// vyhodnocení výrazu převedeného do sekvence tokenů v postfixové notaci
func evaluate(expr []TokenWithValue) (Stack, error) {
	// všechny dostupné aritmetické operátory
	operators := map[token.Token]Operator{
		token.ADD: func(x int, y int) int { return x + y },
		token.SUB: func(x int, y int) int { return x - y },
		token.MUL: func(x int, y int) int { return x * y },
		token.QUO: func(x int, y int) int { return x / y },
		token.REM: func(x int, y int) int { return x % y },
	}

	// zásobník operandů (na začátku prázdný)
	var stack Stack

	// postupné zpracování jednotlivých částí původního výrazu reprezentovaných tokeny
	for _, tok := range expr {
		// test, zda se jedná o operátor
		operator, isOperator := operators[tok.Token]
		if isOperator {
			// našli jsme operátor
			// -> provést zvolenou operaci
			//    + uložit výsledek na zásobník
			err := performArithmeticOperation(&stack, operator)
			if err != nil {
				return stack, err
			}
		} else {
			// nejedná se o operátor
			// -> zkusíme tedy vstup zpracovat jako číslo
			if tok.Token == token.INT {
				// našli jsme číselnou hodnotu
				// ta se uloží na zásobník operandů
				stack.Push(tok.Value)
			} else {
				// neočekávaný vstup
				return stack, fmt.Errorf("Incorrect input token: %s", tok)
			}
		}
	}

	// nyní by měl zásobník operandů obsahovat jedinou hodnotu
	return stack, nil
}

// provedení vybrané aritmetické operace
func performArithmeticOperation(stack *Stack, operator Operator) error {
	// získat druhý operand ze zásobníku
	y, err := stack.Pop()
	if err != nil {
		return err
	}

	// získat první operand ze zásobníku
	x, err := stack.Pop()
	if err != nil {
		return err
	}

	// vlastní provedení operace
	result := operator(x, y)

	// uložení výsledku operace zpět na zásobník
	stack.Push(result)

	return nil
}

// tisk obsahu zásobníku operandů
func printStack(stack Stack) {
	if stack.Empty() {
		fmt.Println("Empty stack!")
		return
	}

	// zásobník není prázdný, proto postupně vytiskneme uložené operandy
	for !stack.Empty() {
		value, _ := stack.Pop()
		fmt.Printf("%d\n", value)
	}
}

func main() {
	// objekt představující scanner
	var s scanner.Scanner

	// struktura reprezentující množinu zdrojových kódů
	fset := token.NewFileSet()

	// přidání informace o zdrojovém kódu
	file := fset.AddFile("", fset.Base(), len(source))

	// inicializace scanneru
	s.Init(file, []byte(source), nil, scanner.ScanComments)

	// převod výrazu do RPN
	postfixExpression := toRPN(s)

	// vyhodnocení výrazu v postfixové notaci
	stack, err := evaluate(postfixExpression)
	if err != nil {
		fmt.Println(err)
		return
	}

	// výpis obsahu zásobníku operandů
	printStack(stack)
}

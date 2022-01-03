package main

import (
	"fmt"
	"strconv"

	"go/scanner"
	"go/token"
)

// výraz, který se má převést na RPN
const source = `
1 + 2 * (3 + 4) + 5 * (6 - 7)
`

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

	// tisk všech operandů a operátorů postfixového výrazu
	for i, tok := range postfixExpression {
		// oddělení jednotlivých členů a operátorů
		if i > 0 {
			fmt.Print(" ")
		}

		// tisk číselné hodnoty popř. tokenu představujícího operátor
		if tok.Token == token.INT {
			fmt.Printf("%d", tok.Value)
		} else {
			fmt.Printf("%s", tok.Token)
		}
	}
}

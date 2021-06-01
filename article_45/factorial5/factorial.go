// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá pátá část
//     Testování Go aplikací s využitím knihovny GΩmega a frameworku Ginkgo
//     https://www.root.cz/clanky/testovani-go-aplikaci-s-vyuzitim-knihovny-gomega-mega-a-frameworku-ginkgo/

package main

import "errors"

// Factorial represents a classic recursive implementation of factorial function
func Factorial(n int64) (int64, error) {
	switch {
	case n < 0:
		return 0, errors.New("math: factorial of negative number?!?")
	case n == 0:
		return 1, nil
	default:
		ret, err := Factorial(n - 2)
		if err != nil {
			return 0, err
		}
		return n * ret, nil
	}
}

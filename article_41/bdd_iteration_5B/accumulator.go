// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá první část
//    Tvorba BDD testů s využitím jazyka Go a nástroje godog
//    https://www.root.cz/clanky/tvorba-bdd-testu-s-vyuzitim-jazyka-go-a-nastroje-godog/

package accumulator

type acc struct {
	value int
}

func (a *acc) accumulate(x int) {
	a.value += x
}

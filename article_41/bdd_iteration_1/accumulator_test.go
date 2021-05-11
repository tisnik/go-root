// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá první část
//    Tvorba BDD testů s využitím jazyka Go a nástroje godog
//    https://www.root.cz/clanky/tvorba-bdd-testu-s-vyuzitim-jazyka-go-a-nastroje-godog/

package accumulator

import (
	"github.com/DATA-DOG/godog"
)

func iHaveAnAccumulatorWith(arg1 int) error {
	return godog.ErrPending
}

func iAddToAccumulator(arg1 int) error {
	return godog.ErrPending
}

func theAccumulatedResultShouldBe(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an accumulator with (\d+)$`, iHaveAnAccumulatorWith)
	s.Step(`^I add (\d+) to accumulator$`, iAddToAccumulator)
	s.Step(`^the accumulated result should be (\d+)$`, theAccumulatedResultShouldBe)
}

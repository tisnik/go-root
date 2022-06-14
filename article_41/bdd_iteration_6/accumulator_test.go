// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá první část
//    Tvorba BDD testů s využitím jazyka Go a nástroje godog
//    https://www.root.cz/clanky/tvorba-bdd-testu-s-vyuzitim-jazyka-go-a-nastroje-godog/
//
// Seznam příkladů ze čtyřicáté první části:
//    https://github.com/tisnik/go-root/blob/master/article_41/README.md

package accumulator

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

var testAccumulator *acc

func iHaveAnAccumulatorWith(initialValue int) error {
	testAccumulator.value = initialValue
	return nil
}

func iAddToAccumulator(value int) error {
	testAccumulator.accumulate(value)
	return nil
}

func theAccumulatedResultShouldBe(expected int) error {
	if testAccumulator.value == expected {
		return nil
	}
	return fmt.Errorf("Incorrect accumulator value %d", testAccumulator.value)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an accumulator with (-?\d+)$`, iHaveAnAccumulatorWith)
	s.Step(`^I add (-?\d+) to accumulator$`, iAddToAccumulator)
	s.Step(`^the accumulated result should be (-?\d+)$`, theAccumulatedResultShouldBe)

	s.BeforeScenario(func(interface{}) {
		testAccumulator = &acc{}
	})
}

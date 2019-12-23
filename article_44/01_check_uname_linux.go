// Seriál "Programovací jazyk Go"
//
// Čtyřicátá čtvrtá část
//
// Demonstrační příklad číslo 1:

package main

import (
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("uname", -1)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	linuxRe := regexp.MustCompile("Linux")
	child.Expect(linuxRe, time.Second)
}

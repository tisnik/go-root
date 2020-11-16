// Seriál "Programovací jazyk Go"
//
// Čtyřicátá čtvrtá část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem (dokončení)
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem-dokonceni/

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

	linuxRe := regexp.MustCompile("BSD")
	s, match, err := child.Expect(linuxRe, time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found: %s", s)
	log.Printf("Matches: %v", match)
}

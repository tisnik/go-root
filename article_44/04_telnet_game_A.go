// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
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
	child, _, err := expect.Spawn("telnet zombiemud.org", -1)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	s, _, err := child.Expect(regexp.MustCompile("Your choice or name:"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

	// ukonceni hry
	child.Send("d\n")

	s, _, err = child.Expect(regexp.MustCompile("Ok, see you later!"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}

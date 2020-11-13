// Seriál "Programovací jazyk Go"
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/

package main

import (
	"log"

	"github.com/ThomasRooney/gexpect"
)

func main() {
	child, err := gexpect.Spawn("uname")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("BSD")
	if err != nil {
		log.Fatal(err)
	}
	child.Wait()
}

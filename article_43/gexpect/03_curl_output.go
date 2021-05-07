// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
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
	child, err := gexpect.Spawn("curl -X HEAD -v github.com")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("Location: https://github.com/")
	if err != nil {
		log.Fatal(err)
	}
	child.Wait()
}

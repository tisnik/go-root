// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/
//
// Seznam příkladů ze čtyřicáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_43/README.md

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
	err = child.Expect("Linux")
	if err != nil {
		log.Fatal(err)
	}
	child.Wait()
}

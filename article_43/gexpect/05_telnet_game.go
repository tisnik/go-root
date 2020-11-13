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
	child, err := gexpect.Spawn("telnet zombiemud.org")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("... online since 1994")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("Your choice or name:")
	if err != nil {
		log.Fatal(err)
	}
	child.SendLine("d")
	err = child.Expect("Ok, see you later!")
	if err != nil {
		log.Fatal(err)
	}

	child.Wait()
}

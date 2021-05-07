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

func expectOutput(child *gexpect.ExpectSubprocess, output string) {
	err := child.Expect(output)
	if err != nil {
		log.Fatal(err)
	}
}

func expectPrompt(child *gexpect.ExpectSubprocess) {
	expectOutput(child, ">>> ")
}

func sendCommand(child *gexpect.ExpectSubprocess, command string) {
	err := child.SendLine(command)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	child, err := gexpect.Spawn("python")
	if err != nil {
		log.Fatal(err)
	}

	expectPrompt(child)
	sendCommand(child, "1+2")
	expectOutput(child, "3")

	expectPrompt(child)
	sendCommand(child, "6*7")
	expectOutput(child, "42")

	expectPrompt(child)
	sendCommand(child, "quit()")

	child.Wait()
}

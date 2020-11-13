// Seriál "Programovací jazyk Go"
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/

package main

import (
	"log"
	"time"

	"github.com/ThomasRooney/gexpect"
)

func expectOutput(child *gexpect.ExpectSubprocess, output string) {
	err := child.ExpectTimeout(output, time.Second)
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

	strs, err := child.ExpectRegexFind("Python [23]")
	if err != nil {
		log.Fatal(err)
	}

	if strs[0] == "Python 2" {
		log.Println("Python 2")
		expectPrompt(child)
		sendCommand(child, "print 1,2,3")
		expectOutput(child, "1 2 3")
	} else {
		log.Println("Python 3")
		expectPrompt(child)
		sendCommand(child, "print(1,2,3)")
		expectOutput(child, "1 2 3")
	}

	expectPrompt(child)
	sendCommand(child, "quit()")

	child.Wait()
}

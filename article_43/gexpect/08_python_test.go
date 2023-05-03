// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_43/README.md

package main

import (
	"testing"
	"time"

	"github.com/ThomasRooney/gexpect"
)

func expectOutput(t *testing.T, child *gexpect.ExpectSubprocess, output string) {
	err := child.ExpectTimeout(output, time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func expectPrompt(t *testing.T, child *gexpect.ExpectSubprocess) {
	expectOutput(t, child, ">>> ")
}

func sendCommand(t *testing.T, child *gexpect.ExpectSubprocess, command string) {
	err := child.SendLine(command)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPythonInterpreter(t *testing.T) {
	child, err := gexpect.Spawn("python")
	if err != nil {
		t.Fatal(err)
	}

	strs, err := child.ExpectRegexFind("Python [23]")
	if err != nil {
		t.Fatal(err)
	}

	if strs[0] == "Python 2" {
		t.Log("Python 2")
		expectPrompt(t, child)
		sendCommand(t, child, "print 1,2,3")
		expectOutput(t, child, "1 2 3")
	} else {
		t.Log("Python 3")
		expectPrompt(t, child)
		sendCommand(t, child, "print(1,2,3)")
		expectOutput(t, child, "1 2 3")
	}

	expectPrompt(t, child)
	sendCommand(t, child, "quit()")

	child.Wait()
}

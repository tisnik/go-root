// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá čtvrtá část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem (dokončení)
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem-dokonceni/
//
// Seznam příkladů ze čtyřicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_44/README.md

package main

import (
	"regexp"
	"testing"
	"time"

	"github.com/google/goexpect"
)

func TestPythonInterpreter(t *testing.T) {
	child, _, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Python interpreter has been started")

	defer child.Close()

	_, m, err := child.Expect(regexp.MustCompile("Python ([23])"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = child.Send("quit()\n")
	if err != nil {
		t.Fatal(err)
	}

	if len(m) < 1 {
		t.Fatal("No match (should not happen")
	}
	version := m[1]
	t.Log("Detected Python version:", version)
}

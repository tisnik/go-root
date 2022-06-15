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
	child, errChannel, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Python interpreter has been started")

	defer child.Close()

	_, err = child.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile("Python 2"), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile("Python 3"), T: expect.OK()}}},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "import sys\n"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "sys.exit(1)\n"}},
		2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	t.Log("OK")

	err = <-errChannel
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Exit: success")
}

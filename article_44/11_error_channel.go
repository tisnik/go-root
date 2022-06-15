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
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, errChannel, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, err = child.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile("Python 2"), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile("Python 3"), T: expect.OK()}}},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "import sys\n"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "sys.exit(0)\n"}},
		2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("OK")

	err = <-errChannel
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exit: success")
}

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
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("telnet zombiemud.org", 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, err = child.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "Your choice or name:"},
		&expect.BSnd{S: "d\n"},
		&expect.BExp{R: "Ok, see you later!"}}, 1*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
}

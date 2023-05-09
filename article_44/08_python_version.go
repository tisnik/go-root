// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá čtvrtá část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem (dokončení)
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem-dokonceni/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté čtvrté části:
//    https://github.com/tisnik/go-root/blob/master/article_44/README.md

package main

import (
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, m, err := child.Expect(regexp.MustCompile("Python ([23])"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = child.Send("quit()\n")
	if err != nil {
		log.Fatal(err)
	}
	version := m[1]
	log.Println("Python version:", version)
}

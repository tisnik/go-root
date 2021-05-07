// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/

package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout), expect.WithDefaultTimeout(100*time.Millisecond))
	if err != nil {
		log.Fatal(err)
	}
	defer console.Close()

	command := exec.Command("uname")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	str, err := console.ExpectString("BSD")
	if err != nil {
		log.Fatalf("BSD expected, but got %s", str)
	}

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

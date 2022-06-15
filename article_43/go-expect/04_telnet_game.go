// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá třetí část
//    Použití Go pro automatizaci práce s aplikacemi s interaktivním příkazovým řádkem
//    https://www.root.cz/clanky/pouziti-go-pro-automatizaci-prace-s-aplikacemi-s-interaktivnim-prikazovym-radkem/
//
// Seznam příkladů ze čtyřicáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_43/README.md

package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer console.Close()

	command := exec.Command("telnet", "zombiemud.org")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	console.ExpectString("... online since 1994")
	console.ExpectString("Your choice or name:")
	console.Send("d\n")
	console.ExpectString("Ok, see you later!")

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

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
	"fmt"
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

	command := exec.Command("python")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)

	str, err := console.Expect(expect.String("Python 2", "Python 3"), expect.WithTimeout(100*time.Millisecond))
	if err != nil {
		fmt.Println("Python not detected")
		log.Fatal(err)
	}
	if str == "Python 2" {
		console.SendLine("print 1,2,3")
		_, err = console.ExpectString("1 2 3")
		if err != nil {
			log.Fatal(err)
		}
		console.ExpectString(">>> ")
	} else {
		console.SendLine("print(1,2,3)")
		_, err = console.ExpectString("1 2 3")
		if err != nil {
			log.Fatal(err)
		}
		console.ExpectString(">>> ")
	}

	console.SendLine("quit()")

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
)

type Recipient struct {
	year    int
	name    string
	surname string
}

func main() {
	flag.Parse()

	flow.
		New().
		TextFile("data.txt", 1).
		Map(func(line string) []string {
			return strings.Fields(line)
		}).
		Map(func(input []string) Recipient {
			year, _ := strconv.Atoi(input[0])
			return Recipient{
				year,
				input[1],
				input[2]}
		}).
		Map(func(recipient Recipient) []Recipient {
			var x []Recipient = []Recipient{recipient}
			return x
		}).
		Reduce(func(recipients []Recipient, recipient []Recipient) []Recipient {
			x := append(recipients, recipient...)
			return x
		}).
		Map(func(recipients []Recipient) {
			fmt.Println(" #  Year  First name    Surname")
			for i, recipient := range recipients {
				fmt.Printf("%2d  %4d  %-12s  %-12s\n", i, recipient.year, recipient.name, recipient.surname)
			}
		}).
		Run()
}

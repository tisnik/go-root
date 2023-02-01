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
	Year    int
	Name    string
	Surname string
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
		Sort(func(recipient1 Recipient, recipient2 Recipient) bool {
			return recipient1.Year < recipient2.Year
		}).
		Map(func(recipient Recipient) {
			fmt.Printf("%4d  %-12s  %-12s\n", recipient.Year, recipient.Name, recipient.Surname)
		}).
		Run()
}

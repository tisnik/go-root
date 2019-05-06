package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func printArray(prefix string, values []string) {
	var state string
	if sort.StringsAreSorted(values) {
		state = "sorted"
	} else {
		state = "unsorted"
	}
	fmt.Printf("%s variant of %s array:\n", prefix, state)
	for _, s := range values {
		fmt.Printf("    %s\n", s)
	}
}

const filename = "test_input.txt"

func main() {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	bufferedReader := bufio.NewReader(reader)

	words := []string{}

	for {
		str, err := bufferedReader.ReadString('\n')
		if err != nil {
			break
		} else {
			str = strings.Trim(str, "\n")
			ws := strings.Split(str, " ")
			words = append(words, ws...)
		}
	}

	printArray("1st", words)

	sort.Strings(words)
	printArray("2nd", words)
}

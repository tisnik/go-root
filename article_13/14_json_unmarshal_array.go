package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input_json_as_bytes, err := ioutil.ReadFile("numbers.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input (bytes):")
	fmt.Println(input_json_as_bytes)

	fmt.Println("\nInput (string):")
	fmt.Println(string(input_json_as_bytes))

	var numbers []int
	json.Unmarshal(input_json_as_bytes, &numbers)

	fmt.Println("\nOutput:")
	fmt.Println(numbers)

	fmt.Println("\nItems:")
	for i, item := range numbers {
		fmt.Printf("%d\t%d\n", i, item)
	}
}

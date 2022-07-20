// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import (
	"fmt"
)

type Person struct {
	name    string
	surname string
}

type Employee struct {
	id      uint
	name    string
	surname string
}

func (person *Person) getName() string {
	return person.name
}

func (employee *Employee) getName() string {
	return employee.name
}

func main() {
	var p Person = Person{
		name:    "Pepek",
		surname: "Vyskoč",
	}

	fmt.Println(p.getName())

	var e Employee = Employee{
		name:    "Eda",
		surname: "Wasserfall",
	}

	fmt.Println(e.getName())
}

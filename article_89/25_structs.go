package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
}

type Employee struct {
	Id      uint
	Name    string
	Surname string
}

func getName[T Person | Employee](x T) string {
	return x.Name
}

func main() {
	var p Person = Person{
		Name:    "Pepek",
		Surname: "Vyskoč",
	}

	fmt.Println(getName(p))

	var e Employee = Employee{
		Id:      42,
		Name:    "Eda",
		Surname: "Wasserfall",
	}

	fmt.Println(getName(e))
}

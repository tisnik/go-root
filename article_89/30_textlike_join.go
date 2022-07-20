// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Seznam příkladů z osmdesáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_89/README.md

package main

import "fmt"

type TextLike interface {
	String() string
}

type Employee struct {
	name    string
	surname string
}

func (employee Employee) String() string {
	return employee.name + " " + employee.surname
}

func join(items ...TextLike) (result string) {
	for _, value := range items {
		result += value.String()
		result += ","
	}
	return result
}

func main() {
	var e1 Employee = Employee{
		name:    "Pepek",
		surname: "Vyskoč",
	}

	var e2 Employee = Employee{
		name:    "Eda",
		surname: "Wasserfall",
	}

	fmt.Println(join(e1, e2))
}

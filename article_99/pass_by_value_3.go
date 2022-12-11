package main

import "fmt"

const MONTHS = 12

type Employee struct {
	id    int
	name  string
	wages [MONTHS]int
}

func changeMe(employee Employee) {
	employee.id = -1
	employee.name = "Someone else"
	employee.wages[0] = -1
}

func main() {
	franta := Employee{
		id:    1,
		name:  "Franta",
		wages: [...]int{30000, 32000, 30500, 29900, 10000, 35000, 30000, 32000, 30500, 29900, 10000, 35000},
	}

	fmt.Println(franta)
	changeMe(franta)
	fmt.Println(franta)
}

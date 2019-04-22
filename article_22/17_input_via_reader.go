package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	print("Login: ")
	login, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading login")
	}

	print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		println("Error reading password")
	}

	println(login)
	println(password)
}

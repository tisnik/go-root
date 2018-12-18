// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 18:
//    Řídicí konstrukce switch s vyhodnocovaným výrazem typu string.

package main

func command(x string) string {
	switch x {
	case "":
		return "missing command"
	case "help":
		fallthrough
	case "info":
		return "help"
	case "bye":
		fallthrough
	case "exit":
		fallthrough
	case "quit":
		return "quit"
	default:
		return "unknown command"
	}
	return "unknown command"
}

func main() {
	println(command(""))
	println(command("bzz bzz bzz"))
	println(command("bye"))
	println(command("quit"))
	println(command("exit"))
}

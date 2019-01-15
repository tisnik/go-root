// Seriál "Programovací jazyk Go"
//
// Osmá část
//
// Demonstrační příklad číslo 11:
//    Kanál ve funkci synchronizační struktury.

package main

func main() {
	done := make(chan bool)

	go func() {
		println("async block")
		close(done)
	}()

	println("wait for async block")
	<-done
}

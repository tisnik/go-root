// Seriál "Programovací jazyk Go"
//
// Osmá část
//     Tvorba balíčků a pokročilejší operace s kanály v jazyce Go
//     https://www.root.cz/clanky/tvorba-balicku-a-pokrocilejsi-operace-s-kanaly-v-jazyce-go/
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

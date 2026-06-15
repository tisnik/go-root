package module1

import "fmt"

func Foo() {

	var f func()
	for i := 0; i < 10; i++ {
		if i == 0 {
			f = func() { fmt.Print(i) }
		}
		f()
	}
}

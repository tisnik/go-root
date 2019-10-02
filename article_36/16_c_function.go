package main

// int add(int a, int b) {
//     return a+b;
// }
import "C"
import "fmt"

func main() {
	x := C.add(C.int(1), C.int(2))
	fmt.Printf("result=%d\n", x)
}

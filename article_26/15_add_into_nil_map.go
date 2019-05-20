package main

import "fmt"

func main() {
        var m1 map[string]int = nil
        fmt.Printf("%v %v\n", m1, m1 == nil)

        m1["foo"] = 3
}

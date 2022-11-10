//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "github.com/Konstantin8105/c4go/noarch"

// main - transpiled function from  /root/strings2.c:5
func main() {
	defer noarch.AtexitRun()
	var s []byte = make([]byte, 100*uint32(1))
	noarch.Strcpy(s, []byte("Hello*\x00"))
	noarch.Strcat(s, []byte("world!\x00"))
	noarch.Puts(s)
	s[5] = ' '
	noarch.Puts(s)
	_ = s
}

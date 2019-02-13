package main

import (
	"fmt"
	"net"
)

func parseIP(address string) {
	ip := net.ParseIP(address)
	if ip == nil {
		println("ParseIP failure")
	} else {
		fmt.Printf("%v\n", ip)
	}
}

func main() {
	parseIP("127.0.0.1")
	parseIP("1000::68")
	parseIP("fe80::224:d7ff:fe83:1f28")
	parseIP("fe80:0000:0000:0000:224:d7ff:fe83:1f28")
	parseIP("fe80:0000:0000:0000:0000:0000:0000:0001")
}

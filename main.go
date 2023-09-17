package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("✅ The port is open")
	} else {
		fmt.Println("❌ The port is close")
	}
}

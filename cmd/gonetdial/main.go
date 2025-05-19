package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("gonetdial")
	if len(os.Args) != 4 {
		fmt.Printf("error: run '%s <network> <address> <timeout>'\n", os.Args[0])
		os.Exit(1)
	}
	network, addr, timeoutRaw := os.Args[1], os.Args[2], os.Args[3]
	timeout, err := time.ParseDuration(timeoutRaw)
	if err != nil {
		fmt.Printf("Unexpected timeout %q: %v\n", timeoutRaw, err)
		os.Exit(1)
	}
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		fmt.Printf("net.DialTimeout(%q, %q, %q): %v\n", network, addr, timeout, err)
		os.Exit(1)
	}
	conn.Close()
}

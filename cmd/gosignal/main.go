package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("Waiting for Ctrl+C or `kill -INT %d`...\n", os.Getpid())

	now := time.Now()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	s := <-sig

	fmt.Printf("\nReceived: '[%d] %s'.\nWaited for %s\n", s, s, time.Since(now))
}

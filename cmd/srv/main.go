package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("error: run '%s <port number>'\n", os.Args[0])
		os.Exit(1)
	}
	port := os.Args[1]
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Println("error: port must be a number")
		os.Exit(1)
	}

	fmt.Println(italic(fmt.Sprintf("Listening on port %s.", port)))
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(
				now(),
				"[", r.RemoteAddr, "]",
				bold(r.Method),
				r.RequestURI,
			)
			w.WriteHeader(http.StatusOK)
		},
	)); err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}

func now() string            { return time.Now().Format("15:04:05.000") }
func bold(s string) string   { return "\x1b[1m" + s + "\x1b[0m" }
func italic(s string) string { return "\x1b[3m" + s + "\x1b[0m" }

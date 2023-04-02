package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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

	fmt.Printf("Listening on port %s.\n", port)
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method, r.RequestURI)
			w.WriteHeader(http.StatusOK)
		},
	)); err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 8888.")
	if err := http.ListenAndServe(":8888", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method, r.RequestURI)
			w.WriteHeader(http.StatusOK)
		},
	)); err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}

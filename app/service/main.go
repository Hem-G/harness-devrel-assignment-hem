package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from myservice\n")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok\n")
	})

	fmt.Println("🚀 myservice started on :8080") // log line for visibility
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server failed:", err)
	}
}

package main

import (
    "fmt"
    "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sample")
	})
	http.ListenAndServe(":5001", nil)
}

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}

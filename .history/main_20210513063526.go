package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>我正在学Golang。</p>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8001", nil)
}

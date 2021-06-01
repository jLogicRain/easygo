package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>我正在学Golang。</p>")
	r.URL.Query()

	fmt.Println("setup: ", &r)
}

func testPointer() *bool {
	setup := true
	return &setup
}

func main() {
	setup := testPointer()
	fmt.Println("setup: ", setup)
	fmt.Println("*setup: ", *setup)

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8001", nil)
}

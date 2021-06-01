package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<p>我正在学Golang。</p>")
	} else if r.URL.Path == "/home" {
		fmt.Fprint(w, "这里是首页")
	} else {
		fmt.Fprint(w, "<p>页面丢失不见了")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8001", nil)
}

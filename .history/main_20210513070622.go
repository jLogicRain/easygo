package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>我正在学Golang。</p>")
	fmt.Fprint(w, "请求路径为："+r.URL.Path)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8001", nil)
}

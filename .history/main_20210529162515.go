package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"-
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<p>我正在学Golang。</p>")
	} else if r.URL.Path == "/home" {
		fmt.Fprint(w, "这里是首页")
	} else if r.URL.Path == "/wxapi" {
		fmt.Fprint(w, "<a href='https://developers.weixin.qq.com/miniprogram/dev/framework/'>这里是微信接口地址</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<p style='color:red;font-size:30px;text-align:center;'>页面丢失不见了</p>")
	}
}

func main() {
	router := mux.NewRouter()

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8001", nil)
}

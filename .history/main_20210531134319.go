package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// 定义首页
func homeHander(w http.ResponseWriter, r *http.Request) {
	// ---  4. 读取成功，显示文章 ---

	// 4.0 设置模板相对路径
	viewDir := "resources/views"

	// 4.1 所有布局模板文件 Slice
	files, _ := filepath.Glob(viewDir + "/layouts/*.gohtml")

	// 4.2 在 Slice 里新增我们的目标文件
	newFiles := append(files, viewDir+"/home/index.gohtml")

	// 4.3 解析模板文件
	tmpl, _ := template.New("index.gohtml").ParseFiles(newFiles...)

	data := true
	// 4.4 渲染模板，将所有文章的数据传输进去
	tmpl.ExecuteTemplate(w, "app", data)
	//fmt.Fprint(w, "这是首页")
}

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

	router.HandleFunc("/", homeHander).Methods("GET").Name("home")

	router.PathPrefix("/dist/css").Handler(http.FileServer(http.Dir("./public")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	http.ListenAndServe(":8001", router)
}

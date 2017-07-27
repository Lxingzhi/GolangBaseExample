package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由 访问规则
	http.HandleFunc("/", sayHello)           /* 第一个参数是rootpath   第二个参数是指定接收什么参数运行*/
	err := http.ListenAndServe(":8080", nil) //传入nil  是默认的handler
	if err != nil {
		log.Fatal(err)
	}
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello word.version 1")
}

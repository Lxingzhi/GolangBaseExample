package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request) //什么样的路径调用什么fun

func main() {
	server := http.Server{
		Addr:        ":8080",         //地址
		Handler:     &myHandler{},    //hadnler
		ReadTimeout: 5 * time.Second, //设置超时
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request)) //make
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

/*自己定一个hadler*/
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //路由转发
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello word.version 3")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bye bye.version 3")
}

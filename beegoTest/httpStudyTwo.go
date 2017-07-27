package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/",
		http.StripPrefix("/static/", h)) //去掉前缀
	err = http.ListenAndServe(":8080", mux) //需要把mux传进去
	if err != nil {
		log.Fatal(err)
	}
}

/*自己定一个hadler*/
type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello word.version 2")
}

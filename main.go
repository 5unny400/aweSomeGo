package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("hello.html")
	//错误处理
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed, err:", err)
		return
	}
}

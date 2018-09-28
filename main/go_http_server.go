package main

import (
	"fmt"
	"net/http"

	"gopkg.in/macaron.v1"

	"gohttpserver/config"
	"gohttpserver/routers"
)

func newInstance() *macaron.Macaron {
	m := macaron.New()
	return m
}

func main() {

	m := newInstance()
	routers.Regist(m)
	listenAddr := fmt.Sprintf("0.0.0.0:%d", conf.Httpport)
	fmt.Println(listenAddr)
	fmt.Println("start")
	if err := http.ListenAndServe(listenAddr, m); err != nil {
		fmt.Println("Err:", err)
	}
}

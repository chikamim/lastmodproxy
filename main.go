package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	website := WebSite{"www.example.com", `(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})`, "2006-01-02 15:04", "Asia/Tokyo"}
	config := &Config{[]WebSite{website}}
	hander := NewLastModifiedHandler(NewBoldTimeStore("/tmp/teststore"), config)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	log.Fatalln(http.ListenAndServe(":8888", proxy))
}

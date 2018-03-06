package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	docwiki := WebSite{"apps.fujisan.co.jp/docwiki", `最終更新: (\d{4}/\d{2}/\d{2} \d{2}:\d{2})`, "2006/01/02 15:04", "Asia/Tokyo"}
	redmine := WebSite{"apps.fujisan.co.jp/redmine", `title="(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})">([^<]+)</a>前に更新.`, "2006-01-02 15:04", "Asia/Tokyo"}

	config := &Config{[]WebSite{docwiki, redmine}}
	hander := NewLastModifiedHandler(NewBoldTimeStore("/tmp/teststore"), config)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	log.Fatalln(http.ListenAndServe(":7777", proxy))
}

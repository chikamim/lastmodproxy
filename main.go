package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	config, err := NewConfig("config.yml")
	if err != nil {
		log.Fatalf("config error - %v\n", err)
	}
	hander := NewLastModifiedHandler(NewBoldTimeStore(config.DBFile), config.Websites)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	log.Fatalln(http.ListenAndServe(":"+config.Port, proxy))
}

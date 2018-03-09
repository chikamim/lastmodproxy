package main

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

func StartLastModifiedProxy(config *Config, force bool) error {
	proxy := goproxy.NewProxyHttpServer()
	hander := NewLastModifiedHandler(NewBoldTimeStore(config.DBFile), config.Websites, force)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	return http.ListenAndServe(":"+config.Port, proxy)
}

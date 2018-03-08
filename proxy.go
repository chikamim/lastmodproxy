package main

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

func StartNotModifiedProxy(config *Config) error {
	proxy := goproxy.NewProxyHttpServer()
	hander := NewLastModifiedHandler(NewBoldTimeStore(config.DBFile), config.Websites, false)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	return http.ListenAndServe(":"+config.NotModifiedPort, proxy)
}

func StartCheckModifiedProxy(config *Config) error {
	proxy := goproxy.NewProxyHttpServer()
	hander := NewLastModifiedHandler(NewBoldTimeStore(config.DBFile), config.Websites, true)

	proxy.OnRequest().DoFunc(hander.OnRequest)
	proxy.OnResponse().DoFunc(hander.OnResponse)
	return http.ListenAndServe(":"+config.CheckModifiedPort, proxy)
}

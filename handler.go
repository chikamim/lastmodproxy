package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elazarl/goproxy"
)

type LastModifiedHandler struct {
	Index *Index
}

func NewLastModifiedHandler(store TimeStorer, config *Config) *LastModifiedHandler {
	index := NewIndex(store, config)
	return &LastModifiedHandler{index}
}

func (h *LastModifiedHandler) OnRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	url := req.URL.String()
	lastModified, err := h.Index.GetLastModified(url)
	if err == nil {
		log.Printf("Index GetLastModified - %v\n", lastModified)
		return req, goproxy.NewResponse(req,
			goproxy.ContentTypeText, http.StatusNotModified,
			"304")
	}
	return req, nil
}

func (h *LastModifiedHandler) OnResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			return resp
	}

	url := ctx.Req.URL.String()
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	io.Copy(writer, resp.Body)
	resp.Body = ioutil.NopCloser(bufio.NewReader(&buf))

	lastModified, err := h.Index.SetLastModified(url, buf.Bytes())
	if err != nil {
		log.Printf("Index SetLastModified error - %v\n", err)
	}

	origin := resp.Header.Get("Last-Modified")
	if len(origin) == 0 {
		log.Printf("Set New LastModified - %v - %v", url, lastModified.Format(time.RFC1123))
		resp.Header.Set("Last-Modified", lastModified.Format(time.RFC1123))
	}
	return resp
}

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

func NewLastModifiedHandler(store TimeStorer, websites []WebSite) *LastModifiedHandler {
	index := NewIndex(store, websites)
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
	log.Printf("Request: %v\n", url)
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	io.Copy(writer, resp.Body)
	bin := bytes.Replace(buf.Bytes(), []byte("nofollow"), []byte(""), -1)
	resp.Body = ioutil.NopCloser(bytes.NewReader(bin))

	lastModified, err := h.Index.SetLastModified(url, buf.Bytes())
	if err != nil {
		log.Printf("SetLastModified error: %v\n", err)
		return resp
	}

	if lastModified.Equal(time.Time{}) {
		log.Println("LastModified Not Found")
		return resp
	}

	log.Printf("LastModified Found: %v\n", lastModified.Format(time.RFC1123))
	resp.Header.Set("Last-Modified", lastModified.Format(time.RFC1123))
	resp.Header.Set("Cache-Control", "private")
	return resp
}

package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Index struct {
	Store    TimeStorer
	Websites []WebSite
}

func NewIndex(store TimeStorer, websites []WebSite) *Index {
	return &Index{store, websites}
}

func urlhash(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
}

func (i *Index) SetLastModified(url string, body []byte) (lastModified time.Time, err error) {
	website := i.MatchedWebsites(url)
	if website == nil {
		return time.Time{}, errors.New("no website config found")
	}

	re, err := regexp.Compile(website.DateMatch)
	if err != nil {
		return time.Time{}, err
	}
	lastModified, err = FindTime(body, re, website.DateLayout, time.Local)
	if err != nil {
		return time.Time{}, err
	}

	return lastModified, i.Store.Set(urlhash(url), lastModified)
}

func (i *Index) MatchedWebsites(url string) *WebSite {
	for _, website := range i.Websites {
		if strings.Contains(url, website.URLFilter) {
			return &website
		}
	}
	return nil
}

func (i *Index) GetLastModified(url string) (time.Time, error) {
	return i.Store.Get(urlhash(url))
}

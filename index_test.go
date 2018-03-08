package main

import (
	"fmt"
)

func ExampleIndex() {
	store := NewBoldTimeStore("/tmp/teststore")

	website := WebSite{"www.example.com", `(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})`, "2006-01-02 15:04", "Asia/Tokyo"}
	index := NewIndex(store, []WebSite{website}, false)
	body := `Title
Edited at 2013-12-10 12:12
Body
`

	index.SetLastModified("http://www.example.com/1", []byte(body))
	lm, _ := index.GetLastModified("http://www.example.com/1")
	fmt.Println(lm)
	// Output: 2013-12-10 12:12:00 +0900 JST
}

func ExampleIndexForceChek() {
	store := NewBoldTimeStore("/tmp/teststore")

	website := WebSite{"www.example.com", `(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})`, "2006-01-02 15:04", "Asia/Tokyo"}
	index := NewIndex(store, []WebSite{website}, true)
	body := `Title
Edited at 2013-12-10 12:12
Body
`

	index.SetLastModified("http://www.example.com/1", []byte(body))
	_, err := index.GetLastModified("http://www.example.com/1")
	fmt.Printf("%v", err)
	// Output: last modified force check
}

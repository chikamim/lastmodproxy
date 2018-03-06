package main

import (
	"regexp"
	"testing"
	"time"
)

func TestFindLastModified(t *testing.T) {
	body := `Title
Edited at 2013-12-10 12:12
posted at 2013-12-09 12:45
Body
`
	re := regexp.MustCompile(`Edited at (\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})`)
	actual, err := FindTime([]byte(body), re, "2006-01-02 15:04", time.Local)
	if err != nil {
		t.Errorf("error occurred - %v", err)
	}
	expected := time.Date(2013, time.December, 10, 12, 12, 0, 0, time.Local)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFindLastModifiedErrorNotFound(t *testing.T) {
	body := `Title
Posted at 2013-12-10 12:12
Body
`
	re := regexp.MustCompile(`Edited at (\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})`)
	_, err := FindTime([]byte(body), re, "2006-01-02 15:04", time.Local)
	if err != ErrTimeNotFound {
		t.Errorf("%v should occurred", ErrTimeNotFound)
	}
}

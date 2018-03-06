package main

import (
	"testing"
	"time"
)

func TestBoldStore(t *testing.T) {
	store := NewBoldTimeStore("/tmp/boltstoretest")
	testdate := time.Date(2013, time.December, 10, 12, 12, 0, 0, time.Local)
	err := store.Set("k1", testdate)
	if err != nil {
		t.Errorf("error occurred - %v", err)
	}

	actual, err := store.Get("k1")
	if err != nil {
		t.Errorf("error occurred - %v", err)
	}
	expected := testdate
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	err = store.Delete("k1")
	if err != nil {
		t.Errorf("error occurred - %v", err)
	}
	_, err = store.Get("k1")
	if err == nil {
		t.Error("error should occurred")
	}
}

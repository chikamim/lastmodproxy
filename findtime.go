package main

import (
	"errors"
	"regexp"
	"time"
)

var ErrTimeNotFound = errors.New("time not found")

func FindTime(b []byte, re *regexp.Regexp, df string, loc *time.Location) (time.Time, error) {
	matches := re.FindSubmatch(b)
	if len(matches) < 2 {
		return time.Now(), ErrTimeNotFound
	}
	return time.ParseInLocation(df, string(matches[1]), loc)
}

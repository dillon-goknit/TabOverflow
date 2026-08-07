package main

import "time"

type Link struct {
	URL   string
	Title string
	Tags  []string
	Added time.Time
	Read  bool
}

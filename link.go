package main

import (
	"net/url"
	"strings"
	"time"
)

type Link struct {
	URL   string
	Title string
	Tags  []string
	Added time.Time
	Read  bool
}

func guessType(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "article"
	}

	host := u.Host

	switch {
	case strings.Contains(host, "youtube.com"):
		return "video"
	case strings.Contains(host, "github.com"):
		return "repo"
	case strings.Contains(host, "imdb.com"), strings.Contains(host, "rottentomatoes.com"):
		return "movie"
	case strings.Contains(host, "goodreads.com"):
		return "book"
	case strings.Contains(host, "x.com"):
		return "twitter link"
	default:
		return "article"
	}
}

func normalizeURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	u.Fragment = ""
	u.Host = strings.ToLower(u.Host)

	if u.Path != "/" && strings.HasSuffix(u.Path, "/") {
		u.Path = strings.TrimSuffix(u.Path, "/")
	}

	trackingParameters := []string{
		"utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content", "fbclid", "gclid", "ref",
	}
	q := u.Query()
	for _, key := range trackingParameters {
		q.Del(key)
	}
	u.RawQuery = q.Encode()

	return u.String()
}

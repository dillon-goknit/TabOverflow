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

func hostMatches(host, domain string) bool {
	return host == domain || strings.HasSuffix(host, "."+domain)
}

func guessType(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "article"
	}

	host := strings.ToLower(u.Hostname())
	host = strings.TrimPrefix(host, "www.")

	switch {
	case hostMatches(host, "youtube.com"), hostMatches(host, "youtu.be"):
		return "video"
	case hostMatches(host, "github.com"):
		return "repo"
	case hostMatches(host, "imdb.com"), hostMatches(host, "rottentomatoes.com"):
		return "movie"
	case hostMatches(host, "goodreads.com"):
		return "book"
	case hostMatches(host, "x.com"), hostMatches(host, "twitter.com"):
		return "twitterlink"
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

package main

import "testing"

func TestGuessType(t *testing.T) {
	cases := []struct {
		name string
		url  string
		want string
	}{
		{"youtube full", "https://www.youtube.com/watch?v=abc", "video"},
		{"github repo", "https://github.com/foo/bar", "repo"},
		{"imdb", "https://www.imdb.com/title/tt123", "movie"},
		{"goodreads", "https://www.goodreads.com/book/show/1", "book"},
		{"unknown domain", "https://some-random-blog.com/post", "article"},
	}

	for _, c := range cases {
		got := guessType(c.url)
		if got != c.want {
			t.Errorf("%s: guessType(%q) = %q, want %q", c.name, c.url, got, c.want)
		}
	}
}

func TestNormalizeURL(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			"strips utm params",
			"https://example.com/post?utm_source=twitter&utm_medium=social",
			"https://example.com/post",
		},
		{
			"lowercases host, keeps path case",
			"https://YouTube.com/watch?v=abc",
			"https://youtube.com/watch?v=abc",
		},
		{
			"strips trailing slash",
			"https://example.com/foo/",
			"https://example.com/foo",
		},
		{
			"drops fragment",
			"https://example.com/post#intro",
			"https://example.com/post",
		},
		{
			"keeps relevant param, strips tracking",
			"https://youtube.com/watch?v=abc&utm_source=x",
			"https://youtube.com/watch?v=abc",
		},
		{
			"strips fbclid and gclid",
			"https://example.com/article?fbclid=123&gclid=456",
			"https://example.com/article",
		},
	}

	for _, c := range cases {
		got := normalizeURL(c.in)
		if got != c.want {
			t.Errorf("%s: normalizeURL(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}

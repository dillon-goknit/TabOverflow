package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCmdAdd(t *testing.T) {
	useTempStore(t)

	if err := cmdAdd("https://example.com/article"); err != nil {
		t.Fatalf("cmdAdd failed: %v", err)
	}

	links, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(links) != 1 {
		t.Fatalf("got %d links, want 1", len(links))
	}

	if err := cmdAdd("https://example.com/article"); err == nil {
		t.Errorf("expected error on duplicate, got nil")
	}
}

func TestCmdAddKeepsRawURL(t *testing.T) {
	useTempStore(t)

	raw := "https://example.com/post?utm_source=twitter&id=42"

	if err := cmdAdd(raw); err != nil {
		t.Fatalf("cmdAdd failed: %v", err)
	}

	links, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if links[0].URL != raw {
		t.Errorf("stored URL = %q, want raw %q", links[0].URL, raw)
	}
}

func TestCmdAddTrimsInput(t *testing.T) {
	useTempStore(t)

	raw := "https://example.com/post"

	if err := cmdAdd("  " + raw + "  "); err != nil {
		t.Fatalf("cmdAdd failed: %v", err)
	}

	links, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if links[0].URL != raw {
		t.Errorf("stored URL = %q, want trimmed %q", links[0].URL, raw)
	}
}

func TestCmdAddDedupesStoredTrackingURL(t *testing.T) {
	useTempStore(t)

	if err := cmdAdd("https://example.com/article?utm_source=x"); err != nil {
		t.Fatalf("first add: %v", err)
	}
	if err := cmdAdd("https://example.com/article"); err == nil {
		t.Errorf("expected error on duplicate of a stored tracking URL, got nil")
	}
}

func TestCmdAddRejectsEmpty(t *testing.T) {
	useTempStore(t)

	if err := cmdAdd("   "); err == nil {
		t.Errorf("expected error on empty URL, got nil")
	}

	links, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(links) != 0 {
		t.Errorf("got %d links, want 0", len(links))
	}
}

func TestCmdListEmpty(t *testing.T) {
	useTempStore(t)

	var buf bytes.Buffer
	if err := cmdList(&buf); err != nil {
		t.Fatalf("cmdList failed: %v", err)
	}

	got := buf.String()
	if !strings.Contains(got, "no links") {
		t.Errorf("output = %q, want it to mention an empty pile", got)
	}
}

func TestCmdListShowsLinks(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/post", Tags: []string{"article"}},
		{URL: "https://github.com/foo/bar", Tags: []string{"repo"}},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdList(&buf); err != nil {
		t.Fatalf("cmdList failed: %v", err)
	}

	got := buf.String()
	for _, want := range []string{
		"https://example.com/post [article]",
		"https://github.com/foo/bar [repo]",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output = %q, want it to contain %q", got, want)
		}
	}
}

func TestCmdListHidesReadLinks(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/unread", Tags: []string{"article"}},
		{URL: "https://example.com/done", Tags: []string{"article"}, Read: true},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdList(&buf); err != nil {
		t.Fatalf("cmdList failed: %v", err)
	}

	got := buf.String()
	if !strings.Contains(got, "https://example.com/unread") {
		t.Errorf("output = %q, want it to contain the unread link", got)
	}
	if strings.Contains(got, "https://example.com/done") {
		t.Errorf("output = %q, want it to omit the read link", got)
	}
}

func TestCmdListAllRead(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/done", Tags: []string{"article"}, Read: true},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdList(&buf); err != nil {
		t.Fatalf("cmdList failed: %v", err)
	}

	got := buf.String()
	if !strings.Contains(got, "no links") {
		t.Errorf("output = %q, want it to mention an empty pile", got)
	}
}

func TestCmdListUntaggedLink(t *testing.T) {
	useTempStore(t)

	links := []Link{{URL: "https://example.com/untagged"}}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdList(&buf); err != nil {
		t.Fatalf("cmdList failed: %v", err)
	}

	got := buf.String()
	if !strings.Contains(got, "https://example.com/untagged") {
		t.Errorf("output = %q, want the URL", got)
	}
	if strings.Contains(got, "[") {
		t.Errorf("output = %q, want no tag brackets", got)
	}
}

func TestCmdDoneMarksRead(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/post", Tags: []string{"article"}},
		{URL: "https://github.com/foo/bar", Tags: []string{"repo"}},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdDone(&buf, "https://example.com/post"); err != nil {
		t.Fatalf("cmdDone failed: %v", err)
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if !got[0].Read {
		t.Errorf("first link Read = false, want true")
	}
	if got[1].Read {
		t.Errorf("second link Read = true, want false")
	}
}

func TestCmdDoneMatchesNormalized(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/stored-dirty?utm_source=x"},
		{URL: "https://example.com/stored-clean"},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdDone(&buf, "https://example.com/stored-dirty"); err != nil {
		t.Fatalf("cmdDone on clean form failed: %v", err)
	}
	if err := cmdDone(&buf, "https://example.com/stored-clean?utm_source=x"); err != nil {
		t.Fatalf("cmdDone on dirty form failed: %v", err)
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	for i, link := range got {
		if !link.Read {
			t.Errorf("link %d (%s) Read = false, want true", i, link.URL)
		}
	}
}

func TestCmdDoneUnknownURL(t *testing.T) {
	useTempStore(t)

	links := []Link{{URL: "https://example.com/post"}}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdDone(&buf, "https://example.com/nope"); err == nil {
		t.Errorf("expected error for a URL not in the pile, got nil")
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if got[0].Read {
		t.Errorf("link was marked read despite no match")
	}
}

func TestCmdDoneRejectsEmpty(t *testing.T) {
	useTempStore(t)

	var buf bytes.Buffer
	if err := cmdDone(&buf, "   "); err == nil {
		t.Errorf("expected error on empty URL, got nil")
	}
}

func TestCmdRmRemovesLink(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/post", Tags: []string{"article"}},
		{URL: "https://github.com/foo/bar", Tags: []string{"repo"}},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdRm(&buf, "https://example.com/post"); err != nil {
		t.Fatalf("cmdRm failed: %v", err)
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d links, want 1", len(got))
	}
	if got[0].URL != "https://github.com/foo/bar" {
		t.Errorf("remaining link = %q, want the github one", got[0].URL)
	}
}

func TestCmdRmMatchesNormalized(t *testing.T) {
	useTempStore(t)

	links := []Link{
		{URL: "https://example.com/stored-dirty?utm_source=x"},
		{URL: "https://example.com/stored-clean"},
	}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdRm(&buf, "https://example.com/stored-dirty"); err != nil {
		t.Fatalf("cmdRm on clean form failed: %v", err)
	}
	if err := cmdRm(&buf, "https://example.com/stored-clean?utm_source=x"); err != nil {
		t.Fatalf("cmdRm on dirty form failed: %v", err)
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %d links, want 0: %v", len(got), got)
	}
}

func TestCmdRmUnknownURL(t *testing.T) {
	useTempStore(t)

	links := []Link{{URL: "https://example.com/post"}}
	if err := save(links); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var buf bytes.Buffer
	if err := cmdRm(&buf, "https://example.com/nope"); err == nil {
		t.Errorf("expected error for a URL not in the pile, got nil")
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(got) != 1 {
		t.Errorf("got %d links, want the pile untouched at 1", len(got))
	}
}

func TestCmdRmRejectsEmpty(t *testing.T) {
	useTempStore(t)

	var buf bytes.Buffer
	if err := cmdRm(&buf, "   "); err == nil {
		t.Errorf("expected error on empty URL, got nil")
	}
}

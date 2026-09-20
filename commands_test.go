package main

import (
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

package main

import (
	"testing"
)

func TestCmdAdd(t *testing.T) {
	t.Chdir(t.TempDir())

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

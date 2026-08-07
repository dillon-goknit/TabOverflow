package main

import "testing"

func TestImportLinks(t *testing.T) {
	got, err := importLinks("testdata/sample.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 3 {
		t.Errorf("got %d links, want 3", len(got))
	}
}

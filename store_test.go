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

func TestSaveLoadRoundTrip(t *testing.T) {
	t.Chdir(t.TempDir())

	want := []Link{
		{URL: "https://example.com"},
		{URL: "https://github/foo/bar"},
	}

	if err := save(want); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	got, err := load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("got %d links, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].URL != want[i].URL {
			t.Errorf("link %d: got URL %q, want %q", i, got[i].URL, want[i].URL)
		}
	}
}

func TestLoadMissingFile(t *testing.T) {
	t.Chdir(t.TempDir())

	got, err := load()
	if err != nil {
		t.Fatalf("load on missing file should not error, got: %v", err)
	}

	if len(got) != 0 {
		t.Errorf("got %d links, want 0", len(got))
	}
}

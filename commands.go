package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

func cmdAdd(rawURL string) error {
	raw := strings.TrimSpace(rawURL)
	if raw == "" {
		return errors.New("empty URL")
	}

	key := normalizeURL(raw)

	links, err := load()
	if err != nil {
		return err
	}

	for _, existing := range links {
		if normalizeURL(existing.URL) == key {
			return errors.New("link already exists")
		}
	}

	link := Link{
		URL:   raw,
		Added: time.Now(),
		Tags:  []string{guessType(raw)},
	}

	links = append(links, link)

	return save(links)
}

func cmdList(w io.Writer) error {
	fmt.Println("list not yet implemented")
	return nil
}

func cmdPick(w io.Writer) error {
	fmt.Println("pick not yet implemented")
	return nil
}

func cmdDone(w io.Writer, url string) error {
	fmt.Println("done not yet implemented")
	return nil
}

func cmdRm(w io.Writer, url string) error {
	fmt.Println("rm not yet implemented")
	return nil
}

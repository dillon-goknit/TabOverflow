package main

import (
	"errors"
	"fmt"
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

func cmdList() {
	fmt.Println("list not yet implemented")
}

func cmdPick() {
	fmt.Println("pick not yet implemented")
}

func cmdDone(url string) {
	fmt.Println("done not yet implemented")
}

func cmdRm(url string) {
	fmt.Println("rm not yet implemented")
}

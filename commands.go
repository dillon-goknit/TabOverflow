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

func cmdList() error {
	fmt.Println("list not yet implemented")
	return nil
}

func cmdPick() error {
	fmt.Println("pick not yet implemented")
	return nil
}

func cmdDone(url string) error {
	fmt.Println("done not yet implemented")
	return nil
}

func cmdRm(url string) error {
	fmt.Println("rm not yet implemented")
	return nil
}

package main

import (
	"errors"
	"fmt"
	"time"
)

func cmdAdd(rawURL string) error {
	url := normalizeURL(rawURL)

	links, err := load()
	if err != nil {
		return err
	}

	for _, existing := range links {
		if existing.URL == url {
			return errors.New("link already exists")
		}
	}

	link := Link{
		URL:   url,
		Added: time.Now(),
		Tags:  []string{guessType(url)},
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

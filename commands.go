package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"strings"
	"time"
)

func cmdAdd(w io.Writer, rawURL string) error {
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

	if err := save(links); err != nil {
		return err
	}

	fmt.Fprintf(w, "added: %s [%s]\n", link.URL, link.Tags[0])
	return nil
}

func cmdList(w io.Writer) error {
	links, err := load()
	if err != nil {
		return err
	}

	var unread []Link
	for _, link := range links {
		if !link.Read {
			unread = append(unread, link)
		}
	}

	if len(unread) == 0 {
		fmt.Fprintln(w, "no links")
		return nil
	}

	for _, link := range unread {
		if len(link.Tags) == 0 {
			fmt.Fprintln(w, link.URL)
		} else {
			fmt.Fprintf(w, "%s [%s]\n", link.URL, link.Tags[0])
		}
	}
	return nil
}

func cmdPick(w io.Writer) error {
	links, err := load()
	if err != nil {
		return err
	}

	var unread []Link
	for _, link := range links {
		if !link.Read {
			unread = append(unread, link)
		}
	}

	if len(unread) == 0 {
		return errors.New("nothing to pick")
	}

	choice := unread[rand.IntN(len(unread))]
	if len(choice.Tags) == 0 {
		fmt.Fprintln(w, choice.URL)
	} else {
		fmt.Fprintf(w, "%s [%s]\n", choice.URL, choice.Tags[0])
	}

	return nil
}

func cmdDone(w io.Writer, url string) error {
	trimmed := strings.TrimSpace(url)
	if trimmed == "" {
		return errors.New("empty URL")
	}

	key := normalizeURL(trimmed)
	links, err := load()
	if err != nil {
		return err
	}

	found := false
	for i := range links {
		if normalizeURL(links[i].URL) == key {
			links[i].Read = true
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("not in pile: %s", trimmed)
	}
	fmt.Fprintf(w, "marked as read: %s\n", trimmed)
	return save(links)
}

func cmdRm(w io.Writer, url string) error {
	trimmed := strings.TrimSpace(url)
	if trimmed == "" {
		return errors.New("empty URL")
	}

	key := normalizeURL(trimmed)
	links, err := load()
	if err != nil {
		return err
	}

	var kept []Link
	found := false
	for _, link := range links {
		if normalizeURL(link.URL) == key {
			found = true
			continue
		}
		kept = append(kept, link)
	}
	if !found {
		return fmt.Errorf("not in pile: %s", trimmed)
	}
	fmt.Fprintf(w, "removed: %s\n", trimmed)
	return save(kept)
}

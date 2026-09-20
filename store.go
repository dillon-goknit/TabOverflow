package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

// import the links from docs/tabs.txt and convert to []Link structs. URL only item populated in struct
func importLinks(path string) ([]Link, error) {
	tabs, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer tabs.Close()

	content, err := io.ReadAll(tabs)
	if err != nil {
		return nil, err
	}

	var links []Link
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		links = append(links, Link{URL: line})
	}

	return links, nil
}

// save writes the full link slice to links.json, overwriting it
func save(links []Link) error {
	data, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("links.json", data, 0o644)
}

func load() ([]Link, error) {
	data, err := os.ReadFile("links.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Link{}, nil
		}
		return nil, err
	}

	var links []Link

	err = json.Unmarshal(data, &links)
	if err != nil {
		return nil, err
	}
	return links, nil
}

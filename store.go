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

// one time function to turn tabs.txt file to json
func save(links []Link) error {
	data, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("links.json", data, 0o644)
}

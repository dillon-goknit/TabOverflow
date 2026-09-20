package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

func dataPath() (string, error) {
	if p := os.Getenv("TABOVERFLOW_DATA"); p != "" {
		return p, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "taboverflow", "links.json"), nil
}

func save(links []Link) error {
	path, err := dataPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func load() ([]Link, error) {
	path, err := dataPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
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

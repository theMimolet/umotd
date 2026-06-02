package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Commands       []Command `json:"commands"`
	InfoFile       string    `json:"info-file"`
	Links          []Link    `json:"links"`
	Prefix         string    `json:"prefix"`
	Suffix         string    `json:"suffix"`
	Tips           []string  `json:"tips"`
	TipsPresets    []string  `json:"tips-presets"`
	UseAccentColor bool      `json:"use-accent-color"`
}

type Command struct {
	Cmd  string `json:"cmd"`
	Desc string `json:"desc"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// sensible defaults
func defaultConfig() Config {
	return Config{
		Commands: []Command{
			{Cmd: "umotd toggle", Desc: "motd_toggle"},
			{Cmd: "fastfetch", Desc: "sys_info"},
			{Cmd: "brew help", Desc: "cli_pkg"},
		},
		Links: []Link{
			{Name: "discuss", URL: "https://universal-blue.discourse.group/"},
			{Name: "discord", URL: "https://discord.com/invite/8RZGC3uFzA"},
			{Name: "mastodon", URL: "https://fosstodon.org/@UniversalBlue"},
		},
		Suffix: " !",
		TipsPresets: []string{
			"default",
		},
	}
}

func writeDefaultConfig(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(defaultConfig(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func getConfig(filepaths ...string) Config {
	cfg := Config{}

	if len(filepaths) == 0 {
		filepaths = []string{
			os.ExpandEnv("$HOME/.config/umotd/config.json"),
			"/etc/umotd/config.json",
		}
	}

	for _, path := range filepaths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			continue
		}
		return cfg
	}

	return defaultConfig()
}

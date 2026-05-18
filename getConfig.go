package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Commands    []Command `json:"commands"`
	InfoFile    string    `json:"info_file"`
	Links       []Link    `json:"links"`
	StyleFile   string    `json:"style_file"`
	TipsPresets []string  `json:"tips_presets"`
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
		InfoFile: "/usr/share/ublue-os/image-info.json",
		Commands: []Command{
			{Cmd: "ujust --choose", Desc: "cmd_list"},
			{Cmd: "ujust toggle-user-motd", Desc: "motd_toggle"},
			{Cmd: "fastfetch", Desc: "fastfetch"},
			{Cmd: "brew help", Desc: "cli_pkg"},
		},
		Links: []Link{
			{Name: "issues", URL: "https://issues.bazzite.gg/"},
			{Name: "docs", URL: "https://docs.bazzite.gg/"},
			{Name: "discord", URL: "https://discord.gg/bazzite"},
			{Name: "bluesky", URL: "https://bluesky.bazzite.gg/"},
		},
		TipsPresets: []string{
			"aurora", "kde", "ublue", "dev",
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
	cfg := defaultConfig()

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
		break
	}

	return cfg
}

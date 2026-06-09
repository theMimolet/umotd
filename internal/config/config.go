package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/leonelquinteros/gotext"
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

// defaultConfig returns a sensible default config
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

func WriteDefaultConfig(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(defaultConfig(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func AddTipsPreset(newPreset string, l *gotext.Locale) error {
	cfg := GetConfig()
	cfg.TipsPresets = append(cfg.TipsPresets, newPreset)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetPath(), data, 0644)
}

func RemoveTipsPreset(presetToRemove string, l *gotext.Locale) error {
	// if strings.Contains(GetPath(), "/etc/") {
	// 	return nil
	// }
	cfg := GetConfig()
	for i, preset := range cfg.TipsPresets {
		if preset == presetToRemove {
			cfg.TipsPresets = append(cfg.TipsPresets[:i], cfg.TipsPresets[i+1:]...)
			break
		}
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetPath(), data, 0644)
}

func ListTipsPresets() []string {
	cfg := GetConfig()
	return cfg.TipsPresets
}

// isConfigOkay checks if the config file at the given path is valid
func isConfigOkay(path string) bool {

	var noError bool = true

	_, err := os.Stat(path)
	if err != nil {
		noError = false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		noError = false
	}

	var tempCfg Config
	if err := json.Unmarshal(data, &tempCfg); err != nil {
		noError = false
	}

	return noError
}

// GetConfigPath returns the path to a valid config file, returns "" if no valid config file is found
func GetPath() string {
	filepath := []string{
		os.ExpandEnv("$HOME/.config/umotd/config.json"),
		"/etc/umotd/config.json",
	}
	for _, path := range filepath {
		if isConfigOkay(path) {
			return path
		}
	}
	return ""
}

// GetConfig returns the config file at the given path, or a default config if no valid config file is found
func GetConfig() Config {
	cfg := Config{}
	path := GetPath()

	if path == "" {
		return defaultConfig()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return defaultConfig()
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return defaultConfig()
	}

	return cfg
}

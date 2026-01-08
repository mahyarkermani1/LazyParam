package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	X8Settings struct {
		Delay      int      `json:"delay"`
		Headers    string   `json:"headers"`
		Methods    []string `json:"methods"`
		ChainCount int      `json:"chain_count"`
	} `json:"x8_settings"`
	Wordlists struct {
		CustomWordlistPath string `json:"custom_wordlist_path"`
	} `json:"wordlists"`
	HTTPSettings struct {
		UserAgent string `json:"user_agent"`
	} `json:"http_settings"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	if len(cfg.X8Settings.Methods) == 0 {
		return nil, fmt.Errorf("at least one method (GET/POST) must be specified in config.json")
	}
	if cfg.Wordlists.CustomWordlistPath == "" {
		return nil, fmt.Errorf("custom_wordlist_path cannot be empty")
	}
	if cfg.HTTPSettings.UserAgent == "" {
		cfg.HTTPSettings.UserAgent = "Mozilla/5.0 (compatible; LazyParam/1.0)"
	}

	return &cfg, nil
}

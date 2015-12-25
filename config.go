package main

import (
	"os"
	"path"
)

type Config struct {
	OAuth struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"oauth"`
}

func loadConfig() (*Config, error) {
	configDir := path.Join(os.Getenv("HOME"), ".config", Name)

	return nil, nil
}

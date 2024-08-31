package config

import (
	"encoding/json"
	"errors"
	"os"
)

type (
	Config struct {
		DB   []Database `json:"database"`
		Port string     `json:"port"`
	}

	Database struct {
		Name             string
		DBType           string
		ConnectionString string
		Username         string
		Password         string
	}
)

var config *Config

func Init() error {
	fileName := "config/config.json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return errors.New("Config not found")
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return errors.New("Error when read config")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return errors.New("Error when marshal config")
	}
	return nil
}

func Get() Config {
	if config == nil {
		return Config{}
	}
	return *config
}

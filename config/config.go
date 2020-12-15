package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Password string `json:"password"`
	BotToken string `json:"botToken"`
	ChatID float64 `json:"chatID"`
	Computers []Computer `json:"computers"`
}

type Computer struct {
	Name string `json:"name"`
	Mac string `json:"mac"`
	IP *string `json:"broadcast"`
	Password *string `json:"password"`
}

func ReadConfig(path string) (config *Config, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to open JSON file: %s", path))
	}
	defer jsonFile.Close()

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var output Config
	err = json.Unmarshal(content, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

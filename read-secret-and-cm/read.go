package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func readMessage(path string) (string, error) {
	msg, err := readTomlFile(path)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func readTomlFile(path string) (string, error) {
	file, err := toml.LoadFile(path)
	if err != nil {
		return "", fmt.Errorf("load toml file: %w", err)
	}

	msg := file.Get("message.text").(string)
	return msg, nil
}

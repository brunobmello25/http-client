package models

import (
	"encoding/json"
	"io"
	"os"
)

type Request struct {
	Name        string            `json:"name"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	Description string            `json:"description"`
}

type Collection struct {
	Name     string    `json:"name"`
	Requests []Request `json:"requests"`
}

func LoadCollection(path string) (*Collection, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var collection Collection
	if err := json.Unmarshal(data, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}

func SaveCollection(collection *Collection, path string) error {
	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

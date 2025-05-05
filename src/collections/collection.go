package collections

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Collection represents a group of related HTTP requests
type Collection struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Requests    []*Request `json:"requests"`
}

// LoadCollection loads a collection from a JSON file
func LoadCollection(path string) (*Collection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var collection Collection
	if err := json.Unmarshal(data, &collection); err != nil {
		return nil, err
	}

	return &collection, nil
}

// Save saves the collection to a JSON file
func (c *Collection) Save(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// AddRequest adds a new request to the collection
func (c *Collection) AddRequest(req *Request) {
	c.Requests = append(c.Requests, req)
}

// RemoveRequest removes a request from the collection by name
func (c *Collection) RemoveRequest(name string) {
	for i, req := range c.Requests {
		if req.Name == name {
			c.Requests = append(c.Requests[:i], c.Requests[i+1:]...)
			return
		}
	}
} 
package collections

import (
	"net/http"
	"time"
)

// Request represents an HTTP request in our collection
type Request struct {
	Name        string            `json:"name"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	Description string            `json:"description"`
}

// Response represents an HTTP response
type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Duration   time.Duration     `json:"duration"`
}

// Execute sends the HTTP request and returns the response
func (r *Request) Execute() (*Response, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: Read response body
	// TODO: Parse response headers

	return &Response{
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
	}, nil
} 
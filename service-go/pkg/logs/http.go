package logs

import (
	"bytes"
	"net/http"
)

type HttpLogWriter struct {
	URL string
}

func (h *HttpLogWriter) Write(p []byte) (n int, err error) {
	// Prepare the HTTP request
	req, err := http.NewRequest("POST", h.URL, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return 0, err
	}

	return len(p), nil
}

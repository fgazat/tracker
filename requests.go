package tracker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) request(
	ctx context.Context,
	method string,
	path string,
	contentType string,
	content []byte,
) (*http.Request, error) {
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(content))
	if err != nil {
		return req, err
	}
	req.Header.Add(hdrAuthorizationKey, "OAuth "+c.token)
	if c.userAgent != "" {
		req.Header.Add(hdrUserAgentKey, c.userAgent)
	}
	return req, nil
}

// Helper function to make HTTP requests
func (c *Client) makeRequest(ctx context.Context, method, apiEndpoint string, payload []byte) (*http.Response, error) {
	url := c.baseURL + apiEndpoint
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set(hdrAuthorizationKey, "OAuth "+c.token)
	req.Header.Set("X-Cloud-Org-ID", c.XCloudOrgID)
	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.userAgent != "" {
		req.Header.Add(hdrUserAgentKey, c.userAgent)
	}
	// Perform the actual HTTP request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Helper function to handle and print responses
func handleResponse(resp *http.Response, successStatusCodes []int) error {
	defer resp.Body.Close()

	// Check if the status code is within the successful status codes
	isSuccess := false
	for _, code := range successStatusCodes {
		if resp.StatusCode == code {
			isSuccess = true
			break
		}
	}

	if isSuccess {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Response: ", string(bodyBytes))
	} else {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

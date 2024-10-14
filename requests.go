package tracker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Helper function to make HTTP requests
func (c *Client) makeRequest(ctx context.Context, method, apiEndpoint string, payload []byte) (*http.Response, error) {
	url := c.baseURL + apiEndpoint
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set(hdrHostKey, c.baseURL)
	req.Header.Set(hdrAuthorizationKey, "OAuth "+c.token)
	if c.xCloudOrgID != "" {
		req.Header.Set(hdrXCloudOrgIDKey, c.xCloudOrgID)
	}
	if c.xOrgID != "" {
		req.Header.Set(hdrXCloudOrgIDKey, c.xOrgID)
	}
	if c.userAgent != "" {
		req.Header.Add(hdrUserAgentKey, c.userAgent)
	}
	if method == methodPost || method == methodPut {
		req.Header.Set(hdrContentType, jsonContentType)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Helper function to handle and print responses
func handleResponse(resp *http.Response, result any) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return err
		}
	}
	return nil
}

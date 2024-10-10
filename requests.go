package tracker

import (
	"bytes"
	"context"
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

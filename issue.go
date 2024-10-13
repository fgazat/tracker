package tracker

import (
	"context"
	"fmt"
	"log"
	"net/url"
)

// Gets Issue by key
// https://yandex.cloud/en/docs/tracker/concepts/issues/get-issue
func (c *Client) GetIssueByKey(ctx context.Context, key string, result any) error {
	url := fmt.Sprintf("/v2/issues/%s", key)
	resp, err := c.makeRequest(ctx, methodGet, url, nil)
	if err != nil {
		return err
	}
	return handleResponse(resp, result)
}

func (c *Client) GetIssues(ctx context.Context, params url.Values, result any) error {
	url := "/v2/issues/"
	if len(params) != 0 {
		url += "?" + params.Encode()
	}
	log.Println(url)
	resp, err := c.makeRequest(ctx, methodGet, url, nil)
	if err != nil {
		return err
	}
	return handleResponse(resp, result)
}

func (c *Client) GetIssuesByQuery(ctx context.Context, query string, result any) error {
	return c.GetIssues(ctx, url.Values{"query": []string{query}}, result)
}

package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) GetIssueByKey(ctx context.Context, key string, result any) error {
	url := fmt.Sprintf("/issues/%s", key)
	r, err := c.makeRequest(ctx, methodGet, url, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode >= 300 {
		return fmt.Errorf("status code != 200: %d", r.StatusCode)
	}
	if result != nil {
		if err = json.NewDecoder(r.Body).Decode(result); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) GetIssues(ctx context.Context, params url.Values, result any) error {
	return nil
}

func (c *Client) GetIssuesByQuery(ctx context.Context, query string, result any) error {
	return nil
}

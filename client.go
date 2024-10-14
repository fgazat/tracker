package tracker

import (
	"log"
	"net/http"
)

const (
	defaultBaseUrl = "https://api.tracker.yandex.net"

	methodGet     = "GET"
	methodPost    = "POST"
	methodPut     = "PUT"
	methodDelete  = "DELETE"
	methodPatch   = "PATCH"
	methodHead    = "HEAD"
	methodOptions = "OPTIONS"

	hdrHostKey          = "Host"
	hdrAuthorizationKey = "Authorization"
	hdrXCloudOrgIDKey   = "X-Cloud-Org-ID"
	hdrXOrgIDKey        = "X-Org-ID"
	hdrUserAgentKey     = "User-Agent"
	hdrContentType      = "Content-Type"

	plainTextType   = "text/plain; charset=utf-8"
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"
)

type Client struct {
	baseURL     string
	token       string
	client      http.Client
	xCloudOrgID string
	xOrgID      string

	userAgent string
	debug     bool
}

func New(token string, opts ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseUrl,
		token:   token,
		client:  http.Client{},
		debug:   false,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithXCloudOrgID(s string) Option {
	return func(c *Client) {
		c.xCloudOrgID = s
	}
}

func WithXOrgID(s string) Option {
	return func(c *Client) {
		c.xOrgID = s
	}
}

func WithUserAgent(s string) Option {
	return func(c *Client) {
		c.userAgent = s
	}
}

func WithLogger(l *log.Logger) Option {
	return func(c *Client) {
	}
}

func WithDebug(d bool) Option {
	return func(c *Client) {
		c.debug = d
	}
}

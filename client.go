package tracker

import (
	"log"
	"net/http"
)

const (
	defaultBaseUrl = "https://api.tracker.yandex.net"

	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

var (
	hdrUserAgentKey       = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey          = http.CanonicalHeaderKey("Accept")
	hdrContentTypeKey     = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey   = http.CanonicalHeaderKey("Content-Length")
	hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")
	hdrLocationKey        = http.CanonicalHeaderKey("Location")
	hdrAuthorizationKey   = http.CanonicalHeaderKey("Authorization")
	hdrWwwAuthenticateKey = http.CanonicalHeaderKey("WWW-Authenticate")

	plainTextType   = "text/plain; charset=utf-8"
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"
)

type Client struct {
	baseURL string
	token   string
	client  http.Client

	userAgent string
	debug     bool
}

func NewClient(token string, opts ...Option) *Client {
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

func WithRetries(count int) Option {
	return func(c *Client) {
	}
}

func WithRetriesCondition(f func()) Option {
	return func(c *Client) {
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithLogger(l *log.Logger) Option {
	return func(c *Client) {
	}
}

func WithDebug(d bool) Option {
	return func(c *Client) {
		c.debug = true
	}
}

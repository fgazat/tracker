package tracker

import (
	"log"
	"net/http"
)

const (
	defaultBaseUrl = "https://api.tracker.yandex.net/v2"

	methodGet     = "GET"
	methodPost    = "POST"
	methodPut     = "PUT"
	methodDelete  = "DELETE"
	methodPatch   = "PATCH"
	methodHead    = "HEAD"
	methodOptions = "OPTIONS"
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
	baseURL     string
	token       string
	client      http.Client
	XCloudOrgID string
	XOrgID      string

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

func WithXCloudOrgID(s string) Option {
	return func(c *Client) {
		c.XCloudOrgID = s
	}
}

func WithXOrgID(s string) Option {
	return func(c *Client) {
		c.XOrgID = s
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

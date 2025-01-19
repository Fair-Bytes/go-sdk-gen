// Code generated by `go-sdk-gen`. DO NOT EDIT.
package petstore

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//go:embed .version
var version string

const (
	// APIUrl is the URL of our API. Currently, SumUp doesn't provide any
	// other environment or APIs thus APIUrl is used as the default URL
	// for the client.
	APIUrl = "https://api.sumup.com"
)

type service struct {
	client *Client
}

type Client struct {
	// service is the shared service struct re-used for all services.
	svc service

	// client is the HTTP client used to communicate with the API.
	client *http.Client
	// url is the url of the API the requests will be sent to.
	url string
	// userAgent is the user-agent header that will be sent with
	// every request.
	userAgent string
	// key is the API key or access token used for authorization.
	key string

	Pets   *PetsService
	Shared *SharedService
}

// NewClient creates new SumUp API client.
// To use APIs that require authentication use [Client.WithAuth].
func NewClient() *Client {
	c := &Client{
		client:    http.DefaultClient,
		userAgent: fmt.Sprintf("sumup-go/%s", version),
		url:       APIUrl,
	}
	c.populate()
	return c
}

// WithAuth returns a copy of the client configured with the provided Authorization key.
func (c *Client) WithAuth(key string) *Client {
	clone := Client{
		client:    c.client,
		url:       APIUrl,
		userAgent: c.userAgent,
		key:       key,
	}
	clone.populate()
	return &clone
}

// WithClient returns a copy of the client configured with the provided http client.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	clone := Client{
		client:    client,
		url:       APIUrl,
		userAgent: c.userAgent,
		key:       c.key,
	}
	clone.populate()
	return &clone
}

func (c *Client) populate() {
	c.svc.client = c
	c.Pets = (*PetsService)(&c.svc)
	c.Shared = (*SharedService)(&c.svc)
}

func (c *Client) NewRequest(
	ctx context.Context,
	method, path string,
	body io.Reader,
) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.url+path,
		body,
	)
	if err != nil {
		return nil, fmt.Errorf("build request: %s", err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+c.key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("SumUp-Version", version)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

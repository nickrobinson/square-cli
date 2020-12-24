package square

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Client is the API client used to sent requests to Square.
type Client struct {
	// The base URL (protocol + hostname) used for all requests sent by this
	// client.
	BaseURL *url.URL

	// Access token used to authenticate requests sent by this client. If left
	// empty, the `Authorization` header will be omitted.
	AccessToken string

	// When this is enabled, request and response headers will be printed to
	// stdout.
	Verbose bool

	// Cached HTTP client, lazily created the first time the Client is used to
	// send a request.
	httpClient *http.Client
}

// PerformRequest sends a request to Square and returns the response.
func (c *Client) PerformRequest(method, path string, params string, configure func(*http.Request)) (*http.Response, error) {
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	url = c.BaseURL.ResolveReference(url)

	var body io.Reader
	if method == http.MethodPost || method == http.MethodPut {
		body = strings.NewReader(params)
	} else {
		url.RawQuery = params
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Content-Type", "application/json")

	if c.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	}

	if configure != nil {
		configure(req)
	}

	if c.httpClient == nil {
		c.httpClient = newHTTPClient(c.Verbose)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func newHTTPClient(verbose bool) *http.Client {
	httpTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	tr := &verboseTransport{
		Transport: httpTransport,
		Verbose:   verbose,
		Out:       os.Stderr,
	}

	return &http.Client{
		Transport: tr,
	}
}

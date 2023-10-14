package config

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"time"
)

// RetryCount - number of retries the service will retry the request
const RetryCount = 3

func NewRetryableClient(config Config) *http.Client {
	transport := &retryableTransport{
		transport: &http.Transport{},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.Server.Timeout.Write) * time.Second,
	}
}

// By default, the Golang HTTP client will close the request body after a request
// is sent. This can cause issues when retrying requests since the body may have
// already been closed. To prevent this from happening, we can create a custom
// RoundTripper that wraps the default Transport and prevents the request body
// from being closed.

type retryableTransport struct {
	transport http.RoundTripper
}

func (t *retryableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request body
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	// Send the request
	resp, err := t.transport.RoundTrip(req)
	// Retry logic
	retries := 0
	for shouldRetry(err, resp) && retries < RetryCount {
		// Wait for the specified backoff period
		time.Sleep(backoff(retries))
		// We're going to retry, consume any response to reuse the connection.
		drainBody(resp)
		// Clone the request body again
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		log.Infof("Retrying Request %+v", req.Body)
		// Retry the request
		resp, err = t.transport.RoundTrip(req)
		retries++
	}
	// Return the response
	return resp, err
}

// backoff - A backoff strategy is a method for delaying retries after a failed
// request. The idea is to increase the delay between retries to give the server
// time to recover.
func backoff(retries int) time.Duration {
	return time.Duration(math.Pow(2, float64(retries))) * time.Second
}

// shouldRetry - We can also implement retry logic for specific network errors
// and response status codes. For example, if we encounter a network error, we
// can retry the request. Similarly, if we receive a 502, 503, or 504 status
// code, we can retry the request.
func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusBadGateway ||
		resp.StatusCode == http.StatusServiceUnavailable ||
		resp.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}

// drainBody - To reuse the same connection when retrying requests, we need to
// drain the response body before closing the connection.
func drainBody(resp *http.Response) {
	if resp.Body != nil {
		_, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			return
		}
		err = resp.Body.Close()
		if err != nil {
			return
		}
	}
}

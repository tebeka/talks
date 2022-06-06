package main

import (
	"fmt"
	"net/http"
	"testing"
)

// START_API_CLIENT OMIT
type APIClient struct {
	c       http.Client
	baseURL string
}

func (c *APIClient) Health() error {
	url := fmt.Sprintf("%s/health", c.baseURL)
	resp, err := c.c.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}
	return nil
}

// END_API_CLIENT OMIT

// START_TEST OMIT
type errTripper struct {
	code int
}

func (e *errTripper) RoundTrip(*http.Request) (*http.Response, error) { // HL
	return &http.Response{StatusCode: e.code}, nil
}

func TestHealthError(t *testing.T) {
	c := APIClient{}
	c.c.Transport = &errTripper{http.StatusInternalServerError} // HL

	err := c.Health()
	if err == nil {
		t.Fatal(err)
	}
}

// END_TEST OMIT

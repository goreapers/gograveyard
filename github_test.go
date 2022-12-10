package gograveyard

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn http.RoundTripper) *http.Client {
	return &http.Client{ //nolint:exhaustivestruct
		Transport: fn,
	}
}

func TestOpenIssuesCount(t *testing.T) {
	expectedCount := 2

	testFn := func(req *http.Request) *http.Response {
		s := SearchIssues{
			TotalCount: expectedCount,
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{ //nolint:exhaustivestruct
			StatusCode: http.StatusOK,
			// Send response to be tested
			Body: io.NopCloser(bytes.NewBuffer(b)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	}

	client := NewTestClient(RoundTripFunc(testFn))

	p := Project{
		client,
		"goreapers",
		"gograveyard",
	}
	c, err := p.OpenIssuesCount()
	if err != nil {
		t.Fatal(err)
	}

	if c != expectedCount {
		t.Fatalf("Expected issue count %d but got %d", 2, c)
	}
}

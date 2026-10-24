package tachosync

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	httpclient "github.com/rojakc96/tachosync-go/internal/client"
	"github.com/rojakc96/tachosync-go/models"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func testClient(handler func(*http.Request) (int, http.Header, string)) *Client {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		status, headers, body := handler(request)
		return &http.Response{
			StatusCode: status,
			Header:     headers,
			Body:       io.NopCloser(strings.NewReader(body)),
			Request:    request,
		}, nil
	})
	return NewClient("https://api.example.test", "secret", WithHTTPClient(&http.Client{Transport: transport}))
}

func TestListDriversBuildsAuthenticatedRequest(t *testing.T) {
	c := testClient(func(r *http.Request) (int, http.Header, string) {
		if r.URL.Path != DriversEndpoint {
			t.Errorf("path = %q", r.URL.Path)
		}
		if got := r.Header.Get("X-Api-Key"); got != "secret" {
			t.Errorf("X-Api-Key = %q", got)
		}
		query := r.URL.Query()
		if got := query["Statuses"]; len(got) != 2 || got[0] != "Active" || got[1] != "Disabled" {
			t.Errorf("Statuses = %#v", got)
		}
		if query.Get("PageNumber") != "2" || query.Get("PageSize") != "25" {
			t.Errorf("pagination = %q/%q", query.Get("PageNumber"), query.Get("PageSize"))
		}
		return http.StatusOK, make(http.Header), `{"page":2,"pageSize":25,"items":[]}`
	})

	result, err := c.ListDrivers(context.Background(), models.ListDriversParams{
		Statuses:   []string{"Active", "Disabled"},
		PageNumber: 2,
		PageSize:   25,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Page != 2 || result.PageSize != 25 {
		t.Fatalf("result = %#v", result)
	}
}

func TestGetDriverEscapesPathParameter(t *testing.T) {
	c := testClient(func(r *http.Request) (int, http.Header, string) {
		if got := r.URL.EscapedPath(); got != "/v1/Drivers/a%2Fb" {
			t.Errorf("escaped path = %q", got)
		}
		return http.StatusOK, make(http.Header), `{"id":"a/b","company":{"id":"company"},"createdAt":"2026-10-17T00:00:00Z","updatedAt":"2026-10-17T00:00:00Z"}`
	})

	result, err := c.GetDriver(context.Background(), "a/b")
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != "a/b" {
		t.Fatalf("id = %q", result.ID)
	}
}

func TestDownloadAndEmptyResponse(t *testing.T) {
	c := testClient(func(r *http.Request) (int, http.Header, string) {
		switch r.URL.Path {
		case DriverFilesDownloadEndpoint:
			return http.StatusOK, make(http.Header), string([]byte{0, 1, 2, 3})
		case "/v1/Drivers/driver":
			return http.StatusOK, make(http.Header), ""
		default:
			return http.StatusNotFound, make(http.Header), ""
		}
	})

	data, err := c.DownloadDriverFiles(context.Background(), models.DownloadFilesParams{IDs: []string{"file"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 4 || data[3] != 3 {
		t.Fatalf("data = %#v", data)
	}
	if err := c.DeleteDriver(context.Background(), "driver"); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPErrorIsExposed(t *testing.T) {
	c := testClient(func(r *http.Request) (int, http.Header, string) {
		return http.StatusBadRequest, make(http.Header), `{"message":"invalid"}`
	})

	_, err := c.GetCompanies(context.Background())
	var httpErr *httpclient.HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("error = %T, want *client.HTTPError", err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d", httpErr.StatusCode)
	}
}

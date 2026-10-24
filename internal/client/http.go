package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Get builds an endpoint URL, makes a GET request, and decodes the JSON response.
func Get[T any](baseURL, endpointPath string, params ...PathParameter) (*T, error) {
	endpoint, err := BuildEndpoint(baseURL, endpointPath, params...)
	if err != nil {
		return nil, err
	}

	result := new(T)
	if err := HTTPRequest(http.MethodGet, endpoint, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Do sends an HTTP request using the supplied client and headers. JSON responses
// are decoded into result; []byte and io.Writer results receive the raw body.
// A nil result is valid for endpoints with an empty response body.
func Do(ctx context.Context, httpClient *http.Client, method, endpoint string, headers http.Header, body, result any) (*http.Response, error) {
	var reqBody io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	for name, values := range headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	if body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return resp, &HTTPError{StatusCode: resp.StatusCode, Body: bodyBytes}
	}

	if result == nil || resp.StatusCode == http.StatusNoContent {
		return resp, nil
	}

	switch destination := result.(type) {
	case *[]byte:
		*destination, err = io.ReadAll(resp.Body)
	case io.Writer:
		_, err = io.Copy(destination, resp.Body)
	default:
		err = json.NewDecoder(resp.Body).Decode(result)
		if err == io.EOF {
			err = nil
		}
	}
	if err != nil {
		return resp, fmt.Errorf("error decoding response: %w", err)
	}

	return resp, nil
}

// HTTPError describes a non-2xx response and keeps its response body available.
type HTTPError struct {
	StatusCode int
	Body       []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("error HTTP %d: %s", e.StatusCode, string(e.Body))
}

// HTTPRequest is a generic function to make HTTP requests and return a JSON response.
func HTTPRequest[T any](method, url string, body any, result *T) error {
	_, err := Do(context.Background(), http.DefaultClient, method, url, nil, body, result)
	return err
}

// PathParameter represents an endpoint path parameter replacement.
type PathParameter struct {
	name  string
	value string
}

// PathParam creates an endpoint path parameter replacement.
func PathParam(name, value string) PathParameter {
	return PathParameter{name: name, value: value}
}

// BuildEndpoint builds a full URL from a base URL, endpoint path, and path parameters.
func BuildEndpoint(baseURL, endpointPath string, params ...PathParameter) (string, error) {
	path := endpointPath

	for _, param := range params {
		placeholder := fmt.Sprintf("{%s}", param.name)
		path = strings.ReplaceAll(path, placeholder, url.PathEscape(param.value))
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %w", err)
	}

	rawPath := joinEscapedPath(parsedBaseURL.EscapedPath(), path)
	decodedPath, err := url.PathUnescape(rawPath)
	if err != nil {
		return "", fmt.Errorf("error decoding URL path: %w", err)
	}

	parsedBaseURL.Path = decodedPath
	parsedBaseURL.RawPath = rawPath

	return parsedBaseURL.String(), nil
}

func joinEscapedPath(basePath, endpointPath string) string {
	if endpointPath == "/" {
		return strings.TrimRight(basePath, "/") + "/"
	}

	return strings.TrimRight(basePath, "/") + "/" + strings.TrimLeft(endpointPath, "/")
}

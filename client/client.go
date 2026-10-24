package tachosync

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	httpclient "github.com/rojakc96/tachosync-go/internal/client"
)

// Client is an authenticated TachoSync API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

type Option func(*Client)

// WithHTTPClient configures the HTTP transport used by the API client.
func WithHTTPClient(value *http.Client) Option {
	return func(c *Client) {
		if value != nil {
			c.httpClient = value
		}
	}
}

// NewClient creates a client. baseURL should be the API origin, without /v1.
func NewClient(baseURL, apiKey string, options ...Option) *Client {
	c := &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *Client) request(
	ctx context.Context,
	method, endpoint string,
	query, body, result any,
	pathParams ...httpclient.PathParameter,
) (*http.Response, error) {
	fullURL, err := httpclient.BuildEndpoint(c.baseURL, endpoint, pathParams...)
	if err != nil {
		return nil, err
	}
	if query != nil {
		values, err := encodeQuery(query)
		if err != nil {
			return nil, err
		}
		if len(values) > 0 {
			fullURL += "?" + values.Encode()
		}
	}

	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("X-Api-Key", c.apiKey)
	return httpclient.Do(ctx, c.httpClient, method, fullURL, headers, body, result)
}

func encodeQuery(input any) (url.Values, error) {
	values := make(url.Values)
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return values, nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("query parameters must be a struct, got %T", input)
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Tag.Get("url")
		if name == "" || name == "-" {
			continue
		}
		if err := addQueryValue(values, name, v.Field(i)); err != nil {
			return nil, fmt.Errorf("query parameter %s: %w", name, err)
		}
	}
	return values, nil
}

func addQueryValue(values url.Values, name string, value reflect.Value) error {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		for i := 0; i < value.Len(); i++ {
			if err := addQueryValue(values, name, value.Index(i)); err != nil {
				return err
			}
		}
		return nil
	}
	if value.Type() == reflect.TypeOf(time.Time{}) {
		values.Add(name, value.Interface().(time.Time).Format(time.RFC3339))
		return nil
	}
	switch value.Kind() {
	case reflect.String:
		if value.String() != "" {
			values.Add(name, value.String())
		}
	case reflect.Bool:
		values.Add(name, strconv.FormatBool(value.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(name, strconv.FormatInt(value.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		values.Add(name, strconv.FormatUint(value.Uint(), 10))
	default:
		return fmt.Errorf("unsupported type %s", value.Type())
	}
	return nil
}

func locationID(response *http.Response) string {
	if response == nil {
		return ""
	}
	location := response.Header.Get("Location")
	if parsed, err := url.Parse(location); err == nil {
		path := strings.TrimRight(parsed.Path, "/")
		if index := strings.LastIndex(path, "/"); index >= 0 {
			return path[index+1:]
		}
	}
	return location
}

func requestJSON[T any](ctx context.Context, c *Client, method, endpoint string, query, body any, pathParams ...httpclient.PathParameter) (*T, error) {
	result := new(T)
	if _, err := c.request(ctx, method, endpoint, query, body, result, pathParams...); err != nil {
		return nil, err
	}
	return result, nil
}

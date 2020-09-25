package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/r57ty7/pver/service"
)

const (
	EndpointSearch = "/rest/api/2/search"
	EndpointIssues = "/rest/api/2/issue/%s"
)

type JiraRepository struct {
	client  *http.Client
	baseURL *url.URL
	auth    Authentication
}

type Authentication struct {
	username string
	password string
}

// NewJiraRepository create JiraRepository
// if httpClient is not defined, http.DefaultClient is used
func NewJiraRepository(httpClient *http.Client, baseURL string, username, password string) (*JiraRepository, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// ensure baseURL has trailling slash
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	auth := Authentication{username, password}

	return &JiraRepository{
		client:  httpClient,
		baseURL: parsedBaseURL,
		auth:    auth,
	}, nil
}

func (r *JiraRepository) SearchIssues(jql string) (*service.SearchResults, error) {
	u := url.URL{
		Path: EndpointSearch,
	}
	uv := url.Values{}
	if jql != "" {
		uv.Add("jql", jql)
	}

	u.RawQuery = uv.Encode()

	req, err := r.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	searchResults := new(service.SearchResults)
	_, err = r.Do(req, searchResults)
	if err != nil {
		return nil, err
	}
	return searchResults, nil
}

// NewRequestWithContext creates http.Request
func (r *JiraRepository) NewRequestWithContext(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// trim preceding slash since base url has trailling slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := r.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Set authentication information
	req.SetBasicAuth(r.auth.username, r.auth.password)

	return req, nil

}

// NewRequest wraps NewRequestWithContext using the background context.
func (r *JiraRepository) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return r.NewRequestWithContext(context.Background(), method, urlStr, body)
}

// Do sends an API request and returns the API response.
func (r *JiraRepository) Do(req *http.Request, v interface{}) (*http.Response, error) {
	httpResp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}

	return httpResp, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}

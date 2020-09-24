package service

import (
	"context"
	"net/http"
	"net/url"

	"github.com/r57ty7/jiracket/domain"
	"github.com/r57ty7/pver/cmd"
)

type jiraRepository struct {
	client *JiraClient
}

type SearchResults struct {
	Expand     string            `json:"expand,omitempty" yaml:"expand,omitempty"`
	Issues     []domain.Issue    `json:"issues,omitempty" yaml:"issues,omitempty"`
	MaxResults int               `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Names      map[string]string `json:"names,omitempty" yaml:"names,omitempty"`
	StartAt    int               `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total      int               `json:"total,omitempty" yaml:"total,omitempty"`
}

// NewJiraRepository returns domain JiraRepository
func NewJiraRepository(client *JiraClient) cmd.JiraRepository {
	return &jiraRepository{
		client: client,
	}
}

func (r *jiraRepository) Search(ctx context.Context, jql string) ([]domain.Issue, error) {
	u := url.URL{
		Path: EndpointSearch,
	}
	uv := url.Values{}
	if jql != "" {
		uv.Add("jql", jql)
	}

	u.RawQuery = uv.Encode()

	req, err := r.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	searchResults := new(SearchResults)
	_, err = r.client.Do(req, searchResults)
	if err != nil {
		return nil, err
	}

	return searchResults.Issues, nil
}

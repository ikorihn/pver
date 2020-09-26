package cmd

import (
	"testing"

	"github.com/r57ty7/pver/service"
)

// FileVersionManager の mock
type mockFvm struct {
	version string
}

func (m *mockFvm) SetConfig(conf service.Config) {
}

func (m *mockFvm) Version() string {
	return m.version
}

func (m *mockFvm) Update(newVersion string) error {
	m.version = newVersion
	return nil
}

type mockGitRepo struct {
}

func (m *mockGitRepo) CommitUpdate(filePath string, updateVer string) error {
	return nil
}

func (m *mockGitRepo) CreateBranch(name string) error {
	return nil
}

type mockJiraRepository struct{}

func (m *mockJiraRepository) SearchIssues(jql string) (*service.SearchResults, error) {
	result := &service.SearchResults{
		Issues: []service.Issue{

			{
				ID:   "",
				Self: "",
				Key:  "MOMONGA-1234",
				Fields: &service.IssueFields{
					Type: service.IssueType{
						ID:          "",
						Description: "",
						Name:        "",
					},
					Created:     service.Time{},
					Duedate:     service.Date{},
					Assignee:    &service.User{},
					Updated:     service.Time{},
					Description: "description fo 1234",
					Summary:     "",
					Creator:     &service.User{},
					Reporter:    &service.User{},
					Status:      &service.Status{},
					IssueLinks:  nil,
					Comments:    &service.Comments{},
					FixVersions: nil,
					Labels:      nil,
					Epic:        &service.Epic{},
					Sprint:      &service.Sprint{},
					Parent:      &service.Parent{},
				},
				Names: map[string]string{
					"": "",
				},
			},

			{
				ID:   "",
				Self: "",
				Key:  "MOMONGA-1235",
				Fields: &service.IssueFields{
					Type: service.IssueType{
						ID:          "",
						Description: "",
						Name:        "",
					},
					Created:     service.Time{},
					Duedate:     service.Date{},
					Assignee:    &service.User{},
					Updated:     service.Time{},
					Description: "description fo 1235",
					Summary:     "",
					Creator:     &service.User{},
					Reporter:    &service.User{},
					Status:      &service.Status{},
					IssueLinks:  nil,
					Comments:    &service.Comments{},
					FixVersions: nil,
					Labels:      nil,
					Epic:        &service.Epic{},
					Sprint:      &service.Sprint{},
					Parent:      &service.Parent{},
				},
				Names: map[string]string{
					"": "",
				},
			},
		},
	}

	return result, nil
}

func setUp(t *testing.T) {
	t.Helper()
	gitRepository = &mockGitRepo{}
}

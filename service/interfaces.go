package service

type JiraRepository interface {
	SearchIssues(jql string) (*SearchResults, error)
}

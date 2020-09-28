package service

import (
	"context"
)

type SearchResults struct {
	Expand     string            `json:"expand,omitempty" yaml:"expand,omitempty"`
	Issues     []Issue           `json:"issues,omitempty" yaml:"issues,omitempty"`
	MaxResults int               `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Names      map[string]string `json:"names,omitempty" yaml:"names,omitempty"`
	StartAt    int               `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total      int               `json:"total,omitempty" yaml:"total,omitempty"`
}

type Issue struct {
	Expand         string               `json:"expand,omitempty" structs:"expand,omitempty"`
	ID             string               `json:"id,omitempty" structs:"id,omitempty"`
	Self           string               `json:"self,omitempty" structs:"self,omitempty"`
	Key            string               `json:"key,omitempty" structs:"key,omitempty"`
	Fields         *IssueFields         `json:"fields,omitempty" structs:"fields,omitempty"`
	RenderedFields *IssueRenderedFields `json:"renderedFields,omitempty" structs:"renderedFields,omitempty"`
	Changelog      *Changelog           `json:"changelog,omitempty" structs:"changelog,omitempty"`
	Transitions    []Transition         `json:"transitions,omitempty" structs:"transitions,omitempty"`
	Names          map[string]string    `json:"names,omitempty" structs:"names,omitempty"`
}

type IssueFields struct {
	// TODO Missing fields
	//      * "workratio": -1,
	//      * "lastViewed": null,
	//      * "environment": null,
	Expand               string        `json:"expand,omitempty" structs:"expand,omitempty"`
	Type                 IssueType     `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Project              Project       `json:"project,omitempty" structs:"project,omitempty"`
	Resolutiondate       Time          `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created              Time          `json:"created,omitempty" structs:"created,omitempty"`
	Duedate              Date          `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Assignee             *User         `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Updated              Time          `json:"updated,omitempty" structs:"updated,omitempty"`
	Description          string        `json:"description,omitempty" structs:"description,omitempty"`
	Summary              string        `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator              *User         `json:"Creator,omitempty" structs:"Creator,omitempty"`
	Reporter             *User         `json:"reporter,omitempty" structs:"reporter,omitempty"`
	Status               *Status       `json:"status,omitempty" structs:"status,omitempty"`
	TimeSpent            int           `json:"timespent,omitempty" structs:"timespent,omitempty"`
	TimeEstimate         int           `json:"timeestimate,omitempty" structs:"timeestimate,omitempty"`
	TimeOriginalEstimate int           `json:"timeoriginalestimate,omitempty" structs:"timeoriginalestimate,omitempty"`
	IssueLinks           []*IssueLink  `json:"issuelinks,omitempty" structs:"issuelinks,omitempty"`
	Comments             *Comments     `json:"comment,omitempty" structs:"comment,omitempty"`
	FixVersions          []*FixVersion `json:"fixVersions,omitempty" structs:"fixVersions,omitempty"`
	Labels               []string      `json:"labels,omitempty" structs:"labels,omitempty"`
	Subtasks             []*Subtasks   `json:"subtasks,omitempty" structs:"subtasks,omitempty"`
	Epic                 *Epic         `json:"epic,omitempty" structs:"epic,omitempty"`
	Sprint               *Sprint       `json:"sprint,omitempty" structs:"sprint,omitempty"`
	Parent               *Parent       `json:"parent,omitempty" structs:"parent,omitempty"`
}

type IssueType struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Subtask     bool   `json:"subtask,omitempty" structs:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty" structs:"avatarId,omitempty"`
}
type IssueRenderedFields struct {
	Resolutiondate string `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created        string `json:"created,omitempty" structs:"created,omitempty"`
	Duedate        string `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Updated        string `json:"updated,omitempty" structs:"updated,omitempty"`
	Description    string `json:"description,omitempty" structs:"description,omitempty"`
}
type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}
type ChangelogHistory struct {
	Id      string           `json:"id" structs:"id"`
	Author  User             `json:"author" structs:"author"`
	Created string           `json:"created" structs:"created"`
	Items   []ChangelogItems `json:"items" structs:"items"`
} // ChangelogItems reflects one single changelog item of a history item
type ChangelogItems struct {
	Field      string      `json:"field" structs:"field"`
	FieldType  string      `json:"fieldtype" structs:"fieldtype"`
	From       interface{} `json:"from" structs:"from"`
	FromString string      `json:"fromString" structs:"fromString"`
	To         interface{} `json:"to" structs:"to"`
	ToString   string      `json:"toString" structs:"toString"`
}
type Transition struct {
	ID     string                     `json:"id" structs:"id"`
	Name   string                     `json:"name" structs:"name"`
	To     Status                     `json:"to" structs:"status"`
	Fields map[string]TransitionField `json:"fields" structs:"fields"`
}

// TransitionField represents the value of one Transition
type TransitionField struct {
	Required bool `json:"required" structs:"required"`
}

// Project represents a Jira Project.
type Project struct {
	Expand       string            `json:"expand,omitempty" structs:"expand,omitempty"`
	Self         string            `json:"self,omitempty" structs:"self,omitempty"`
	ID           string            `json:"id,omitempty" structs:"id,omitempty"`
	Key          string            `json:"key,omitempty" structs:"key,omitempty"`
	Description  string            `json:"description,omitempty" structs:"description,omitempty"`
	Lead         User              `json:"lead,omitempty" structs:"lead,omitempty"`
	IssueTypes   []IssueType       `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
	URL          string            `json:"url,omitempty" structs:"url,omitempty"`
	Email        string            `json:"email,omitempty" structs:"email,omitempty"`
	AssigneeType string            `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Versions     []Version         `json:"versions,omitempty" structs:"versions,omitempty"`
	Name         string            `json:"name,omitempty" structs:"name,omitempty"`
	Roles        map[string]string `json:"roles,omitempty" structs:"roles,omitempty"`
}

type Date string
type Status struct{}
type IssueLink struct{}
type Comments struct{}
type FixVersion struct{}
type Subtasks struct{}
type Epic struct{}
type Parent struct{}
type Sprint struct{}
type Time string
type User struct{}
type Version struct{}

type jiraService struct {
	repository JiraRepository
}

// NewJiraService returns domain JiraService
func NewJiraService(repository JiraRepository) *jiraService {
	return &jiraService{
		repository: repository,
	}
}

func (r *jiraService) Search(ctx context.Context, jql string) ([]Issue, error) {
	result, err := r.repository.SearchIssues(jql)
	if err != nil {
		return nil, err
	}

	return result.Issues, nil
}

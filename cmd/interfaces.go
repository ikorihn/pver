package cmd

import (
	"context"

	"github.com/r57ty7/pver/service"
)

// FileVersionManager バージョンファイル操作インターフェース
type FileVersionManager interface {
	SetConfig(conf service.Config)
	Version() string
	Update(newVersion string) error
}

// GitRepository Git関連
type GitRepository interface {
	CommitUpdate(filePath string, updateVer string) error
}

// JiraService JIRA
type JiraService interface {
	Search(ctx context.Context, jql string) ([]service.Issue, error)
}

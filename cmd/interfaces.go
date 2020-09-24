package cmd

import (
	"context"

	"github.com/r57ty7/jiracket/domain"
)

// FileVersionManager バージョンファイル操作インターフェース
type FileVersionManager interface {
	SetConfig(conf Config)
	Version() string
	Update(newVersion string) error
}

// GitRepository Git関連
type GitRepository interface {
	CommitUpdate(filePath string, updateVer string) error
}

// JiraRepository JIRA
type JiraRepository interface {
	Search(ctx context.Context, jql string) ([]domain.Issue, error)
}

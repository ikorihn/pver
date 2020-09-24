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

func setUp(t *testing.T) {
	t.Helper()
	gitRepository = &mockGitRepo{}
}

package cmd

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

package cmd

// FileVersionManager バージョンファイル操作インターフェース
type FileVersionManager interface {
	SetFile(filePath string)
	Version() string
	Update(newVersion string) error
}

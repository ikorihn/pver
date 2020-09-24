package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/r57ty7/pver/cmd"
)

type PackgeJson struct {
	Version string `json:"version"`
}

type NpmProject struct {
	filePath string
	indent   string
}

func NewNpmProject() *NpmProject {
	return &NpmProject{}
}

func (p *NpmProject) parse() (PackgeJson, error) {
	var packageJson PackgeJson

	bytes, err := ioutil.ReadFile(p.filePath)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return packageJson, err
	}

	if err := json.Unmarshal(bytes, &packageJson); err != nil {
		fmt.Printf("parse err: %v\n", err)
		return packageJson, err
	}

	return packageJson, nil
}
func (p *NpmProject) SetConfig(config cmd.Config) {
	p.filePath = config.Npm.Filepath
}

func (p *NpmProject) Version() string {
	packgeJson, err := p.parse()
	if err != nil {
		return ""
	}
	return packgeJson.Version
}

func (p *NpmProject) Update(newVersion string) error {
	bytes, err := ioutil.ReadFile(p.filePath)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return err
	}

	pattern := `("version":\s*)(".*")`
	format := regexp.MustCompile(pattern)

	if !format.Match(bytes) {
		return fmt.Errorf("version tag not found")
	}
	updatedJSON := format.ReplaceAll(bytes, []byte(fmt.Sprintf(`${1}"%s"`, newVersion)))

	err = writeFile(p.filePath, updatedJSON)

	if err != nil {
		fmt.Printf("write file err: %v\n", err)
		return err
	}

	return nil
}

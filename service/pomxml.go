package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
)

type MavenProject struct {
	filePath string
	indent   string
}

type PomXml struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	GroupID              string               `xml:"groupId"`
	ArtifactID           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Packaging            string               `xml:"packaging"`
	Name                 string               `xml:"name"`
	Build                string               `xml:"build"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
	Dependencies         []Dependency         `xml:"dependencies>dependency"`
}

type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
	Classifier string `xml:"classifier"`
	Type       string `xml:"type"`
	Scope      string `xml:"scope"`
}

type Build struct {
	Plugins []Plugin `xml:"plugins>plugin"`
}

type Plugin struct {
	XMLName    xml.Name `xml:"plugin"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version"`
}

type DependencyManagement struct {
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

type Repository struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

func NewMavenProject() *MavenProject {
	return &MavenProject{}
}

var writeFile func(filePath string, content []byte) error

func init() {
	writeFile = func(filePath string, content []byte) error {
		return ioutil.WriteFile(filePath, content, 0644)
	}
}

func (p *MavenProject) parse() (PomXml, error) {
	var pom PomXml

	bytes, err := ioutil.ReadFile(p.filePath)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return pom, err
	}

	if err := xml.Unmarshal(bytes, &pom); err != nil {
		fmt.Printf("parse xml err: %v\n", err)
		return pom, err
	}

	return pom, nil
}

func (p *MavenProject) SetConfig(config Config) {
	p.filePath = config.Pom.Filepath
	p.indent = config.Pom.Indent
}

func (p *MavenProject) Version() string {
	pomXML, err := p.parse()
	if err != nil {
		return ""
	}
	return pomXML.Version
}

func (p *MavenProject) Update(newVersion string) error {
	bytes, err := ioutil.ReadFile(p.filePath)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return err
	}

	// ルート直下のバージョンタグの中身を更新する
	// TODO ルート直下を表す正規表現がかけなかったのでインデント幅で判定する インデントはオプションで設定する
	pattern := fmt.Sprintf(`(?m)^(%s<version>)(.*?)(</version>)$`, p.indent)
	format := regexp.MustCompile(pattern)

	if !format.Match(bytes) {
		return fmt.Errorf("version tag not found")
	}
	updatedXML := format.ReplaceAll(bytes, []byte(fmt.Sprintf("${1}%s$3", newVersion)))

	err = writeFile(p.filePath, updatedXML)

	if err != nil {
		fmt.Printf("write file err: %v\n", err)
		return err
	}

	return nil
}

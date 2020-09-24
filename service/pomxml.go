package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type MavenProject struct {
	filePath string
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

func (p *MavenProject) SetFile(filePath string) {
	p.filePath = filePath
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

	// TODO parseが冗長なので正規表現をつかいたい
	format := "<version>%s</version>"
	pomXML, err := p.parse()
	if err != nil {
		return err
	}

	updatedXML := strings.Replace(string(bytes), fmt.Sprintf(format, pomXML.Version), fmt.Sprintf(format, newVersion), 1)
	err = ioutil.WriteFile(p.filePath, []byte(updatedXML), 0644)
	if err != nil {
		fmt.Printf("write file err: %v\n", err)
		return err
	}

	return nil
}

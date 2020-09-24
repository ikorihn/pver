/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// FileVersionManager の mock
type mockFvm struct {
	version string
}

func (m *mockFvm) SetFile(filePath string) {
}

func (m *mockFvm) Version() string {
	return m.version
}

func (m *mockFvm) Update(newVersion string) error {
	m.version = newVersion
	return nil
}

// pomver pom のテスト
func Test_pomCmd_Execute(t *testing.T) {
	mockFvm := &mockFvm{
		version: "1.0.0",
	}

	cmd := newPomCmd(mockFvm)

	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	want := "Version: 1.0.0\n"

	if string(got) != want {
		t.Fatalf("want '%s', got '%s'", want, got)
	}

}

// pomver pom -u xxx のテスト
func Test_pomCmd_Execute_Update(t *testing.T) {
	mockFvm := &mockFvm{
		version: "1.0.0",
	}

	cmd := newPomCmd(mockFvm)

	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"-u", "1.2.3"})
	cmd.Execute()
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}

	wantMessage := "Version: 1.0.0\nUpdate to => 1.2.3\n"
	wantVersion := "1.2.3"

	if string(got) != wantMessage {
		t.Fatalf("want '%s', got '%s'", wantMessage, got)
	}

	if mockFvm.version != wantVersion {
		t.Fatalf("want '%s', got '%s'", wantVersion, mockFvm.version)
	}

}

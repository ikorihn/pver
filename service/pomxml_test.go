package service

import (
	"testing"
)

func TestMavenProject_Version(t *testing.T) {
	type fields struct {
		filePath string
		indent   string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "normal",
			fields: fields{
				filePath: "testdata/pom.xml",
				indent:   "  ",
			},
			want: "1.1.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MavenProject{
				filePath: tt.fields.filePath,
				indent:   tt.fields.indent,
			}
			if got := p.Version(); got != tt.want {
				t.Errorf("MavenProject.Version() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestMavenProject_Update(t *testing.T) {
	outputMap := map[string]string{}
	writeFile = func(filePath string, content []byte) error {
		outputMap[filePath] = string(content)
		return nil
	}

	type fields struct {
		filePath string
		indent   string
	}
	type args struct {
		newVersion string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantXml string
		wantErr bool
	}{

		{
			name: "normal",
			fields: fields{
				filePath: "testdata/pom.xml",
				indent:   "  ",
			},
			args: args{
				newVersion: "1.2.3",
			},
			wantXml: `<?xml version="1.0" encoding="UTF-8"?>
			<project xmlns="http://maven.apache.org/POM/4.0.0">
				<modelVersion>4.0.0</modelVersion>
				<groupId>com.example</groupId>
				<artifactId>my-app</artifactId>
				<version>1.2.3</version>
				<dependencies>
					<dependency>
						<groupId>junit</groupId>
						<artifactId>junit</artifactId>
						<scope>test</scope>
						<version>3.8.0</version>
					</dependency>
				</dependencies>
			</project>`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MavenProject{
				filePath: tt.fields.filePath,
				indent:   tt.fields.indent,
			}
			if err := p.Update(tt.args.newVersion); (err != nil) != tt.wantErr {
				t.Errorf("MavenProject.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

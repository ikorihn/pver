package service

import (
	"testing"
)

func TestNpmProject_Version(t *testing.T) {
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
				filePath: "testdata/package.json",
				indent:   "  ",
			},
			want: "1.1.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NpmProject{
				filePath: tt.fields.filePath,
			}
			if got := p.Version(); got != tt.want {
				t.Errorf("NpmProject.Version() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNpmProject_Update(t *testing.T) {
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
		name     string
		fields   fields
		args     args
		wantJson string
		wantErr  bool
	}{

		{
			name: "normal",
			fields: fields{
				filePath: "testdata/package.json",
				indent:   "  ",
			},
			args: args{
				newVersion: "1.2.3",
			},
			wantJson: `{
  "name": "sample",
  "version": "1.2.3",
  "dependencies": {
    "typescript": "~3.7.2"
  },
  "scripts": {
    "test": "echo test"
  },
  "devDependencies": {
    "prettier": "^2.0.5"
  }
}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NpmProject{
				filePath: tt.fields.filePath,
				indent:   tt.fields.indent,
			}
			if err := p.Update(tt.args.newVersion); (err != nil) != tt.wantErr {
				t.Errorf("NpmProject.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			if outputMap[tt.fields.filePath] != tt.wantJson {
				t.Errorf("NpmProject.Update() got = %v, want %v", outputMap[tt.fields.filePath], tt.wantJson)
			}
		})
	}
}

package cmd

type Config struct {
	Pom  pom
	Npm  npm
	Jira jira
}

type pom struct {
	Filepath string
	Indent   string
}

type npm struct {
	Filepath string
}

type jiraConfig struct {
	BaseURL  string
	Username string
	Password string
}

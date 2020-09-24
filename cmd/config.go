package cmd

type Config struct {
	Pom pom
	Npm npm
}

type pom struct {
	Filepath string
	Indent   string
}

type npm struct {
	Filepath string
}

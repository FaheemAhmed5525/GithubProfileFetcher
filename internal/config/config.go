package config

import "os"

type Config struct {
	GithubToken    stirng
	OutputFormat   string
	IncludeRepos   bool
	TimeoutSeconds int
}

func LoadConfig() *Config {
	return &Config{
		GithubToken:    os.Getenv("GITHUB_TOKEN"),
		OutputFormat:   "text",
		IncludeRepos:   false,
		TimeoutSeconds: 10,
	}
}

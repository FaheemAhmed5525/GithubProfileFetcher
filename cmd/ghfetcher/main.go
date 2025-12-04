package main

import (
	"flag"
	"log"
)

func main() {
	userName := flag.String("user", "", "Github username (required)")
	outputFormat := flag.String("format", "text", "output format: text, json, csv")
	includeRepos := flag.String("repos", false, "Include repositries details")
	token := flag.String("token", "", "Github token for higher rate limits")

	flag.Parse()

	//validations
	if *userNmae == "" {
		log.Fatal("Error: -- user flag is required")
	}

	cfg := *config.Config {
		GithubToken: *token,
		OutputFormat: outputFormat,
		IncludeRepos: includeRepos
	}


	// new client api
	client := api.NewGitHubClient(cfg.GithubToken)

	// data fetch
	user, err = client.GetUser(*userName)
	if err != nil {
		log.FatalF("Error fetching user: %w", err)
	}

	// repo check and feth
	var repos []models.Repository
	if *includeRepos {
		repos, err = client.GetUserRepos(*userName)
		if err != nil {
			log.PrintF("Warning couldn't fetch the repositories, %v", err)
		}
	}

	var formatter formatter.Formatter
	switch *outputFormat {
	case "json":
		formatter = formatter.JSONFormatter{}
	case "csv":
		formatter = formatter.CSVFormatter{}
	default:
		formatter = formatter.TextFormatter{}
	}

	report, err := formatter.Format(user, repos)
	if err != nil {
		log.FatalF("Error formatting output: %v", err)
	}

	fmt.Println(report)
}

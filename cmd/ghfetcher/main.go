package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/api"
	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/formatter"
	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"
)

func main() {
	// CLI Flags
	username := flag.String("user", "", "GitHub username (required)")
	outputFormat := flag.String("format", "text", "Output format: text, json, csv")
	includeRepos := flag.Bool("repos", false, "Include repository details")
	limitRepos := flag.Int("limit", 10, "Limit number of repositories to fetch")
	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "GitHub token (or set GITHUB_TOKEN env)")
	showRateLimit := flag.Bool("rate-limit", false, "Show GitHub API rate limit info")

	flag.Parse()

	// Validation
	if *username == "" && !*showRateLimit {
		flag.Usage()
		log.Fatal("Error: --user flag is required (or use --rate-limit)")
	}

	// Initialize client
	client := api.NewGitHubClient(*token)

	// Show rate limit if requested
	if *showRateLimit {
		rateLimit, err := client.RateLimit()
		if err != nil {
			log.Fatalf("Error fetching rate limit: %v", err)
		}

		resetTime := time.Unix(rateLimit.Reset, 0)
		fmt.Printf("GitHub API Rate Limits:\n")
		fmt.Printf("  Limit:     %d requests/hour\n", rateLimit.Limit)
		fmt.Printf("  Remaining: %d requests\n", rateLimit.Remaining)
		fmt.Printf("  Resets at: %s\n", resetTime.Format("2006-01-02 15:04:05 MST"))
		fmt.Printf("  Reset in:  %v\n", time.Until(resetTime).Round(time.Minute))
		return
	}

	// Fetch user data
	fmt.Printf("Fetching profile for %s...\n", *username)
	user, err := client.GetUser(*username)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}

	// Fetch repositories if requested
	var repos []models.Repository
	if *includeRepos {
		fmt.Printf("Fetching repositories...\n")
		allRepos, err := client.GetUserRepos(*username)
		if err != nil {
			log.Printf("Warning: Could not fetch repositories: %v", err)
		} else {
			// Limit repos if specified
			if *limitRepos > 0 && *limitRepos < len(allRepos) {
				repos = allRepos[:*limitRepos]
			} else {
				repos = allRepos
			}
		}
	}

	// Create formatter based on output format
	var formatterInstance formatter.Formatter
	switch *outputFormat {
	case "json":
		formatterInstance = formatter.JSONFormatter{}
	case "csv":
		formatterInstance = formatter.CSVFormatter{}
	default:
		formatterInstance = formatter.TextFormatter{}
	}

	// Generate and output report
	report, err := formatterInstance.Format(user, repos)
	if err != nil {
		log.Fatalf("Error formatting output: %v", err)
	}

	fmt.Println(report)

	// Show stats if we have repos
	if len(repos) > 0 {
		stats := models.CalculateStats(user, repos)
		if statsOutput, err := formatterInstance.FormatStats(stats); err == nil {
			fmt.Println("\n" + statsOutput)
		}
	}
}

package formatter

import (
	"fmt"
	"strings"

	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"
)

type TextFormatter struct{}

func (formatter TextFormatter) Format(user *models.User, repos []models.Repository) (string, error) {
	var builder strings.Builder

	totalStars := 0
	language := make(map[string]int)
	for _, repo := range repos {
		totalStars += repo.Stars
		if repo.Language != "" {
			language[repo.Language]++
		}
	}

	builder.WriteString("╔══════════════════════════════════════════════════════════════╗\n")
	builder.WriteString("║                     GITHUB USER PROFILE                      ║\n")
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString(fmt.Sprintf("║ Username:  %-50s ║\n", user.Login))
	builder.WriteString(fmt.Sprintf("║ Name:      %-50s ║\n", user.Name))
	builder.WriteString(fmt.Sprintf("║ Company:   %-50s ║\n", user.Company))
	builder.WriteString(fmt.Sprintf("║ Location:  %-50s ║\n", user.Location))
	builder.WriteString(fmt.Sprintf("║ Email:     %-50s ║\n", user.Email))
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString("║ STATISTICS                                                   ║\n")
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString(fmt.Sprintf("║ • Public Repositories: %-41d ║\n", user.PublicRepos))
	builder.WriteString(fmt.Sprintf("║ • Followers:           %-41d ║\n", user.Followers))
	builder.WriteString(fmt.Sprintf("║ • Following:           %-41d ║\n", user.Following))
	builder.WriteString(fmt.Sprintf("║ • Public Gists:        %-41d ║\n", user.PublicGists))
	builder.WriteString(fmt.Sprintf("║ • Total Stars:         %-41d ║\n", totalStars))
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString("║ TOP LANGUAGES                                                ║\n")
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	// top 5 languages
	count := 0
	for lang, freq := range language {
		if count >= 5 {
			break
		}
		builder.WriteString(fmt.Sprintf("║ • %-20s: %-35d ║\n", lang, freq))
		count++
	}

	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString("║ LINKS                                                        ║\n")
	builder.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	builder.WriteString(fmt.Sprintf("║ • Profile:     %-50s ║\n", user.HTMLURL))
	builder.WriteString(fmt.Sprintf("║ • Blog:        %-50s ║\n", user.Blog))
	builder.WriteString(fmt.Sprintf("║ • Twitter:     @%-49s ║\n", user.TwitterUsername))
	builder.WriteString("╚══════════════════════════════════════════════════════════════╝")

	return builder.String(), nil
}

func (formatter TextFormatter) FormatStats(stats *models.UserStats) (string, error) {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("=== STATISTICS FOR %s ===\n", stats.User.Login))
	builder.WriteString(fmt.Sprintf("Total Repositories: %d\n", len(stats.Repositories)))
	builder.WriteString(fmt.Sprintf("Total Stars: %d\n", stats.TotalStars))
	builder.WriteString(fmt.Sprintf("Total Forks: %d\n", stats.TotalForks))
	builder.WriteString("\nLanguage Distribution:\n")

	for lang, count := range stats.MostUsedLanguages {
		builder.WriteString(fmt.Sprintf(" %s: %d repositories\n", lang, count))
	}

	return builder.String(), nil
}

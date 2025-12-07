package formatter

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"
)

type CSVFormatter struct{}

func (formatter CSVFormatter) Format(user *models.User, repos []models.Repository) (string, error) {
	var builder strings.Builder
	writer := csv.NewWriter(&builder)

	// header for user
	userHeader := []string{
		"Username", "Name", "Company", "Location", "Email",
		"Public Repos", "Followers", "Following", "Public Gists",
		"Profile URL", "Blog",
	}
	if err := writer.Write(userHeader); err != nil {
		return "", err
	}

	// user data
	userData := []string{
		user.Login,
		user.Name,
		user.Company,
		user.Location,
		user.Email,
		fmt.Sprintf("%d", user.PublicRepos),
		fmt.Sprintf("%d", user.Followers),
		fmt.Sprintf("%d", user.Following),
		fmt.Sprintf("%d", user.PublicGists),
		user.HTMLURL,
		user.Blog,
	}
	if err := writer.Write(userData); err != nil {
		return "", err
	}

	writer.Write([]string{})
	writer.Write([]string{})

	// resposts heder
	if len(repos) > 0 {
		repoHeader := []string{
			"Repository", "Description", "Language", "Stars", "Forks", "Updated At",
		}
		writer.Write(repoHeader)

		// data
		for _, repo := range repos {
			repoData := []string{
				repo.FullName,
				repo.Description,
				repo.Language,
				fmt.Sprintf("%d", repo.Stars),
				fmt.Sprintf("%d", repo.Forks),
				repo.UpdatedAt.Format("2006-01-02"),
			}

			writer.Write(repoData)
		}
	}

	// finalize
	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (c CSVFormatter) FormatStats(stats *models.UserStats) (string, error) {
	var sb strings.Builder
	writer := csv.NewWriter(&sb)

	header := []string{"Metric", "Value"}
	writer.Write(header)

	writer.Write([]string{"Total Repositories", fmt.Sprintf("%d", len(stats.Repositories))})
	writer.Write([]string{"Total Stars", fmt.Sprintf("%d", stats.TotalStars)})
	writer.Write([]string{"Total Forks", fmt.Sprintf("%d", stats.TotalForks)})
	writer.Write([]string{"", ""})

	// Languages
	writer.Write([]string{"Language", "Repository Count"})
	for lang, count := range stats.MostUsedLanguages {
		writer.Write([]string{lang, fmt.Sprintf("%d", count)})
	}

	writer.Flush()
	return sb.String(), nil
}

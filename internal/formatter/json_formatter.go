package formatter

import (
	"encoding/json"

	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"
)

type JSONFormatter struct{}

func (j JSONFormatter) Format(user *models.User, repos []models.Repository) (string, error) {
	// Calculate aggregated data
	totalStars := 0
	totalForks := 0
	languages := make(map[string]int)

	for _, repo := range repos {
		totalStars += repo.Stars
		totalForks += repo.Forks
		if repo.Language != "" {
			languages[repo.Language]++
		}
	}

	data := map[string]interface{}{
		"user": user,
		"metadata": map[string]interface{}{
			"repositories_count": len(repos),
			"total_stars":        totalStars,
			"total_forks":        totalForks,
			"languages":          languages,
		},
		"repositories": repos,
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (j JSONFormatter) FormatStats(stats *models.UserStats) (string, error) {
	jsonBytes, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

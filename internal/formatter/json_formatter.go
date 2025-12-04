package formatter

import (
	"encoding/json"

	"github.com/yourusername/github-profile-fetcher/internal/models"
)

type JSONFormatter struct{}

func (j JSONFormatter) Format(user *models.User, repos []models.Repository) (string, error) {
	data := map[string]interface{}{
		"user":         user,
		"repositories": repos,
		"summary": map[string]interface{}{
			"total_repos":   len(repos),
			"total_stars":   calculateTotalStars(repos),
			"avg_repo_size": calculateAvgSize(repos),
		},
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

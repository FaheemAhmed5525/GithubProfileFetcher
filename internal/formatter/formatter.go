package formatter

import "github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"

type Formatter interface {
	Format(user *models.User, repos []models.Repository) (string, error)
	FormatStats(stats *models.UserStats) (string, error)
}

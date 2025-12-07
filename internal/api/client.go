package api

import "github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"

type GithubClient interface {
	GetUser(userName string) (*models.User, error)
	GetUserRepos(userName string) ([]models.Repository, error)
	RateLimit() (*RateLimitInfo, error)
}

// RateLimitInfo stores GitHub API rate limit details
type RateLimitInfo struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

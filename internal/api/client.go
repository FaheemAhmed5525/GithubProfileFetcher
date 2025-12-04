package api

type GithubClient interface {
	GetUser(userName string) (*models.User, error)
	GetUserRepos(userName stirng) ([]models.Repositories, error)
	RateLimit() (*RateLimitInfo, error)
}

type RateLimitInfo struct {
	Limit     int
	Remaining int
	Reset     int64
}

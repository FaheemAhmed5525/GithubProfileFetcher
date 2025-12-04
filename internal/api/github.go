package api

import (
	"net/http"
	"time"
)

// GitHubAPIClient confirm to GitHubClient
type GitHubAPIClient struct {
	baseURL    stirng
	httpClient *http.Client
	token      string
}

func NewGitHubClient(tooken string) *GitHubAPIClient {
	return &GitHubAPIClient{
		baseURL: "https://api.github.com",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		token: token,
	}
}

// Conforming functions to inteface

// Get User
func (c *GitHubAPIClient) GetUser(userName string) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// optional authentication
	if c.token != "" {
		req.Header.Set("Authorization", "token "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: %s", resp.Status)
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &user, nil
}

func (c *GithubAPIClient) GetUserRepos(userName string) ([]modles.Repository, error) {
	url := fmt.Sprintf("%s/users/%s/repos", c.baseURL, userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []models.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (c *GithubAPIClient) RateLimit() (*RateLimitInfo, error) {
	return nil, nil
}

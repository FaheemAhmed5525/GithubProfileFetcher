package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/FaheemAhmed5525/GithubProfileFetcher/internal/models"
)

// GitHubAPIClient confirm to GitHubClient
type GitHubAPIClient struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

func NewGitHubClient(token string) *GitHubAPIClient {
	return &GitHubAPIClient{
		baseURL: "https://api.github.com",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    20 * time.Second,
				DisableCompression: false,
			},
		},
		token: token,
	}
}

// Conforming functions to inteface

// Get User
func (c *GitHubAPIClient) GetUser(userName string) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, userName)

	req, err := c.createRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	// Read entire body first for error handling
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		// Try to parse GitHub error message
		var githubErr struct {
			Message          string `json:"message"`
			DocumentationURL string `json:"documentation_url"`
		}
		if err := json.Unmarshal(body, &githubErr); err == nil && githubErr.Message != "" {
			return nil, fmt.Errorf("GitHub API error: %s", githubErr.Message)
		}
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}

	return &user, nil
}

func (c *GitHubAPIClient) GetUserRepos(userName string) ([]models.Repository, error) {
	url := fmt.Sprintf("%s/users/%s/repos", c.baseURL, userName)
	var allRepos []models.Repository
	page := 1
	perPage := 100 // GitHub max per page

	for {
		reqUrl := fmt.Sprintf("%s?page=%d&per_page=%d&sort=updated", url, page, perPage)

		req, err := c.createRequest("GET", reqUrl, nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			break
		}

		var repos []models.Repository
		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			return nil, err
		}

		if len(repos) == 0 {
			break // No more repos
		}

		allRepos = append(allRepos, repos...)

		// Check if we got fewer than perPage, means last page
		if len(repos) < perPage {
			break
		}

		page++

		// Safety limit
		if page > 10 {
			break
		}
	}

	return allRepos, nil
}

func (c *GitHubAPIClient) RateLimit() (*RateLimitInfo, error) {
	url := fmt.Sprintf("%s/rate_limit", c.baseURL)

	req, err := c.createRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse rate limit headers (GitHub provides them in headers too)
	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	reset, _ := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)

	if limit == 0 || remaining == 0 {
		// Fallback to API endpoint
		var rateLimitResponse struct {
			Resources struct {
				Core struct {
					Limit     int   `json:"limit"`
					Remaining int   `json:"remaining"`
					Reset     int64 `json:"reset"`
				} `json:"core"`
			} `json:"resources"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&rateLimitResponse); err != nil {
			return &RateLimitInfo{
				Limit:     limit,
				Remaining: remaining,
				Reset:     reset,
			}, nil
		}

		limit = rateLimitResponse.Resources.Core.Limit
		remaining = rateLimitResponse.Resources.Core.Remaining
		reset = rateLimitResponse.Resources.Core.Reset
	}

	return &RateLimitInfo{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}, nil
}

func (c *GitHubAPIClient) createRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "GitHub-Profile-Fetcher/1.0")

	// Add authentication if token exists
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	return req, nil
}

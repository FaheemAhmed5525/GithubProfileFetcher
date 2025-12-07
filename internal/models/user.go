package models

import (
	"sort"
	"time"
)

// User represents GitHub user
type User struct {
	Login           string    `json:"login"`
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Company         string    `json:"company"`
	Blog            string    `json:"blog"`
	Location        string    `json:"location"`
	Email           string    `json:"email"`
	Bio             string    `json:"bio"`
	TwitterUsername string    `json:"twitter_username"`
	PublicRepos     int       `json:"public_repos"`
	PublicGists     int       `json:"public_gists"`
	Followers       int       `json:"followers"`
	Following       int       `json:"following"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	HTMLURL         string    `json:"html_url"`
}

// CalculateStats processes repositories to generate statistics
func CalculateStats(user *User, repos []Repository) *UserStats {
	stats := &UserStats{
		User:              user,
		Repositories:      repos,
		MostUsedLanguages: make(map[string]int),
		TotalStars:        0,
		TotalForks:        0,
	}

	for _, repo := range repos {
		stats.TotalStars += repo.Stars
		stats.TotalForks += repo.Forks

		if repo.Language != "" {
			stats.MostUsedLanguages[repo.Language]++
		}
	}

	return stats
}

// GetTopLanguages returns top N languages by usage
func (s *UserStats) GetTopLanguages(n int) []struct {
	Language string
	Count    int
} {
	type langCount struct {
		Language string
		Count    int
	}

	var langs []langCount
	for lang, count := range s.MostUsedLanguages {
		langs = append(langs, langCount{Language: lang, Count: count})
	}

	// Sort by count descending
	sort.Slice(langs, func(i, j int) bool {
		return langs[i].Count > langs[j].Count
	})

	// Return top N
	if n > len(langs) {
		n = len(langs)
	}

	result := make([]struct {
		Language string
		Count    int
	}, n)
	for i := 0; i < n; i++ {
		result[i].Language = langs[i].Language
		result[i].Count = langs[i].Count
	}
	return result
}

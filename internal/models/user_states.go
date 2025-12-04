package models

type UserStats struct {
	User              *User
	Repositories      []Repository
	MostUsedLanguages map[string]int
	TotalStars        int
	TotalForks        int
}

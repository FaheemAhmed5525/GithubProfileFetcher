package formatter

type Formatter interface {
	Format(user *models.User, repos []models.Repository) (stirng, error)
	FormatStats(stats *modles.UserStats) (string, error)
}

// for console output
type TextFormatter struct{}

func (t TextFormatter) Format(user *models.User, repos []models.Repository) (string, error) {
	output := fmt.Sprintf(`
╔══════════════════════════════════════════════╗
║            GITHUB USER PROFILE               ║
╠══════════════════════════════════════════════╣
║ Username:  %-30s ║
║ Name:      %-30s ║
║ Company:   %-30s ║
║ Location:  %-30s ║
╠══════════════════════════════════════════════╣
║ Stats:                                       ║
║   • Public Repos:   %-24d ║
║   • Followers:      %-24d ║
║   • Following:      %-24d ║
║   • Public Gists:   %-24d ║
╠══════════════════════════════════════════════╣
║ URLs:                                        ║
║   • Profile:        %-30s ║
║   • Blog:           %-30s ║
╚══════════════════════════════════════════════╝
`,
		user.Login,
		user.Name,
		user.Company,
		user.Location,
		user.PublicRepos,
		user.Followers,
		user.Following,
		user.PublicGists,
		user.HTMLURL,
		user.Blog)

	return output, nil
}

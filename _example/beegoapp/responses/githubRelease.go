package responses

import "time"

type GithubRelease struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	RepoName    string
	TagName     string
	UserName    string
	PreRelease  bool
	Message     string
}

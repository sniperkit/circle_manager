package responses

import "time"

type GithubCommit struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	RepoName    string
	Comments    string
	UserName    string
	BranchName  string
}

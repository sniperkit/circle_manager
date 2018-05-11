package requests

import "time"

var _ = time.Time{}

type CreateGithubRelease struct {
	RepoName   string
	TagName    string
	UserName   string
	PreRelease bool
	Message    string
}

type UpdateGithubRelease struct {
	RepoName   string
	TagName    string
	UserName   string
	PreRelease bool
	Message    string
}

func (c *CreateGithubRelease) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateGithubRelease) Valid() error {
	return validate.Struct(c)
}

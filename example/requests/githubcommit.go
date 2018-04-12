package requests

import "time"

type CreateGithubCommit struct {
	RepoName   string
	Comments   string
	UserName   string
	BranchName string
}

type UpdateGithubCommit struct {
	RepoName   string
	Comments   string
	UserName   string
	BranchName string
}

func (c *CreateGithubCommit) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateGithubCommit) Valid() error {
	return validate.Struct(c)
}

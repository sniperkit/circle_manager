package models

import "time"

var _ = time.Time{}

// gen:qs
type GithubCommit struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	RepoName    string
	Comments    string
	UserName    string
	BranchName  string
}

func AddGithubCommit(githubCommit *GithubCommit) (id uint, err error) {
	err = githubCommit.Create(gGormDB)
	id = githubCommit.ID
	return
}

func GetGithubCommitByID(id uint) (githubCommit *GithubCommit, err error) {
	githubCommit = &GithubCommit{
		ID: id,
	}
	err = NewGithubCommitQuerySet(gGormDB).
		One(githubCommit)
	return
}

func GetAllGithubCommit(queryPage *QueryPage) (githubCommits []GithubCommit, err error) {
	err = NewGithubCommitQuerySet(gGormDB).
		All(&githubCommits)
	return
}

func UpdateGithubCommitByID(githubCommit *GithubCommit) (err error) {
	err = githubCommit.Update(gGormDB,
		GithubCommitDBSchema.Name,
		GithubCommitDBSchema.Description,
		GithubCommitDBSchema.RepoName,
		GithubCommitDBSchema.Comments,
		GithubCommitDBSchema.UserName,
		GithubCommitDBSchema.BranchName,
	)
	return
}

func DeleteGithubCommit(id uint) (err error) {
	githubCommit := &GithubCommit{
		ID: id,
	}
	err = githubCommit.Delete(gGormDB)
	return
}

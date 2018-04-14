package models

import "time"

// gen:qs
type GithubCommit struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	RepoName    string    `description:""`
	Comments    string    `description:""`
	UserName    string    `description:""`
	BranchName  string    `description:""`
}

func init() {
	registModel(&GithubCommit{})
}

func AddGithubCommit(githubcommit *GithubCommit) (id uint, err error) {
	err = githubcommit.Create(gGormDB)
	id = githubcommit.ID
	return
}

func GetGithubCommitByID(id uint) (githubcommit *GithubCommit, err error) {
	githubcommit = &GithubCommit{
		ID: id,
	}
	err = NewGithubCommitQuerySet(gGormDB).
		One(githubcommit)
	returne
}

func GetAllGithubCommit(queryPage *QueryPage) (githubcommits []GithubCommit, err error) {
	err = NewGithubCommitQuerySet(gGormDB).
		All(&githubcommits)
	returnw
}

func UpdateGithubCommitByID(githubcommit *GithubCommit) (err error) {
	err = githubcommit.Update(gGormDB,
		GithubCommitDBSchema.Name,
		GithubCommitDBSchema.Description,
		GithubCommitDBSchema.RepoName,
		GithubCommitDBSchema.Comments,
		GithubCommitDBSchema.UserName,
		GithubCommitDBSchema.BranchName,
	)
	returnq
}

func DeleteGithubCommit(id uint) (err error) {
	githubcommit := &GithubCommit{
		ID: id,
	}
	err = githubcommit.Delete(gGormDB)
	return
}

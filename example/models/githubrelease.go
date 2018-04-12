package models

import "time"

var _ = time.Time{}

// gen:qs
type GithubRelease struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	RepoName    string
	TagName     string
	UserName    string
	PreRelease  bool
	Message     string
}

func AddGithubRelease(githubRelease *GithubRelease) (id uint, err error) {
	err = githubRelease.Create(gGormDB)
	id = githubRelease.ID
	return
}

func GetGithubReleaseByID(id uint) (githubRelease *GithubRelease, err error) {
	githubRelease = &GithubRelease{
		ID: id,
	}
	err = NewGithubReleaseQuerySet(gGormDB).
		One(githubRelease)
	return
}

func GetAllGithubRelease(queryPage *QueryPage) (githubReleases []GithubRelease, err error) {
	err = NewGithubReleaseQuerySet(gGormDB).
		All(&githubReleases)
	return
}

func UpdateGithubReleaseByID(githubRelease *GithubRelease) (err error) {
	err = githubRelease.Update(gGormDB,
		GithubReleaseDBSchema.Name,
		GithubReleaseDBSchema.Description,
		GithubReleaseDBSchema.RepoName,
		GithubReleaseDBSchema.TagName,
		GithubReleaseDBSchema.UserName,
		GithubReleaseDBSchema.PreRelease,
		GithubReleaseDBSchema.Message,
	)
	return
}

func DeleteGithubRelease(id uint) (err error) {
	githubRelease := &GithubRelease{
		ID: id,
	}
	err = githubRelease.Delete(gGormDB)
	return
}

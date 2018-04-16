package models

import "time"

// gen:qs
type GithubRelease struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	RepoName    string    `description:""`
	TagName     string    `description:""`
	UserName    string    `description:""`
	PreRelease  bool      `description:""`
	Message     string    `description:""`
}

func init() {
	registModel(&GithubRelease{})
}

func AddGithubRelease(githubrelease *GithubRelease) (id uint, err error) {
	err = githubrelease.Create(gGormDB)
	id = githubrelease.ID
	return
}

func GetGithubReleaseByID(id uint) (githubrelease *GithubRelease, err error) {
	githubrelease = &GithubRelease{
		ID: id,
	}
	err = NewGithubReleaseQuerySet(gGormDB).
		One(githubrelease)
	return
}

func GetAllGithubRelease(queryPage *QueryPage) (githubreleases []GithubRelease, err error) {
	err = NewGithubReleaseQuerySet(gGormDB).
		All(&githubreleases)
	return
}

func UpdateGithubReleaseByID(githubrelease *GithubRelease) (err error) {
	err = githubrelease.Update(gGormDB,
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
	githubrelease := &GithubRelease{
		ID: id,
	}
	err = githubrelease.Delete(gGormDB)
	return
}

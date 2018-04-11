// @APIVersion 1.0.0
// @Title Circle API
// @Description Circle API
// @Contact leejungju.go@gmail.com
// @TermsOfServiceUrl http://circle.com
// @License Private
// @SecurityDefinition userAPIKey apiKey X-USER-AUTH-TOKEN header "I love auto-generated docs
package models

func setTable() error {
	return gGormDB.AutoMigrate(
		// @manual start
		&GithubCommit{},
		&GithubRelease{},
		&Event{},
		&Employee{},
		&KeyEvent{},
		&Project{},
		&Todo{},
		&Team{},
		// @manual end

		// @auto start
		// @auto end
	).Error
}

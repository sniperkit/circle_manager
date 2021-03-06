package admin

import (
	"github.com/jungju/circle/models"
	"github.com/qor/admin"
)

func setCircle(a *admin.Admin) {
	// circle:system:start
	// circle:system:end

	// circle:manual:start
	addResourceAndMenu(&models.GithubCommit{}, "Commits", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.GithubRelease{}, "Releases", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.Event{}, "이벤트", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.Employee{}, "직원", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.KeyEvent{}, "주요일정", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.Project{}, "프로젝트", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.Todo{}, "할일", "이벤트관리", anyoneAllow, -1)
	addResourceAndMenu(&models.Team{}, "팀", "이벤트관리", anyoneAllow, -1)
	// circle:manual:end

	// circle:auto:start

	addResourceAndMenu(&models.GithubCommit{}, "Commits", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.GithubRelease{}, "Releases", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.Event{}, "이벤트", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.Employee{}, "직원", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.KeyEvent{}, "주요일정", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.Project{}, "프로젝트", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.Todo{}, "할일", "이벤트관리", anyoneAllow, -1)

	addResourceAndMenu(&models.Team{}, "팀", "이벤트관리", anyoneAllow, -1)

	// circle:auto:end
}

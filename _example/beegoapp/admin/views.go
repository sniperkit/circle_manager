package admin

import (
	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/modules"
	"github.com/qor/admin"
	"github.com/qor/roles"
)

var (
	adminPage  *admin.Admin
	anyoneRole *roles.Permission

	notificationRes     *admin.Resource
	notificationTypeRes *admin.Resource

	githubCommitRes  *admin.Resource
	githubReleaseRes *admin.Resource
	keyEventRes      *admin.Resource
	sprintRes        *admin.Resource
	todoRes          *admin.Resource
	teamRes          *admin.Resource
	employeeRes      *admin.Resource
	eventRes         *admin.Resource

	icsRes    *admin.Resource
	trelloRes *admin.Resource

	anyoneAllow *roles.Permission

	circleSetRes          *admin.Resource
	circleUnitRes         *admin.Resource
	circleUnitPropertyRes *admin.Resource
)

func init() {
	anyoneAllow = roles.Allow(roles.CRUD, roles.Anyone)
}

func setViews(a *admin.Admin) {
	adminPage = a
	a.SetSiteName("Circle")

	circleSetRes = addResourceAndMenu(&modules.CircleSet{}, "Set", "Circle", anyoneAllow, -1)
	circleSetRes.NewAttrs("-Units")
	circleSetRes.EditAttrs("-Units")
	circleSetRes.ShowAttrs("-Units")
	circleSetRes.IndexAttrs("-Units")

	circleUnitRes = addResourceAndMenu(&modules.CircleUnit{}, "Unit", "Circle", anyoneAllow, -1)
	circleUnitRes.NewAttrs("-Properties")
	circleUnitRes.EditAttrs("-Properties")
	circleUnitRes.ShowAttrs("-Properties")
	circleUnitRes.IndexAttrs("-Properties")

	circleUnitPropertyRes = addResourceAndMenu(&modules.CircleUnitProperty{}, "속성", "Circle", anyoneAllow, -1)

	notificationRes = addResourceAndMenu(&modules.Notification{}, "알림", "사용자관리", anyoneAllow, -1)
	notificationTypeRes = addResourceAndMenu(&modules.NotificationType{}, "알림타입", "사용자관리", anyoneAllow, -1)
	icsRes = addResourceAndMenu(&models.Trello{}, "Trello", "동기화관리", anyoneAllow, -1)
	trelloRes = addResourceAndMenu(&models.Ics{}, "ICS", "동기화관리", anyoneAllow, -1)

	setCircle(a)
}

func addResourceAndMenu(value interface{}, menuViewName string, parentMenu string, permission *roles.Permission, priority int) *admin.Resource {
	res := adminPage.AddResource(value, &admin.Config{Menu: []string{parentMenu}, Permission: permission, Priority: priority})

	return res
}

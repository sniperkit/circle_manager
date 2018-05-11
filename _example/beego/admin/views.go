package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/modules"
	"github.com/qor/admin"
	"github.com/qor/qor"
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
	res.SaveHandler = func(result interface{}, context *qor.Context) error {

		if modelItem, ok := result.(modules.ModelItem); ok {
			user, ok := context.CurrentUser.(*models.User)

			// 이벤트 처리
			if context.Request.Method == "POST" {
				if ok {
					modelItem.SetCreatorID(user.ID)
				}
				go modules.EventThenCreate(modelItem, &user.ID)
			} else if context.Request.Method == "PUT" {
				go modules.EventThenUpdate(modelItem, nil, &user.ID)
			}
		}

		//https://github.com/qor/qor/blob/d696f1942afc36458ef5bc19710145ea6fa93e7e/resource/crud.go#L129
		if (context.GetDB().NewScope(result).PrimaryKeyZero() &&
			res.HasPermission(roles.Create, context)) || // has create permission
			res.HasPermission(roles.Update, context) { // has update permission
			return context.GetDB().Save(result).Error
		}
		return roles.ErrPermissionDenied
	}
	res.DeleteHandler = func(result interface{}, context *qor.Context) error {
		if modelItem, ok := result.(modules.ModelItem); ok {
			if user, ok := context.CurrentUser.(*models.User); ok {
				go modules.EventThenDelete(modelItem, &user.ID)
			}
		}

		if res.HasPermission(roles.Delete, context) {
			if primaryQuerySQL, primaryParams := res.ToPrimaryQueryParams(context.ResourceID, context); primaryQuerySQL != "" {
				if !context.GetDB().First(result, append([]interface{}{primaryQuerySQL}, primaryParams...)...).RecordNotFound() {
					return context.GetDB().Delete(result).Error
				}
			}
			return gorm.ErrRecordNotFound
		}
		return roles.ErrPermissionDenied
	}

	menuName := res.Name
	if !res.Config.Singleton {
		menuName = inflection.Plural(res.Name)
	}
	menu := adminPage.GetMenu(menuName)
	menu.Name = menuViewName

	if meta := res.GetMeta("CreatorID"); meta != nil {
		res.IndexAttrs("-Description")
		res.EditAttrs("-CreatorID")
		res.NewAttrs("-CreatorID")
		res.Meta(&admin.Meta{Name: "CreatorID", Label: "작성자", Valuer: func(result interface{}, context *qor.Context) interface{} {
			if modelItem, ok := result.(modules.ModelItem); ok {
				if modelItem.GetCreatorID() > 0 {
					if user, err := models.GetOnlyUserByID(modelItem.GetCreatorID()); err == nil {
						return user.Name
					}
				}
			}

			return "-"
		}})
	}

	return res
}

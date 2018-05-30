package modules

import (
	"html/template"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

type CircleQor struct {
	QorAdmin *admin.Admin
}

func (m *CircleQor) CrudEvent(currentUserID uint, result interface{}, context *qor.Context, oldData string) {
	actionName := ""
	if context.Request.Method == "POST" || context.Request.Method == "PUT" {
		if context.ResourceID == "" {
			actionName = CreateActionTypeName
		} else {
			actionName = UpdateActionTypeName
		}
	} else if context.Request.Method == "DELETE" {
		actionName = DeleteActionTypeName
	} else {
		logrus.
			WithField("method", context.Request.Method).
			WithField("resourceID", context.ResourceID).
			Warn("Unknown action")
	}

	// TODO: context의 ResourceID를 사용?
	resourceID := uint(0)
	if field, ok := structs.New(result).FieldOk("ID"); ok {
		resourceID = field.Value().(uint)
	}

	if _, err := AddCrudEvent(&CrudEvent{
		ActionName:   actionName,
		ActionType:   actionName,
		ResourceID:   resourceID,
		ResourceName: structs.Name(result),
		CreatorID:    currentUserID,
		Where:        "QOR",
		UpdatedData:  ConvJsonData(result),
		OldData:      oldData,
	}); err != nil {
		logrus.WithError(err).Error("")
	}
}

func (m *CircleQor) AddResourceAndMenu(value interface{}, menuViewName string, parentMenu string, permission *roles.Permission, priority int) *admin.Resource {
	res := m.QorAdmin.AddResource(value, &admin.Config{Menu: []string{parentMenu}, Permission: permission, Priority: priority})

	menuName := res.Name
	if !res.Config.Singleton {
		menuName = inflection.Plural(res.Name)
	}
	menu := m.QorAdmin.GetMenu(menuName)
	menu.Name = menuViewName

	matas := res.GetMetas(nil)
	resStruct := structs.New(res.NewStruct())
	appendAttr := []interface{}{}
	for _, mata := range matas {
		name := mata.GetName()
		switch name {
		case "ID", "CreatedAt", "CreatorID", "UpdaterID", "Name", "Description":
		default:
			appendAttr = append(appendAttr, name)
		}

		if resStruct.Field(name).Kind() == reflect.Bool {
			res.Meta(&admin.Meta{Name: name, Setter: mata.GetSetter(), Valuer: func(result interface{}, context *qor.Context) interface{} {
				value := structs.New(result).Field(name).Value()
				if context.ResourceID == "" {
					if boolValue, ok := value.(bool); ok && boolValue {
						return template.HTML(`<input type="checkbox" checked="checked" readonly/>`)
					}
					return template.HTML(`<input type="checkbox" readonly/>`)
				}
				return value
			}})
		}
	}

	if meta := res.GetMeta("CreatorID"); meta != nil {
		res.Meta(&admin.Meta{Name: "CreatorID", Label: "작성자", Type: "readonly", Valuer: func(result interface{}, context *qor.Context) interface{} {
			return extractUserNameByField("CreatorID", result)
		}})
	}

	if meta := res.GetMeta("UpdaterID"); meta != nil {
		res.EditAttrs("-UpdaterID")
		res.NewAttrs("-UpdaterID")
		res.Meta(&admin.Meta{Name: "UpdaterID", Label: "최종수정자", Type: "readonly", Valuer: func(result interface{}, context *qor.Context) interface{} {
			return extractUserNameByField("UpdaterID", result)
		}})
	}

	for _, meta := range []struct {
		FieldName string
		Label     string
		Type      string
	}{
		{"Description", "설명", ""},
		{"Name", "이름", ""},
		{"CreatedAt", "작성일", "readonly"},
		{"UpdatedAt", "수정일", "readonly"},
	} {
		if _, ok := resStruct.FieldOk(meta.FieldName); ok {
			res.Meta(&admin.Meta{Name: meta.FieldName, Label: meta.Label, Type: meta.Type})
		}
	}

	_, creatorIDOK := resStruct.FieldOk("CreatorID")
	_, updaterIDOK := resStruct.FieldOk("UpdaterID")
	_, createdAtIDOK := resStruct.FieldOk("CreatedAt")
	_, updatedAtIDOK := resStruct.FieldOk("UpdatedAt")
	if creatorIDOK && updaterIDOK && createdAtIDOK && updatedAtIDOK {
		res.EditAttrs("-CreatorID", "-CreatedAt", "-UpdaterID", "-UpdatedAt")
		res.NewAttrs("-CreatorID", "-CreatedAt", "-UpdaterID", "-UpdatedAt")
	}

	SetIndexAttrs(res, appendAttr...)

	res.SaveHandler = func(result interface{}, context *qor.Context) error {
		currentUserID := structs.New(context.CurrentUser).Field("ID").Value().(uint)
		if currentUserID > 0 {
			if context.ResourceID == "" {
				structs.New(result).Field("CreatorID").Set(currentUserID)
			}
			structs.New(result).Field("UpdaterID").Set(currentUserID)
		}

		oldData := ""
		if context.ResourceID != "" {
			if resIDUint64, err := strconv.ParseUint(context.ResourceID, 10, 64); err == nil {
				oldModelItem := reflect.New(reflect.ValueOf(result).Elem().Type()).Interface()
				if err := GetItemByID(uint(resIDUint64), oldModelItem); err == nil {
					oldData = ConvJsonData(oldModelItem)
				}
			}
		}

		//https://github.com/qor/qor/blob/d696f1942afc36458ef5bc19710145ea6fa93e7e/resource/crud.go#L129
		if (context.GetDB().NewScope(result).PrimaryKeyZero() &&
			res.HasPermission(roles.Create, context)) || // has create permission
			res.HasPermission(roles.Update, context) { // has update permission
			if err := context.GetDB().Save(result).Error; err != nil {
				return err
			}
			go m.CrudEvent(currentUserID, result, context, oldData)
			return nil
		}
		return roles.ErrPermissionDenied
	}
	res.DeleteHandler = func(result interface{}, context *qor.Context) error {
		currentUserID := uint(0)
		if _, ok := result.(ModelItem); ok {
			currentUserID = structs.New(context.CurrentUser).Field("ID").Value().(uint)
		}

		if res.HasPermission(roles.Delete, context) {
			if primaryQuerySQL, primaryParams := res.ToPrimaryQueryParams(context.ResourceID, context); primaryQuerySQL != "" {
				if !context.GetDB().First(result, append([]interface{}{primaryQuerySQL}, primaryParams...)...).RecordNotFound() {
					if err := context.GetDB().Delete(result).Error; err != nil {
						return err
					}
					go m.CrudEvent(currentUserID, result, context, "")
					return nil
				}
			}
			return gorm.ErrRecordNotFound
		}
		return roles.ErrPermissionDenied
	}

	return res
}

func extractValueByField(fieldName string, result interface{}) interface{} {
	if field := structs.New(result).Field(fieldName); field != nil {
		return field.Value()
	}
	return nil
}

func extractUserNameByField(fieldName string, result interface{}) string {
	value := extractValueByField(fieldName, result)
	if userID, ok := value.(uint); ok {
		if value, err := GetValueByKeyOfTableName("users", "name", userID); err == nil {
			return value.(string)
		}
	}

	return "-"
}

func SetIndexAttrs(res *admin.Resource, attr ...interface{}) {
	indexAttr := []interface{}{"ID", "Name"}
	indexAttr = append(indexAttr, attr...)
	indexAttr = append(indexAttr, "CreatorID", "CreatedAt", "UpdaterID", "UpdatedAt")
	res.IndexAttrs(indexAttr...)
}

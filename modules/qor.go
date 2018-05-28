package modules

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"

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
	modelItem, ok := result.(ModelItem)
	if !ok {
		return
	}

	userID := structs.New(context.CurrentUser).Field("ID").Value().(uint)
	modelItem.SetCreatorID(userID)

	action := ""
	if context.ResourceID == "" && context.Request.Method == "POST" {
		action = "create"
	} else if (context.Request.Method == "PUT") ||
		(context.ResourceID != "" && context.Request.Method == "POST") {
		action = "update"
	} else if context.Request.Method == "DELETE" {
		action = "delete"
	}

	targetID := uint(0)
	if field, ok := structs.New(result).FieldOk("ID"); ok {
		targetID = field.Value().(uint)
	}

	if _, err := AddCrudEvent(&CrudEvent{
		Action:       action,
		TargetID:     targetID,
		TargetObject: structs.Name(modelItem),
		CreatorID:    currentUserID,
		Where:        "QOR",
		UpdatedData:  convJsonData(modelItem),
		OldData:      oldData,
	}); err != nil {
		fmt.Println(err)
	}
}

func (m *CircleQor) AddResourceAndMenu(value interface{}, menuViewName string, parentMenu string, permission *roles.Permission, priority int) *admin.Resource {
	res := m.QorAdmin.AddResource(value, &admin.Config{Menu: []string{parentMenu}, Permission: permission, Priority: priority})
	res.SaveHandler = func(result interface{}, context *qor.Context) error {
		currentUserID := uint(0)
		if modelItem, ok := result.(ModelItem); ok {
			currentUserID = structs.New(context.CurrentUser).Field("ID").Value().(uint)
			if context.ResourceID == "" {
				modelItem.SetCreatorID(currentUserID)
			}
			modelItem.SetUpdaterID(currentUserID)
		}

		oldData := ""
		if context.ResourceID != "" {
			if resIDUint64, err := strconv.ParseUint(context.ResourceID, 10, 64); err == nil {
				oldModelItem := reflect.New(reflect.ValueOf(result).Elem().Type()).Interface()
				if err := GetItemByID(uint(resIDUint64), oldModelItem); err == nil {
					oldData = convJsonData(oldModelItem)
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
		if name != "ID" &&
			name != "CreatedAt" &&
			name != "UpdatedAt" &&
			name != "CreatorID" &&
			name != "UpdaterID" &&
			name != "Name" &&
			name != "Description" {
			appendAttr = append(appendAttr, name)
		}
		if resStruct.Field(name).Kind() == reflect.Bool {
			res.Meta(&admin.Meta{Name: name, Setter: mata.GetSetter(), Valuer: func(result interface{}, context *qor.Context) interface{} {
				value := structs.New(result).Field(name).Value()
				if context.ResourceID == "" {
					if boolValue, ok := value.(bool); ok {
						if boolValue {
							return template.HTML(`<input type="checkbox" checked="checked" readonly/>`)
						}
					}
					return template.HTML(`<input type="checkbox" readonly/>`)
				}
				return value
			}})
		}
	}

	if meta := res.GetMeta("CreatorID"); meta != nil {
		res.Meta(&admin.Meta{Name: "CreatorID", Label: "작성자", Type: "readonly", Valuer: func(result interface{}, context *qor.Context) interface{} {
			if modelItem, ok := result.(ModelItem); ok {
				if modelItem.GetCreatorID() > 0 {
					if value, err := GetValueByKeyOfTableName("users", "name", modelItem.GetCreatorID()); err == nil {
						return value.(string)
					}
				}
			}

			return "-"
		}})
	}

	if meta := res.GetMeta("UpdaterID"); meta != nil {
		res.EditAttrs("-UpdaterID")
		res.NewAttrs("-UpdaterID")
		res.Meta(&admin.Meta{Name: "UpdaterID", Label: "최종수정자", Type: "readonly", Valuer: func(result interface{}, context *qor.Context) interface{} {
			if updateField := structs.New(result).Field("UpdaterID"); updateField != nil {
				updaterID := updateField.Value().(uint)
				if value, err := GetValueByKeyOfTableName("users", "name", updaterID); err == nil {
					return value.(string)
				}
			}

			return "-"
		}})
	}

	if _, ok := value.(ModelItem); ok {
		res.Meta(&admin.Meta{Name: "Description", Label: "설명"})
		res.Meta(&admin.Meta{Name: "Name", Label: "이름"})
		res.Meta(&admin.Meta{Name: "CreatedAt", Label: "작성일", Type: "readonly"})
		res.Meta(&admin.Meta{Name: "UpdatedAt", Label: "수정일", Type: "readonly"})
		res.EditAttrs("-CreatorID", "-CreatedAt", "-UpdaterID", "-UpdatedAt")
		res.NewAttrs("-CreatorID", "-CreatedAt", "-UpdaterID", "-UpdatedAt")
		SetIndexAttrs(res, appendAttr...)
	}
	return res
}

func SetIndexAttrs(res *admin.Resource, attr ...interface{}) {
	indexAttr := []interface{}{"ID", "Name"}
	indexAttr = append(indexAttr, attr...)
	indexAttr = append(indexAttr, "CreatorID", "CreatedAt", "UpdaterID", "UpdatedAt")
	res.IndexAttrs(indexAttr...)
}

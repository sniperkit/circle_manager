package modules

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set CircleUnitQuerySet

// CircleUnitQuerySet is an queryset type for CircleUnit
type CircleUnitQuerySet struct {
	db *gorm.DB
}

// NewCircleUnitQuerySet constructs new CircleUnitQuerySet
func NewCircleUnitQuerySet(db *gorm.DB) CircleUnitQuerySet {
	return CircleUnitQuerySet{
		db: db.Model(&CircleUnit{}),
	}
}

func (qs CircleUnitQuerySet) w(db *gorm.DB) CircleUnitQuerySet {
	return NewCircleUnitQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) All(ret *[]CircleUnit) error {
	return qs.db.Find(ret).Error
}

// CircleSetIDEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDEq(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id = ?", circleSetID))
}

// CircleSetIDGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDGt(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id > ?", circleSetID))
}

// CircleSetIDGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDGte(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id >= ?", circleSetID))
}

// CircleSetIDIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDIn(circleSetID uint, circleSetIDRest ...uint) CircleUnitQuerySet {
	iArgs := []interface{}{circleSetID}
	for _, arg := range circleSetIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("circle_set_id IN (?)", iArgs))
}

// CircleSetIDLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDLt(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id < ?", circleSetID))
}

// CircleSetIDLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDLte(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id <= ?", circleSetID))
}

// CircleSetIDNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDNe(circleSetID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("circle_set_id != ?", circleSetID))
}

// CircleSetIDNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CircleSetIDNotIn(circleSetID uint, circleSetIDRest ...uint) CircleUnitQuerySet {
	iArgs := []interface{}{circleSetID}
	for _, arg := range circleSetIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("circle_set_id NOT IN (?)", iArgs))
}

// ControllerNameEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ControllerNameEq(controllerName string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("controller_name = ?", controllerName))
}

// ControllerNameIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ControllerNameIn(controllerName string, controllerNameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{controllerName}
	for _, arg := range controllerNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("controller_name IN (?)", iArgs))
}

// ControllerNameNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ControllerNameNe(controllerName string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("controller_name != ?", controllerName))
}

// ControllerNameNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ControllerNameNotIn(controllerName string, controllerNameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{controllerName}
	for _, arg := range controllerNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("controller_name NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *CircleUnit) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtEq(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtGt(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtGte(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtLt(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtLte(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) CreatedAtNe(createdAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *CircleUnit) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) Delete() error {
	return qs.db.Delete(CircleUnit{}).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) DescriptionEq(description string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) DescriptionIn(description string, descriptionRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) DescriptionNe(description string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) DescriptionNotIn(description string, descriptionRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) GetUpdater() CircleUnitUpdater {
	return NewCircleUnitUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDEq(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDGt(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDGte(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDIn(ID uint, IDRest ...uint) CircleUnitQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDLt(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDLte(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDNe(ID uint) CircleUnitQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IDNotIn(ID uint, IDRest ...uint) CircleUnitQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// ImportEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ImportEq(importValue string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("import = ?", importValue))
}

// ImportIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ImportIn(importValue string, importValueRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{importValue}
	for _, arg := range importValueRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("import IN (?)", iArgs))
}

// ImportNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ImportNe(importValue string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("import != ?", importValue))
}

// ImportNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) ImportNotIn(importValue string, importValueRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{importValue}
	for _, arg := range importValueRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("import NOT IN (?)", iArgs))
}

// IsEnableEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsEnableEq(isEnable bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_enable = ?", isEnable))
}

// IsEnableIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsEnableIn(isEnable bool, isEnableRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable IN (?)", iArgs))
}

// IsEnableNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsEnableNe(isEnable bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_enable != ?", isEnable))
}

// IsEnableNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsEnableNotIn(isEnable bool, isEnableRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable NOT IN (?)", iArgs))
}

// IsManualEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsManualEq(isManual bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_manual = ?", isManual))
}

// IsManualIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsManualIn(isManual bool, isManualRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual IN (?)", iArgs))
}

// IsManualNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsManualNe(isManual bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_manual != ?", isManual))
}

// IsManualNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsManualNotIn(isManual bool, isManualRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual NOT IN (?)", iArgs))
}

// IsSystemEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsSystemEq(isSystem bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_system = ?", isSystem))
}

// IsSystemIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsSystemIn(isSystem bool, isSystemRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isSystem}
	for _, arg := range isSystemRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_system IN (?)", iArgs))
}

// IsSystemNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsSystemNe(isSystem bool) CircleUnitQuerySet {
	return qs.w(qs.db.Where("is_system != ?", isSystem))
}

// IsSystemNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) IsSystemNotIn(isSystem bool, isSystemRest ...bool) CircleUnitQuerySet {
	iArgs := []interface{}{isSystem}
	for _, arg := range isSystemRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_system NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) Limit(limit int) CircleUnitQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// MenuGroupEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuGroupEq(menuGroup string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("menu_group = ?", menuGroup))
}

// MenuGroupIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuGroupIn(menuGroup string, menuGroupRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{menuGroup}
	for _, arg := range menuGroupRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("menu_group IN (?)", iArgs))
}

// MenuGroupNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuGroupNe(menuGroup string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("menu_group != ?", menuGroup))
}

// MenuGroupNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuGroupNotIn(menuGroup string, menuGroupRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{menuGroup}
	for _, arg := range menuGroupRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("menu_group NOT IN (?)", iArgs))
}

// MenuNameEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuNameEq(menuName string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("menu_name = ?", menuName))
}

// MenuNameIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuNameIn(menuName string, menuNameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{menuName}
	for _, arg := range menuNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("menu_name IN (?)", iArgs))
}

// MenuNameNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuNameNe(menuName string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("menu_name != ?", menuName))
}

// MenuNameNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) MenuNameNotIn(menuName string, menuNameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{menuName}
	for _, arg := range menuNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("menu_name NOT IN (?)", iArgs))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) NameEq(name string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) NameIn(name string, nameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) NameNe(name string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) NameNotIn(name string, nameRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs CircleUnitQuerySet) One(ret *CircleUnit) error {
	return qs.db.First(ret).Error
}

// OrderAscByCircleSetID is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderAscByCircleSetID() CircleUnitQuerySet {
	return qs.w(qs.db.Order("circle_set_id ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderAscByCreatedAt() CircleUnitQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderAscByID() CircleUnitQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderAscByUpdatedAt() CircleUnitQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCircleSetID is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderDescByCircleSetID() CircleUnitQuerySet {
	return qs.w(qs.db.Order("circle_set_id DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderDescByCreatedAt() CircleUnitQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderDescByID() CircleUnitQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) OrderDescByUpdatedAt() CircleUnitQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadCircleSet is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) PreloadCircleSet() CircleUnitQuerySet {
	return qs.w(qs.db.Preload("CircleSet"))
}

// SetCircleSet is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetCircleSet(circleSet CircleSet) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.CircleSet)] = circleSet
	return u
}

// SetCircleSetID is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetCircleSetID(circleSetID uint) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.CircleSetID)] = circleSetID
	return u
}

// SetControllerName is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetControllerName(controllerName string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.ControllerName)] = controllerName
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetCreatedAt(createdAt time.Time) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetDescription(description string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetID(ID uint) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.ID)] = ID
	return u
}

// SetImport is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetImport(importValue string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.Import)] = importValue
	return u
}

// SetIsEnable is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetIsEnable(isEnable bool) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.IsEnable)] = isEnable
	return u
}

// SetIsManual is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetIsManual(isManual bool) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.IsManual)] = isManual
	return u
}

// SetIsSystem is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetIsSystem(isSystem bool) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.IsSystem)] = isSystem
	return u
}

// SetMenuGroup is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetMenuGroup(menuGroup string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.MenuGroup)] = menuGroup
	return u
}

// SetMenuName is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetMenuName(menuName string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.MenuName)] = menuName
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetName(name string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.Name)] = name
	return u
}

// SetURL is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetURL(URL string) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.URL)] = URL
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) SetUpdatedAt(updatedAt time.Time) CircleUnitUpdater {
	u.fields[string(CircleUnitDBSchema.UpdatedAt)] = updatedAt
	return u
}

// URLEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) URLEq(URL string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("url = ?", URL))
}

// URLIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) URLIn(URL string, URLRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{URL}
	for _, arg := range URLRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("url IN (?)", iArgs))
}

// URLNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) URLNe(URL string) CircleUnitQuerySet {
	return qs.w(qs.db.Where("url != ?", URL))
}

// URLNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) URLNotIn(URL string, URLRest ...string) CircleUnitQuerySet {
	iArgs := []interface{}{URL}
	for _, arg := range URLRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("url NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u CircleUnitUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtEq(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtGt(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtGte(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtLt(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtLte(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitQuerySet) UpdatedAtNe(updatedAt time.Time) CircleUnitQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set CircleUnitQuerySet

// ===== BEGIN of CircleUnit modifiers

type circleUnitDBSchemaField string

func (f circleUnitDBSchemaField) String() string {
	return string(f)
}

// CircleUnitDBSchema stores db field names of CircleUnit
var CircleUnitDBSchema = struct {
	ID             circleUnitDBSchemaField
	CreatedAt      circleUnitDBSchemaField
	UpdatedAt      circleUnitDBSchemaField
	Name           circleUnitDBSchemaField
	Description    circleUnitDBSchemaField
	CircleSet      circleUnitDBSchemaField
	CircleSetID    circleUnitDBSchemaField
	ControllerName circleUnitDBSchemaField
	Import         circleUnitDBSchemaField
	URL            circleUnitDBSchemaField
	MenuName       circleUnitDBSchemaField
	MenuGroup      circleUnitDBSchemaField
	IsEnable       circleUnitDBSchemaField
	IsManual       circleUnitDBSchemaField
	IsSystem       circleUnitDBSchemaField
}{

	ID:             circleUnitDBSchemaField("id"),
	CreatedAt:      circleUnitDBSchemaField("created_at"),
	UpdatedAt:      circleUnitDBSchemaField("updated_at"),
	Name:           circleUnitDBSchemaField("name"),
	Description:    circleUnitDBSchemaField("description"),
	CircleSet:      circleUnitDBSchemaField("circle_set"),
	CircleSetID:    circleUnitDBSchemaField("circle_set_id"),
	ControllerName: circleUnitDBSchemaField("controller_name"),
	Import:         circleUnitDBSchemaField("import"),
	URL:            circleUnitDBSchemaField("url"),
	MenuName:       circleUnitDBSchemaField("menu_name"),
	MenuGroup:      circleUnitDBSchemaField("menu_group"),
	IsEnable:       circleUnitDBSchemaField("is_enable"),
	IsManual:       circleUnitDBSchemaField("is_manual"),
	IsSystem:       circleUnitDBSchemaField("is_system"),
}

// Update updates CircleUnit fields by primary key
func (o *CircleUnit) Update(db *gorm.DB, fields ...circleUnitDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":              o.ID,
		"created_at":      o.CreatedAt,
		"updated_at":      o.UpdatedAt,
		"name":            o.Name,
		"description":     o.Description,
		"circle_set":      o.CircleSet,
		"circle_set_id":   o.CircleSetID,
		"controller_name": o.ControllerName,
		"import":          o.Import,
		"url":             o.URL,
		"menu_name":       o.MenuName,
		"menu_group":      o.MenuGroup,
		"is_enable":       o.IsEnable,
		"is_manual":       o.IsManual,
		"is_system":       o.IsSystem,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update CircleUnit %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// CircleUnitUpdater is an CircleUnit updates manager
type CircleUnitUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewCircleUnitUpdater creates new CircleUnit updater
func NewCircleUnitUpdater(db *gorm.DB) CircleUnitUpdater {
	return CircleUnitUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&CircleUnit{}),
	}
}

// ===== END of CircleUnit modifiers

// ===== END of all query sets

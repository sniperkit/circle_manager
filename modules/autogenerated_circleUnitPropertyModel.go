package modules

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set CircleUnitPropertyQuerySet

// CircleUnitPropertyQuerySet is an queryset type for CircleUnitProperty
type CircleUnitPropertyQuerySet struct {
	db *gorm.DB
}

// NewCircleUnitPropertyQuerySet constructs new CircleUnitPropertyQuerySet
func NewCircleUnitPropertyQuerySet(db *gorm.DB) CircleUnitPropertyQuerySet {
	return CircleUnitPropertyQuerySet{
		db: db.Model(&CircleUnitProperty{}),
	}
}

func (qs CircleUnitPropertyQuerySet) w(db *gorm.DB) CircleUnitPropertyQuerySet {
	return NewCircleUnitPropertyQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) All(ret *[]CircleUnitProperty) error {
	return qs.db.Find(ret).Error
}

// CircleUnitIDEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDEq(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id = ?", circleUnitID))
}

// CircleUnitIDGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDGt(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id > ?", circleUnitID))
}

// CircleUnitIDGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDGte(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id >= ?", circleUnitID))
}

// CircleUnitIDIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDIn(circleUnitID uint, circleUnitIDRest ...uint) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{circleUnitID}
	for _, arg := range circleUnitIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("circle_unit_id IN (?)", iArgs))
}

// CircleUnitIDLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDLt(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id < ?", circleUnitID))
}

// CircleUnitIDLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDLte(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id <= ?", circleUnitID))
}

// CircleUnitIDNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDNe(circleUnitID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("circle_unit_id != ?", circleUnitID))
}

// CircleUnitIDNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CircleUnitIDNotIn(circleUnitID uint, circleUnitIDRest ...uint) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{circleUnitID}
	for _, arg := range circleUnitIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("circle_unit_id NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *CircleUnitProperty) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtEq(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtGt(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtGte(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtLt(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtLte(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) CreatedAtNe(createdAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *CircleUnitProperty) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) Delete() error {
	return qs.db.Delete(CircleUnitProperty{}).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) DescriptionEq(description string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) DescriptionIn(description string, descriptionRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) DescriptionNe(description string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) DescriptionNotIn(description string, descriptionRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) GetUpdater() CircleUnitPropertyUpdater {
	return NewCircleUnitPropertyUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDEq(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDGt(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDGte(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDIn(ID uint, IDRest ...uint) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDLt(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDLte(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDNe(ID uint) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IDNotIn(ID uint, IDRest ...uint) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// IsEnableEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsEnableEq(isEnable bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_enable = ?", isEnable))
}

// IsEnableIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsEnableIn(isEnable bool, isEnableRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable IN (?)", iArgs))
}

// IsEnableNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsEnableNe(isEnable bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_enable != ?", isEnable))
}

// IsEnableNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsEnableNotIn(isEnable bool, isEnableRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable NOT IN (?)", iArgs))
}

// IsManualEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsManualEq(isManual bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_manual = ?", isManual))
}

// IsManualIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsManualIn(isManual bool, isManualRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual IN (?)", iArgs))
}

// IsManualNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsManualNe(isManual bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_manual != ?", isManual))
}

// IsManualNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsManualNotIn(isManual bool, isManualRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual NOT IN (?)", iArgs))
}

// IsSystemEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsSystemEq(isSystem bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_system = ?", isSystem))
}

// IsSystemIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsSystemIn(isSystem bool, isSystemRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isSystem}
	for _, arg := range isSystemRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_system IN (?)", iArgs))
}

// IsSystemNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsSystemNe(isSystem bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("is_system != ?", isSystem))
}

// IsSystemNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) IsSystemNotIn(isSystem bool, isSystemRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{isSystem}
	for _, arg := range isSystemRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_system NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) Limit(limit int) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NameEq(name string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NameIn(name string, nameRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NameNe(name string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NameNotIn(name string, nameRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// NullableEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NullableEq(nullable bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("nullable = ?", nullable))
}

// NullableIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NullableIn(nullable bool, nullableRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{nullable}
	for _, arg := range nullableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("nullable IN (?)", iArgs))
}

// NullableNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NullableNe(nullable bool) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("nullable != ?", nullable))
}

// NullableNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) NullableNotIn(nullable bool, nullableRest ...bool) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{nullable}
	for _, arg := range nullableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("nullable NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs CircleUnitPropertyQuerySet) One(ret *CircleUnitProperty) error {
	return qs.db.First(ret).Error
}

// OrderAscByCircleUnitID is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderAscByCircleUnitID() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("circle_unit_id ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderAscByCreatedAt() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderAscByID() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderAscByUpdatedAt() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCircleUnitID is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderDescByCircleUnitID() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("circle_unit_id DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderDescByCreatedAt() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderDescByID() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) OrderDescByUpdatedAt() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadCircleUnit is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) PreloadCircleUnit() CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Preload("CircleUnit"))
}

// SetCircleUnit is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetCircleUnit(circleUnit CircleUnit) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.CircleUnit)] = circleUnit
	return u
}

// SetCircleUnitID is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetCircleUnitID(circleUnitID uint) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.CircleUnitID)] = circleUnitID
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetCreatedAt(createdAt time.Time) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetDescription(description string) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetID(ID uint) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.ID)] = ID
	return u
}

// SetIsEnable is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetIsEnable(isEnable bool) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.IsEnable)] = isEnable
	return u
}

// SetIsManual is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetIsManual(isManual bool) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.IsManual)] = isManual
	return u
}

// SetIsSystem is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetIsSystem(isSystem bool) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.IsSystem)] = isSystem
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetName(name string) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.Name)] = name
	return u
}

// SetNullable is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetNullable(nullable bool) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.Nullable)] = nullable
	return u
}

// SetType is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetType(typeValue string) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.Type)] = typeValue
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) SetUpdatedAt(updatedAt time.Time) CircleUnitPropertyUpdater {
	u.fields[string(CircleUnitPropertyDBSchema.UpdatedAt)] = updatedAt
	return u
}

// TypeEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) TypeEq(typeValue string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("type = ?", typeValue))
}

// TypeIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) TypeIn(typeValue string, typeValueRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{typeValue}
	for _, arg := range typeValueRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("type IN (?)", iArgs))
}

// TypeNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) TypeNe(typeValue string) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("type != ?", typeValue))
}

// TypeNotIn is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) TypeNotIn(typeValue string, typeValueRest ...string) CircleUnitPropertyQuerySet {
	iArgs := []interface{}{typeValue}
	for _, arg := range typeValueRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("type NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u CircleUnitPropertyUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtEq(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtGt(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtGte(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtLt(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtLte(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs CircleUnitPropertyQuerySet) UpdatedAtNe(updatedAt time.Time) CircleUnitPropertyQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set CircleUnitPropertyQuerySet

// ===== BEGIN of CircleUnitProperty modifiers

type circleUnitPropertyDBSchemaField string

func (f circleUnitPropertyDBSchemaField) String() string {
	return string(f)
}

// CircleUnitPropertyDBSchema stores db field names of CircleUnitProperty
var CircleUnitPropertyDBSchema = struct {
	ID           circleUnitPropertyDBSchemaField
	CreatedAt    circleUnitPropertyDBSchemaField
	UpdatedAt    circleUnitPropertyDBSchemaField
	Name         circleUnitPropertyDBSchemaField
	Description  circleUnitPropertyDBSchemaField
	CircleUnit   circleUnitPropertyDBSchemaField
	CircleUnitID circleUnitPropertyDBSchemaField
	Type         circleUnitPropertyDBSchemaField
	Nullable     circleUnitPropertyDBSchemaField
	IsEnable     circleUnitPropertyDBSchemaField
	IsManual     circleUnitPropertyDBSchemaField
	IsSystem     circleUnitPropertyDBSchemaField
}{

	ID:           circleUnitPropertyDBSchemaField("id"),
	CreatedAt:    circleUnitPropertyDBSchemaField("created_at"),
	UpdatedAt:    circleUnitPropertyDBSchemaField("updated_at"),
	Name:         circleUnitPropertyDBSchemaField("name"),
	Description:  circleUnitPropertyDBSchemaField("description"),
	CircleUnit:   circleUnitPropertyDBSchemaField("circle_unit"),
	CircleUnitID: circleUnitPropertyDBSchemaField("circle_unit_id"),
	Type:         circleUnitPropertyDBSchemaField("type"),
	Nullable:     circleUnitPropertyDBSchemaField("nullable"),
	IsEnable:     circleUnitPropertyDBSchemaField("is_enable"),
	IsManual:     circleUnitPropertyDBSchemaField("is_manual"),
	IsSystem:     circleUnitPropertyDBSchemaField("is_system"),
}

// Update updates CircleUnitProperty fields by primary key
func (o *CircleUnitProperty) Update(db *gorm.DB, fields ...circleUnitPropertyDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":             o.ID,
		"created_at":     o.CreatedAt,
		"updated_at":     o.UpdatedAt,
		"name":           o.Name,
		"description":    o.Description,
		"circle_unit":    o.CircleUnit,
		"circle_unit_id": o.CircleUnitID,
		"type":           o.Type,
		"nullable":       o.Nullable,
		"is_enable":      o.IsEnable,
		"is_manual":      o.IsManual,
		"is_system":      o.IsSystem,
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

		return fmt.Errorf("can't update CircleUnitProperty %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// CircleUnitPropertyUpdater is an CircleUnitProperty updates manager
type CircleUnitPropertyUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewCircleUnitPropertyUpdater creates new CircleUnitProperty updater
func NewCircleUnitPropertyUpdater(db *gorm.DB) CircleUnitPropertyUpdater {
	return CircleUnitPropertyUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&CircleUnitProperty{}),
	}
}

// ===== END of CircleUnitProperty modifiers

// ===== END of all query sets

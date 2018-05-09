package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SprintQuerySet

// SprintQuerySet is an queryset type for Sprint
type SprintQuerySet struct {
	db *gorm.DB
}

// NewSprintQuerySet constructs new SprintQuerySet
func NewSprintQuerySet(db *gorm.DB) SprintQuerySet {
	return SprintQuerySet{
		db: db.Model(&Sprint{}),
	}
}

func (qs SprintQuerySet) w(db *gorm.DB) SprintQuerySet {
	return NewSprintQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) All(ret *[]Sprint) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Sprint) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtEq(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtGt(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtGte(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtLt(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtLte(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CreatedAtNe(createdAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CurrentEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CurrentEq(current bool) SprintQuerySet {
	return qs.w(qs.db.Where("current = ?", current))
}

// CurrentIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CurrentIn(current bool, currentRest ...bool) SprintQuerySet {
	iArgs := []interface{}{current}
	for _, arg := range currentRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("current IN (?)", iArgs))
}

// CurrentNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CurrentNe(current bool) SprintQuerySet {
	return qs.w(qs.db.Where("current != ?", current))
}

// CurrentNotIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) CurrentNotIn(current bool, currentRest ...bool) SprintQuerySet {
	iArgs := []interface{}{current}
	for _, arg := range currentRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("current NOT IN (?)", iArgs))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Sprint) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) Delete() error {
	return qs.db.Delete(Sprint{}).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) DescriptionEq(description string) SprintQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) DescriptionIn(description string, descriptionRest ...string) SprintQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) DescriptionNe(description string) SprintQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) DescriptionNotIn(description string, descriptionRest ...string) SprintQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) GetUpdater() SprintUpdater {
	return NewSprintUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDEq(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDGt(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDGte(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDIn(ID uint, IDRest ...uint) SprintQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDLt(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDLte(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDNe(ID uint) SprintQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) IDNotIn(ID uint, IDRest ...uint) SprintQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) Limit(limit int) SprintQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) NameEq(name string) SprintQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) NameIn(name string, nameRest ...string) SprintQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) NameNe(name string) SprintQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) NameNotIn(name string, nameRest ...string) SprintQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SprintQuerySet) One(ret *Sprint) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderAscByCreatedAt() SprintQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderAscByID() SprintQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderAscByUpdatedAt() SprintQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderDescByCreatedAt() SprintQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderDescByID() SprintQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) OrderDescByUpdatedAt() SprintQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetCreatedAt(createdAt time.Time) SprintUpdater {
	u.fields[string(SprintDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCurrent is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetCurrent(current bool) SprintUpdater {
	u.fields[string(SprintDBSchema.Current)] = current
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetDescription(description string) SprintUpdater {
	u.fields[string(SprintDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetID(ID uint) SprintUpdater {
	u.fields[string(SprintDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetName(name string) SprintUpdater {
	u.fields[string(SprintDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u SprintUpdater) SetUpdatedAt(updatedAt time.Time) SprintUpdater {
	u.fields[string(SprintDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u SprintUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SprintUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtEq(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtGt(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtGte(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtLt(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtLte(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs SprintQuerySet) UpdatedAtNe(updatedAt time.Time) SprintQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set SprintQuerySet

// ===== BEGIN of Sprint modifiers

type sprintDBSchemaField string

func (f sprintDBSchemaField) String() string {
	return string(f)
}

// SprintDBSchema stores db field names of Sprint
var SprintDBSchema = struct {
	ID          sprintDBSchemaField
	CreatedAt   sprintDBSchemaField
	UpdatedAt   sprintDBSchemaField
	Name        sprintDBSchemaField
	Description sprintDBSchemaField
	Current     sprintDBSchemaField
}{

	ID:          sprintDBSchemaField("id"),
	CreatedAt:   sprintDBSchemaField("created_at"),
	UpdatedAt:   sprintDBSchemaField("updated_at"),
	Name:        sprintDBSchemaField("name"),
	Description: sprintDBSchemaField("description"),
	Current:     sprintDBSchemaField("current"),
}

// Update updates Sprint fields by primary key
func (o *Sprint) Update(db *gorm.DB, fields ...sprintDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"name":        o.Name,
		"description": o.Description,
		"current":     o.Current,
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

		return fmt.Errorf("can't update Sprint %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SprintUpdater is an Sprint updates manager
type SprintUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSprintUpdater creates new Sprint updater
func NewSprintUpdater(db *gorm.DB) SprintUpdater {
	return SprintUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Sprint{}),
	}
}

// ===== END of Sprint modifiers

// ===== END of all query sets

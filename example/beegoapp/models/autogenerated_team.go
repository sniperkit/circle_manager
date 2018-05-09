package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set TeamQuerySet

// TeamQuerySet is an queryset type for Team
type TeamQuerySet struct {
	db *gorm.DB
}

// NewTeamQuerySet constructs new TeamQuerySet
func NewTeamQuerySet(db *gorm.DB) TeamQuerySet {
	return TeamQuerySet{
		db: db.Model(&Team{}),
	}
}

func (qs TeamQuerySet) w(db *gorm.DB) TeamQuerySet {
	return NewTeamQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) All(ret *[]Team) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Team) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtEq(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtGt(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtGte(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtLt(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtLte(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) CreatedAtNe(createdAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Team) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) Delete() error {
	return qs.db.Delete(Team{}).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) DescriptionEq(description string) TeamQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) DescriptionIn(description string, descriptionRest ...string) TeamQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) DescriptionNe(description string) TeamQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) DescriptionNotIn(description string, descriptionRest ...string) TeamQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) GetUpdater() TeamUpdater {
	return NewTeamUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDEq(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDGt(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDGte(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDIn(ID uint, IDRest ...uint) TeamQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDLt(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDLte(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDNe(ID uint) TeamQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) IDNotIn(ID uint, IDRest ...uint) TeamQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) Limit(limit int) TeamQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) NameEq(name string) TeamQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) NameIn(name string, nameRest ...string) TeamQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) NameNe(name string) TeamQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) NameNotIn(name string, nameRest ...string) TeamQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs TeamQuerySet) One(ret *Team) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderAscByCreatedAt() TeamQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderAscByID() TeamQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderAscByUpdatedAt() TeamQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderDescByCreatedAt() TeamQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderDescByID() TeamQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) OrderDescByUpdatedAt() TeamQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u TeamUpdater) SetCreatedAt(createdAt time.Time) TeamUpdater {
	u.fields[string(TeamDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u TeamUpdater) SetDescription(description string) TeamUpdater {
	u.fields[string(TeamDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u TeamUpdater) SetID(ID uint) TeamUpdater {
	u.fields[string(TeamDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u TeamUpdater) SetName(name string) TeamUpdater {
	u.fields[string(TeamDBSchema.Name)] = name
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u TeamUpdater) SetUpdatedAt(updatedAt time.Time) TeamUpdater {
	u.fields[string(TeamDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u TeamUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u TeamUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtEq(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtGt(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtGte(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtLt(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtLte(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs TeamQuerySet) UpdatedAtNe(updatedAt time.Time) TeamQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set TeamQuerySet

// ===== BEGIN of Team modifiers

type teamDBSchemaField string

func (f teamDBSchemaField) String() string {
	return string(f)
}

// TeamDBSchema stores db field names of Team
var TeamDBSchema = struct {
	ID          teamDBSchemaField
	CreatedAt   teamDBSchemaField
	UpdatedAt   teamDBSchemaField
	Name        teamDBSchemaField
	Description teamDBSchemaField
}{

	ID:          teamDBSchemaField("id"),
	CreatedAt:   teamDBSchemaField("created_at"),
	UpdatedAt:   teamDBSchemaField("updated_at"),
	Name:        teamDBSchemaField("name"),
	Description: teamDBSchemaField("description"),
}

// Update updates Team fields by primary key
func (o *Team) Update(db *gorm.DB, fields ...teamDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"name":        o.Name,
		"description": o.Description,
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

		return fmt.Errorf("can't update Team %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// TeamUpdater is an Team updates manager
type TeamUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewTeamUpdater creates new Team updater
func NewTeamUpdater(db *gorm.DB) TeamUpdater {
	return TeamUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Team{}),
	}
}

// ===== END of Team modifiers

// ===== END of all query sets
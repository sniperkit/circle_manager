package modules

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set NotificationTypeQuerySet

// NotificationTypeQuerySet is an queryset type for NotificationType
type NotificationTypeQuerySet struct {
	db *gorm.DB
}

// NewNotificationTypeQuerySet constructs new NotificationTypeQuerySet
func NewNotificationTypeQuerySet(db *gorm.DB) NotificationTypeQuerySet {
	return NotificationTypeQuerySet{
		db: db.Model(&NotificationType{}),
	}
}

func (qs NotificationTypeQuerySet) w(db *gorm.DB) NotificationTypeQuerySet {
	return NewNotificationTypeQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) All(ret *[]NotificationType) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *NotificationType) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtEq(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtGt(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtGte(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtLt(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtLte(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) CreatedAtNe(createdAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) Delete() error {
	return qs.db.Delete(NotificationType{}).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *NotificationType) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) DescriptionEq(description string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) DescriptionIn(description string, descriptionRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) DescriptionNe(description string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) DescriptionNotIn(description string, descriptionRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) GetUpdater() NotificationTypeUpdater {
	return NewNotificationTypeUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDEq(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDGt(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDGte(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDIn(ID uint, IDRest ...uint) NotificationTypeQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDLt(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDLte(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDNe(ID uint) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IDNotIn(ID uint, IDRest ...uint) NotificationTypeQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// IsEnableEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsEnableEq(isEnable bool) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("is_enable = ?", isEnable))
}

// IsEnableIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsEnableIn(isEnable bool, isEnableRest ...bool) NotificationTypeQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable IN (?)", iArgs))
}

// IsEnableNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsEnableNe(isEnable bool) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("is_enable != ?", isEnable))
}

// IsEnableNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsEnableNotIn(isEnable bool, isEnableRest ...bool) NotificationTypeQuerySet {
	iArgs := []interface{}{isEnable}
	for _, arg := range isEnableRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_enable NOT IN (?)", iArgs))
}

// IsManualEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsManualEq(isManual bool) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("is_manual = ?", isManual))
}

// IsManualIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsManualIn(isManual bool, isManualRest ...bool) NotificationTypeQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual IN (?)", iArgs))
}

// IsManualNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsManualNe(isManual bool) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("is_manual != ?", isManual))
}

// IsManualNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) IsManualNotIn(isManual bool, isManualRest ...bool) NotificationTypeQuerySet {
	iArgs := []interface{}{isManual}
	for _, arg := range isManualRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("is_manual NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) Limit(limit int) NotificationTypeQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// ListItemTemplateEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ListItemTemplateEq(listItemTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("list_item_template = ?", listItemTemplate))
}

// ListItemTemplateIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ListItemTemplateIn(listItemTemplate string, listItemTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{listItemTemplate}
	for _, arg := range listItemTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("list_item_template IN (?)", iArgs))
}

// ListItemTemplateNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ListItemTemplateNe(listItemTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("list_item_template != ?", listItemTemplate))
}

// ListItemTemplateNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ListItemTemplateNotIn(listItemTemplate string, listItemTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{listItemTemplate}
	for _, arg := range listItemTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("list_item_template NOT IN (?)", iArgs))
}

// MessageTemplateEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) MessageTemplateEq(messageTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("message_template = ?", messageTemplate))
}

// MessageTemplateIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) MessageTemplateIn(messageTemplate string, messageTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{messageTemplate}
	for _, arg := range messageTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("message_template IN (?)", iArgs))
}

// MessageTemplateNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) MessageTemplateNe(messageTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("message_template != ?", messageTemplate))
}

// MessageTemplateNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) MessageTemplateNotIn(messageTemplate string, messageTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{messageTemplate}
	for _, arg := range messageTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("message_template NOT IN (?)", iArgs))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) NameEq(name string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) NameIn(name string, nameRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) NameNe(name string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) NameNotIn(name string, nameRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs NotificationTypeQuerySet) One(ret *NotificationType) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderAscByCreatedAt() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderAscByID() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderAscByUpdatedAt() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderDescByCreatedAt() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderDescByID() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) OrderDescByUpdatedAt() NotificationTypeQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// ReplaceTextEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ReplaceTextEq(replaceText string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("replace_text = ?", replaceText))
}

// ReplaceTextIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ReplaceTextIn(replaceText string, replaceTextRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{replaceText}
	for _, arg := range replaceTextRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("replace_text IN (?)", iArgs))
}

// ReplaceTextNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ReplaceTextNe(replaceText string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("replace_text != ?", replaceText))
}

// ReplaceTextNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) ReplaceTextNotIn(replaceText string, replaceTextRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{replaceText}
	for _, arg := range replaceTextRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("replace_text NOT IN (?)", iArgs))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetCreatedAt(createdAt time.Time) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetDescription(description string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.Description)] = description
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetID(ID uint) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.ID)] = ID
	return u
}

// SetIsEnable is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetIsEnable(isEnable bool) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.IsEnable)] = isEnable
	return u
}

// SetIsManual is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetIsManual(isManual bool) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.IsManual)] = isManual
	return u
}

// SetListItemTemplate is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetListItemTemplate(listItemTemplate string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.ListItemTemplate)] = listItemTemplate
	return u
}

// SetMessageTemplate is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetMessageTemplate(messageTemplate string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.MessageTemplate)] = messageTemplate
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetName(name string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.Name)] = name
	return u
}

// SetReplaceText is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetReplaceText(replaceText string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.ReplaceText)] = replaceText
	return u
}

// SetTags is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetTags(tags string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.Tags)] = tags
	return u
}

// SetTargetObject is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetTargetObject(targetObject string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.TargetObject)] = targetObject
	return u
}

// SetTargetWhere is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetTargetWhere(targetWhere string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.TargetWhere)] = targetWhere
	return u
}

// SetTitleTemplate is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetTitleTemplate(titleTemplate string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.TitleTemplate)] = titleTemplate
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetUpdatedAt(updatedAt time.Time) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetWebhookURLs is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) SetWebhookURLs(webhookURLs string) NotificationTypeUpdater {
	u.fields[string(NotificationTypeDBSchema.WebhookURLs)] = webhookURLs
	return u
}

// TagsEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TagsEq(tags string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("tags = ?", tags))
}

// TagsIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TagsIn(tags string, tagsRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{tags}
	for _, arg := range tagsRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("tags IN (?)", iArgs))
}

// TagsNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TagsNe(tags string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("tags != ?", tags))
}

// TagsNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TagsNotIn(tags string, tagsRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{tags}
	for _, arg := range tagsRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("tags NOT IN (?)", iArgs))
}

// TargetObjectEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetObjectEq(targetObject string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("target_object = ?", targetObject))
}

// TargetObjectIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetObjectIn(targetObject string, targetObjectRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{targetObject}
	for _, arg := range targetObjectRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("target_object IN (?)", iArgs))
}

// TargetObjectNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetObjectNe(targetObject string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("target_object != ?", targetObject))
}

// TargetObjectNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetObjectNotIn(targetObject string, targetObjectRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{targetObject}
	for _, arg := range targetObjectRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("target_object NOT IN (?)", iArgs))
}

// TargetWhereEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetWhereEq(targetWhere string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("target_where = ?", targetWhere))
}

// TargetWhereIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetWhereIn(targetWhere string, targetWhereRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{targetWhere}
	for _, arg := range targetWhereRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("target_where IN (?)", iArgs))
}

// TargetWhereNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetWhereNe(targetWhere string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("target_where != ?", targetWhere))
}

// TargetWhereNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TargetWhereNotIn(targetWhere string, targetWhereRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{targetWhere}
	for _, arg := range targetWhereRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("target_where NOT IN (?)", iArgs))
}

// TitleTemplateEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TitleTemplateEq(titleTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("title_template = ?", titleTemplate))
}

// TitleTemplateIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TitleTemplateIn(titleTemplate string, titleTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{titleTemplate}
	for _, arg := range titleTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("title_template IN (?)", iArgs))
}

// TitleTemplateNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TitleTemplateNe(titleTemplate string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("title_template != ?", titleTemplate))
}

// TitleTemplateNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) TitleTemplateNotIn(titleTemplate string, titleTemplateRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{titleTemplate}
	for _, arg := range titleTemplateRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("title_template NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u NotificationTypeUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtEq(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtGt(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtGte(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtLt(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtLte(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) UpdatedAtNe(updatedAt time.Time) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// WebhookURLsEq is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) WebhookURLsEq(webhookURLs string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("webhook_urls = ?", webhookURLs))
}

// WebhookURLsIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) WebhookURLsIn(webhookURLs string, webhookURLsRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{webhookURLs}
	for _, arg := range webhookURLsRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("webhook_urls IN (?)", iArgs))
}

// WebhookURLsNe is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) WebhookURLsNe(webhookURLs string) NotificationTypeQuerySet {
	return qs.w(qs.db.Where("webhook_urls != ?", webhookURLs))
}

// WebhookURLsNotIn is an autogenerated method
// nolint: dupl
func (qs NotificationTypeQuerySet) WebhookURLsNotIn(webhookURLs string, webhookURLsRest ...string) NotificationTypeQuerySet {
	iArgs := []interface{}{webhookURLs}
	for _, arg := range webhookURLsRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("webhook_urls NOT IN (?)", iArgs))
}

// ===== END of query set NotificationTypeQuerySet

// ===== BEGIN of NotificationType modifiers

type notificationTypeDBSchemaField string

func (f notificationTypeDBSchemaField) String() string {
	return string(f)
}

// NotificationTypeDBSchema stores db field names of NotificationType
var NotificationTypeDBSchema = struct {
	ID               notificationTypeDBSchemaField
	CreatedAt        notificationTypeDBSchemaField
	UpdatedAt        notificationTypeDBSchemaField
	Name             notificationTypeDBSchemaField
	Description      notificationTypeDBSchemaField
	IsEnable         notificationTypeDBSchemaField
	IsManual         notificationTypeDBSchemaField
	TargetObject     notificationTypeDBSchemaField
	TargetWhere      notificationTypeDBSchemaField
	Tags             notificationTypeDBSchemaField
	TitleTemplate    notificationTypeDBSchemaField
	MessageTemplate  notificationTypeDBSchemaField
	ListItemTemplate notificationTypeDBSchemaField
	WebhookURLs      notificationTypeDBSchemaField
	ReplaceText      notificationTypeDBSchemaField
}{

	ID:               notificationTypeDBSchemaField("id"),
	CreatedAt:        notificationTypeDBSchemaField("created_at"),
	UpdatedAt:        notificationTypeDBSchemaField("updated_at"),
	Name:             notificationTypeDBSchemaField("name"),
	Description:      notificationTypeDBSchemaField("description"),
	IsEnable:         notificationTypeDBSchemaField("is_enable"),
	IsManual:         notificationTypeDBSchemaField("is_manual"),
	TargetObject:     notificationTypeDBSchemaField("target_object"),
	TargetWhere:      notificationTypeDBSchemaField("target_where"),
	Tags:             notificationTypeDBSchemaField("tags"),
	TitleTemplate:    notificationTypeDBSchemaField("title_template"),
	MessageTemplate:  notificationTypeDBSchemaField("message_template"),
	ListItemTemplate: notificationTypeDBSchemaField("list_item_template"),
	WebhookURLs:      notificationTypeDBSchemaField("webhook_urls"),
	ReplaceText:      notificationTypeDBSchemaField("replace_text"),
}

// Update updates NotificationType fields by primary key
func (o *NotificationType) Update(db *gorm.DB, fields ...notificationTypeDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                 o.ID,
		"created_at":         o.CreatedAt,
		"updated_at":         o.UpdatedAt,
		"name":               o.Name,
		"description":        o.Description,
		"is_enable":          o.IsEnable,
		"is_manual":          o.IsManual,
		"target_object":      o.TargetObject,
		"target_where":       o.TargetWhere,
		"tags":               o.Tags,
		"title_template":     o.TitleTemplate,
		"message_template":   o.MessageTemplate,
		"list_item_template": o.ListItemTemplate,
		"webhook_urls":       o.WebhookURLs,
		"replace_text":       o.ReplaceText,
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

		return fmt.Errorf("can't update NotificationType %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// NotificationTypeUpdater is an NotificationType updates manager
type NotificationTypeUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewNotificationTypeUpdater creates new NotificationType updater
func NewNotificationTypeUpdater(db *gorm.DB) NotificationTypeUpdater {
	return NotificationTypeUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&NotificationType{}),
	}
}

// ===== END of NotificationType modifiers

// ===== END of all query sets

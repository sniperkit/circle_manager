package modules

//  NotificationTypeController operations for NotificationType
type NotificationTypeController struct {
	BaseUserController
}

// Post ...
// @Title Post
// @Description create NotificationType
// @Param	body		body 	modules.NotificationType	true		"body for NotificationType content"
// @Success 201 {int} modules.NotificationType
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *NotificationTypeController) Post() {
	c.BasePost(&CreateNotificationType{}, &NotificationType{})
}

// GetOne ...
// @Title Get One
// @Description get NotificationType by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} modules.NotificationType
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *NotificationTypeController) GetOne() {
	c.BaseGetOne(&NotificationType{})
}

// GetAll ...
// @Title Get All
// @Description get NotificationType
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []modules.NotificationType
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *NotificationTypeController) GetAll() {
	c.BaseGetAll([]NotificationType{})
}

// Put ...
// @Title Put
// @Description update the NotificationType
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	modules.NotificationType	true		"body for NotificationType content"
// @Success 200 {object} modules.NotificationType
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *NotificationTypeController) Put() {
	c.BasePut(&UpdateNotificationType{}, &NotificationType{})
}

// Delete ...
// @Title Delete
// @Description delete the NotificationType
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *NotificationTypeController) Delete() {
	c.BaseDelete(&NotificationType{})
}

type CreateNotificationType struct {
	Name             string
	Description      string
	ManualSend       bool
	TargetObject     string
	TargetAction     string
	IsEnable         bool
	UseLink          bool
	TitleTemplate    string
	MessageTemplate  string
	ListItemTemplate string
}

type UpdateNotificationType struct {
	Name             string
	Description      string
	ManualSend       bool
	TargetObject     string
	TargetAction     string
	IsEnable         bool
	UseLink          bool
	TitleTemplate    string
	MessageTemplate  string
	ListItemTemplate string
}

func (c *CreateNotificationType) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateNotificationType) Valid() error {
	return validate.Struct(c)
}

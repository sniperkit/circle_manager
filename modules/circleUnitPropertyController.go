package modules

//  CircleUnitPropertyController operations for CircleUnit
type CircleUnitPropertyController struct {
	BaseUserController
}

func (c *CircleUnitPropertyController) Prepare() {
	c.RequestCreateItem = &CreateCircleUnit{}
	c.RequestUpdateItem = &UpdateCircleUnit{}
	c.ModelItem = &CircleUnit{}
	c.ModelItems = &[]CircleUnit{}
	c.ResponseItem = &CircleUnit{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create CircleUnit
// @Param	body		body 	modules.CircleUnit	true		"body for CircleUnit content"
// @Success 201 {int} modules.CircleUnit
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *CircleUnitPropertyController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get CircleUnit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} modules.CircleUnit
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *CircleUnitPropertyController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get CircleUnit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []modules.CircleUnit
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *CircleUnitPropertyController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the CircleUnit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	modules.CircleUnit	true		"body for CircleUnit content"
// @Success 200 {object} modules.CircleUnit
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *CircleUnitPropertyController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the CircleUnit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *CircleUnitPropertyController) Delete() {
	c.BaseDelete()
}

type CreateCircleUnitProperty struct {
	Name         string
	Description  string
	CircleUnitID uint
	Type         string
	Nullable     bool
	IsEnable     bool
	IsManual     bool
	IsSystem     bool
}

type UpdateCircleUnitProperty struct {
	Name         string
	Description  string
	CircleUnitID uint
	Type         string
	Nullable     bool
	IsEnable     bool
	IsManual     bool
	IsSystem     bool
}

func (c *CreateCircleUnitProperty) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateCircleUnitProperty) Valid() error {
	return validate.Struct(c)
}

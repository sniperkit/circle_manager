package modules

//  CircleUnitController operations for CircleUnit
type CircleUnitController struct {
	BaseUserController
}

// Post ...
// @Title Post
// @Description create CircleUnit
// @Param	body		body 	modules.CreateCircleUnit	true		"body for CircleUnit content"
// @Success 201 {int} modules.CircleUnit
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *CircleUnitController) Post() {
	c.BasePost(&CreateCircleUnit{}, &CircleUnit{})
}

// GetOne ...
// @Title Get One
// @Description get CircleUnit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} modules.CircleUnit
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *CircleUnitController) GetOne() {
	c.BaseGetOne(&CircleUnit{})
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
func (c *CircleUnitController) GetAll() {
	c.BaseGetAll(&[]CircleUnit{})
}

// Put ...
// @Title Put
// @Description update the CircleUnit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	modules.UpdateCircleUnit	true		"body for CircleUnit content"
// @Success 200 {object} modules.CircleUnit
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *CircleUnitController) Put() {
	c.BasePut(&UpdateCircleUnit{}, &CircleUnit{})
}

// Delete ...
// @Title Delete
// @Description delete the CircleUnit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *CircleUnitController) Delete() {
	c.BaseDelete(&CircleUnit{})
}

type CreateCircleUnit struct {
	Name           string
	Description    string
	CircleSetID    uint
	ControllerName string
	Import         string
	URL            string
	MenuName       string
	MenuGroup      string
	IsEnable       bool
	IsManual       bool
	IsSystem       bool
}

type UpdateCircleUnit struct {
	Name           string
	Description    string
	CircleSetID    uint
	ControllerName string
	Import         string
	URL            string
	MenuName       string
	MenuGroup      string
	IsEnable       bool
	IsManual       bool
	IsSystem       bool
}

func (c *CreateCircleUnit) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateCircleUnit) Valid() error {
	return validate.Struct(c)
}

package modules

//  CircleSetController operations for CircleSet
type CircleSetController struct {
	BaseUserController
}

func (c *CircleSetController) Prepare() {
	c.RequestCreateItem = &CreateCircleSet{}
	c.RequestUpdateItem = &UpdateCircleSet{}
	c.ModelItem = &CircleSet{}
	c.ModelItems = &[]CircleSet{}
	c.ResponseItem = &CircleSet{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create CircleSet
// @Param	body		body 	modules.CircleSet	true		"body for CircleSet content"
// @Success 201 {int} modules.CircleSet
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *CircleSetController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get modules.CircleSet by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} modules.CircleSet
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *CircleSetController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get CircleSet
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []modules.CircleSet
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *CircleSetController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the CircleSet
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	modules.CircleSet	true		"body for CircleSet content"
// @Success 200 {object} modules.CircleSet
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *CircleSetController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the CircleSet
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *CircleSetController) Delete() {
	c.BaseDelete()
}

type CreateCircleSet struct {
	Name                  string
	Description           string
	Import                string
	IsEnable              bool
	AppVersion            string
	AppTitle              string
	AppDescription        string
	AppContact            string
	AppTermsOfServiceUrl  string
	AppLicense            string
	AppSecurityDefinition string
	RunAppEnvs            string
}

type UpdateCircleSet struct {
	Name                  string
	Description           string
	Import                string
	IsEnable              bool
	AppVersion            string
	AppTitle              string
	AppDescription        string
	AppContact            string
	AppTermsOfServiceUrl  string
	AppLicense            string
	AppSecurityDefinition string
	RunAppEnvs            string
}

func (c *CreateCircleSet) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateCircleSet) Valid() error {
	return validate.Struct(c)
}

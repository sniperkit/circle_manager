package controllers

import (
	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/_example/beegoapp/requests"
	"github.com/jungju/circle_manager/_example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

//  TeamController operations for Team
type TeamController struct {
	modules.BaseUserController
}

func (c *TeamController) Prepare() {
	c.RequestCreateItem = &requests.CreateTeam{}
	c.RequestUpdateItem = &requests.UpdateTeam{}
	c.ModelItem = &models.Team{}
	c.ModelItems = &[]models.Team{}
	c.ResponseItem = &responses.Team{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create Team
// @Param	body		body 	models.Team	true		"body for Team content"
// @Success 201 {object} responses.Team
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *TeamController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get Team by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Team
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *TeamController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get Team
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Team
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *TeamController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the Team
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Team	true		"body for Team content"
// @Success 200 {object} responses.Team
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *TeamController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the Team
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *TeamController) Delete() {
	c.BaseDelete()
}

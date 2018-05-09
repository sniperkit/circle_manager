package controllers

import (
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/requests"
	"github.com/jungju/circle_manager/example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

//  IcsController operations for Ics
type IcsController struct {
	modules.BaseUserController
}

func (c *IcsController) Prepare() {
	c.RequestCreateItem = &requests.CreateIcs{}
	c.RequestUpdateItem = &requests.UpdateIcs{}
	c.ModelItem = &models.Ics{}
	c.ModelItems = &[]models.Ics{}
	c.ResponseItem = &responses.Ics{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create Ics
// @Param	body		body 	models.Ics	true		"body for Ics content"
// @Success 201 {object} responses.Ics
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *IcsController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get Ics by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Ics
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *IcsController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get Ics
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Ics
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *IcsController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the Ics
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ics	true		"body for Ics content"
// @Success 200 {object} responses.Ics
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *IcsController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the Ics
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *IcsController) Delete() {
	c.BaseDelete()
}

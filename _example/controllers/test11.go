package controllers

import (
	"github.com//models"
	"github.com//requests"
	"github.com//responses"
	"github.com/jungju/circle_manager/modules"
)

//  Test11Controller operations for Test11
type Test11Controller struct {
	modules.BaseUserController
}

func (c *Test11Controller) Prepare() {
	c.RequestCreateItem = &requests.CreateTest11{}
	c.RequestUpdateItem = &requests.UpdateTest11{}
	c.ModelItem = &models.Test11{}
	c.ModelItems = &[]models.Test11{}
	c.ResponseItem = &responses.Test11{}
}

// Post ...
// @Title Post
// @Description create Test11
// @Param	body		body 	models.Test11	true		"body for Test11 content"
// @Success 201 {int} responses.Test11
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *Test11Controller) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get Test11 by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Test11
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *Test11Controller) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get Test11
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.Test11
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *Test11Controller) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the Test11
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Test11	true		"body for Test11 content"
// @Success 200 {object} responses.Test11
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *Test11Controller) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the Test11
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *Test11Controller) Delete() {
	c.BaseDelete()
}

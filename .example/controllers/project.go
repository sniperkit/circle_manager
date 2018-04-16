package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
)

//  ProjectController operations for Project
type ProjectController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Project
// @Param	body		body 	requests.CreateProject	true		"body for Project content"
// @Success 201 {int} responses.ResponseProject
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *ProjectController) Post() {
	c.post(&requests.CreateProject{}, &models.Project{})
}

// GetOne ...
// @Title Get One
// @Description get Project by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.ResponseProject
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *ProjectController) GetOne() {
	c.getOne(&models.Project{})
}

// GetAll ...
// @Title Get All
// @Description get Project
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseProject
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *ProjectController) GetAll() {
	c.getAll([]models.Project{})
}

// Put ...
// @Title Put
// @Description update the Project
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateProject	true		"body for Project content"
// @Success 200 {object} responses.ResponseProject
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *ProjectController) Put() {
	c.put(&requests.UpdateProject{}, &models.Project{})
}

// Delete ...
// @Title Delete
// @Description delete the Project
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *ProjectController) Delete() {
	c.delete(&models.Project{})
}

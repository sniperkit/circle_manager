package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
)

//  TodoController operations for Todo
type TodoController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Todo
// @Param	body		body 	requests.CreateTodo	true		"body for Todo content"
// @Success 201 {int} responses.ResponseTodo
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *TodoController) Post() {
	c.post(&requests.CreateTodo{}, &models.Todo{})
}

// GetOne ...
// @Title Get One
// @Description get Todo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Todo
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *TodoController) GetOne() {
	c.getOne(&models.Todo{})
}

// GetAll ...
// @Title Get All
// @Description get Todo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseTodo
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *TodoController) GetAll() {
	c.getAll([]models.Todo{})
}

// Put ...
// @Title Put
// @Description update the Todo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateTodo	true		"body for Todo content"
// @Success 200 {object} responses.ResponseTodo
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *TodoController) Put() {
	c.put(&requests.UpdateTodo{}, &models.Todo{})
}

// Delete ...
// @Title Delete
// @Description delete the Todo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *TodoController) Delete() {
	c.delete(&models.Todo{})
}

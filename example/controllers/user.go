package controllers

import (
	"github.com/jungju/demo/models"
	"github.com/jungju/demo/requests"
)

//  UserController operations for User
type UserController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	requests.CreateUser	true		"body for User content"
// @Success 201 {int} responses.ResponseUser
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *UserController) Post() {
	c.post(&requests.CreateUser{}, &models.User{})
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.User
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *UserController) GetOne() {
	c.getOne(&models.User{})
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseUser
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *UserController) GetAll() {
	c.getAll([]models.User{})
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateUser	true		"body for User content"
// @Success 200 {object} responses.ResponseUser
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *UserController) Put() {
	c.put(&requests.UpdateUser{}, &models.User{})
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *UserController) Delete() {
	c.delete(&models.User{})
}

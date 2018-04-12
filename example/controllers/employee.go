package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
)

//  EmployeeController operations for Employee
type EmployeeController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Employee
// @Param	body		body 	models.Employee	true		"body for Employee content"
// @Success 201 {int} models.Employee
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *EmployeeController) Post() {
	c.post(&requests.CreateEmployee{}, &models.Employee{})
}

// GetOne ...
// @Title Get One
// @Description get Employee by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Employee
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *EmployeeController) GetOne() {
	c.getOne(&models.Employee{})
}

// GetAll ...
// @Title Get All
// @Description get Employee
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Employee
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *EmployeeController) GetAll() {
	c.getAll([]models.Employee{})
}

// Put ...
// @Title Put
// @Description update the Employee
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Employee	true		"body for Employee content"
// @Success 200 {object} models.Employee
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *EmployeeController) Put() {
	c.put(&requests.UpdateEmployee{}, &models.Employee{})
}

// Delete ...
// @Title Delete
// @Description delete the Employee
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *EmployeeController) Delete() {
	c.delete(&models.Employee{})
}

package controllers

import (
	"github.com/jungju/demo/models"
	"github.com/jungju/demo/requests"
)

//  CarController operations for Car
type CarController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Car
// @Param	body		body 	requests.CreateCar	true		"body for Car content"
// @Success 201 {int} responses.ResponseCar
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *CarController) Post() {
	c.post(&requests.CreateCar{}, &models.Car{})
}

// GetOne ...
// @Title Get One
// @Description get Car by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Car
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *CarController) GetOne() {
	c.getOne(&models.Car{})
}

// GetAll ...
// @Title Get All
// @Description get Car
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseCar
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *CarController) GetAll() {
	c.getAll([]models.Car{})
}

// Put ...
// @Title Put
// @Description update the Car
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateCar	true		"body for Car content"
// @Success 200 {object} responses.ResponseCar
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *CarController) Put() {
	c.put(&requests.UpdateCar{}, &models.Car{})
}

// Delete ...
// @Title Delete
// @Description delete the Car
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *CarController) Delete() {
	c.delete(&models.Car{})
}

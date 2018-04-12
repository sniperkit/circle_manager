package controllers

import (
	"github.com//models"
	"github.com//requests"
)

//  EventController operations for Event
type EventController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Event
// @Param	body		body 	requests.CreateEvent	true		"body for Event content"
// @Success 201 {int} responses.ResponseEvent
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *EventController) Post() {
	c.post(&requests.CreateEvent{}, &models.Event{})
}

// GetOne ...
// @Title Get One
// @Description get Event by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Event
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *EventController) GetOne() {
	c.getOne(&models.Event{})
}

// GetAll ...
// @Title Get All
// @Description get Event
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseEvent
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *EventController) GetAll() {
	c.getAll([]models.Event{})
}

// Put ...
// @Title Put
// @Description update the Event
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateEvent	true		"body for Event content"
// @Success 200 {object} responses.ResponseEvent
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *EventController) Put() {
	c.put(&requests.UpdateEvent{}, &models.Event{})
}

// Delete ...
// @Title Delete
// @Description delete the Event
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *EventController) Delete() {
	c.delete(&models.Event{})
}

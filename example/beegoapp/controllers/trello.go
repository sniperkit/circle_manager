package controllers

import (
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/requests"
	"github.com/jungju/circle_manager/example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

//  TrelloController operations for Trello
type TrelloController struct {
	modules.BaseUserController
}

func (c *TrelloController) Prepare() {
	c.RequestCreateItem = &requests.CreateTrello{}
	c.RequestUpdateItem = &requests.UpdateTrello{}
	c.ModelItem = &models.Trello{}
	c.ModelItems = &[]models.Trello{}
	c.ResponseItem = &responses.Trello{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create Trello
// @Param	body		body 	models.Trello	true		"body for Trello content"
// @Success 201 {object} responses.Trello
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *TrelloController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get Trello by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.Trello
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *TrelloController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get Trello
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Trello
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *TrelloController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the Trello
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Trello	true		"body for Trello content"
// @Success 200 {object} responses.Trello
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *TrelloController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the Trello
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *TrelloController) Delete() {
	c.BaseDelete()
}

package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
)

//  TeamController operations for Team
type TeamController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create Team
// @Param	body		body 	models.Team	true		"body for Team content"
// @Success 201 {int} models.Team
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *TeamController) Post() {
	c.post(&requests.CreateTeam{}, &models.Team{})
}

// GetOne ...
// @Title Get One
// @Description get Team by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Team
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *TeamController) GetOne() {
	c.getOne(&models.Team{})
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
	c.getAll([]models.Team{})
}

// Put ...
// @Title Put
// @Description update the Team
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Team	true		"body for Team content"
// @Success 200 {object} models.Team
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *TeamController) Put() {
	c.put(&requests.UpdateTeam{}, &models.Team{})
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
	c.delete(&models.Team{})
}

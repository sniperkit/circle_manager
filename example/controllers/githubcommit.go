package controllers

import (
	"github.com//models"
	"github.com//requests"
)

//  GithubCommitController operations for GithubCommit
type GithubCommitController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create GithubCommit
// @Param	body		body 	requests.CreateGithubCommit	true		"body for GithubCommit content"
// @Success 201 {int} responses.ResponseGithubCommit
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *GithubCommitController) Post() {
	c.post(&requests.CreateGithubCommit{}, &models.GithubCommit{})
}

// GetOne ...
// @Title Get One
// @Description get GithubCommit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.GithubCommit
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *GithubCommitController) GetOne() {
	c.getOne(&models.GithubCommit{})
}

// GetAll ...
// @Title Get All
// @Description get GithubCommit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []responses.ResponseGithubCommit
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *GithubCommitController) GetAll() {
	c.getAll([]models.GithubCommit{})
}

// Put ...
// @Title Put
// @Description update the GithubCommit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateGithubCommit	true		"body for GithubCommit content"
// @Success 200 {object} responses.ResponseGithubCommit
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *GithubCommitController) Put() {
	c.put(&requests.UpdateGithubCommit{}, &models.GithubCommit{})
}

// Delete ...
// @Title Delete
// @Description delete the GithubCommit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *GithubCommitController) Delete() {
	c.delete(&models.GithubCommit{})
}

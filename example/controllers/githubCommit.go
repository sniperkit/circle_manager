package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
	"github.com/jungju/circle/responses"
	"github.com/jungju/circle_manager/modules"
)

//  GithubCommitController operations for GithubCommit
type GithubCommitController struct {
	modules.BaseUserController
}

func (c *GithubCommitController) Prepare() {
	c.RequestCreateItem = &requests.CreateGithubCommit{}
	c.RequestUpdateItem = &requests.UpdateGithubCommit{}
	c.ModelItem = &models.GithubCommit{}
	c.ModelItems = &[]models.GithubCommit{}
	c.ResponseItem = &responses.GithubCommit{}
}

// Post ...
// @Title Post
// @Description create GithubCommit
// @Param	body		body 	models.GithubCommit	true		"body for GithubCommit content"
// @Success 201 {int} responses.GithubCommit
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *GithubCommitController) Post() {
	c.BasePost()
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
	c.BaseGetOne()
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
// @Success 200 {object} []responses.GithubCommit
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *GithubCommitController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the GithubCommit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GithubCommit	true		"body for GithubCommit content"
// @Success 200 {object} responses.GithubCommit
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *GithubCommitController) Put() {
	c.BasePut()
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
	c.BaseDelete()
}

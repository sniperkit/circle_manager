package controllers

import (
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/requests"
	"github.com/jungju/circle_manager/example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

//  GithubReleaseController operations for GithubRelease
type GithubReleaseController struct {
	modules.BaseUserController
}

func (c *GithubReleaseController) Prepare() {
	c.RequestCreateItem = &requests.CreateGithubRelease{}
	c.RequestUpdateItem = &requests.UpdateGithubRelease{}
	c.ModelItem = &models.GithubRelease{}
	c.ModelItems = &[]models.GithubRelease{}
	c.ResponseItem = &responses.GithubRelease{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create GithubRelease
// @Param	body		body 	models.GithubRelease	true		"body for GithubRelease content"
// @Success 201 {object} responses.GithubRelease
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *GithubReleaseController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get GithubRelease by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.GithubRelease
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *GithubReleaseController) GetOne() {
	c.BaseGetOne()
}

// GetAll ...
// @Title Get All
// @Description get GithubRelease
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.GithubRelease
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *GithubReleaseController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the GithubRelease
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GithubRelease	true		"body for GithubRelease content"
// @Success 200 {object} responses.GithubRelease
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *GithubReleaseController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the GithubRelease
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *GithubReleaseController) Delete() {
	c.BaseDelete()
}

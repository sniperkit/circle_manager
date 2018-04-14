package controllers

import (
	"github.com/jungju/circle/models"
	"github.com/jungju/circle/requests"
)

//  GithubReleaseController operations for GithubRelease
type GithubReleaseController struct {
	BaseCircleController
}

// Post ...
// @Title Post
// @Description create GithubRelease
// @Param	body		body 	requests.CreateGithubRelease	true		"body for GithubRelease content"
// @Success 201 {int} responses.ResponseGithubRelease
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *GithubReleaseController) Post() {
	c.post(&requests.CreateGithubRelease{}, &models.GithubRelease{})
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
	c.getOne(&models.GithubRelease{})
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
// @Success 200 {object} []responses.ResponseGithubRelease
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *GithubReleaseController) GetAll() {
	c.getAll([]models.GithubRelease{})
}

// Put ...
// @Title Put
// @Description update the GithubRelease
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateGithubRelease	true		"body for GithubRelease content"
// @Success 200 {object} responses.ResponseGithubRelease
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *GithubReleaseController) Put() {
	c.put(&requests.UpdateGithubRelease{}, &models.GithubRelease{})
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
	c.delete(&models.GithubRelease{})
}

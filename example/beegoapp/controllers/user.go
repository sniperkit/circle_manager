package controllers

import (
	"net/http"

	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/example/beegoapp/requests"
	"github.com/jungju/circle_manager/example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

// UserController operations for User
type UserController struct {
	modules.BaseUserController
}

func (c *UserController) Prepare() {
	c.RequestCreateItem = &requests.CreateUser{}
	c.RequestUpdateItem = &requests.UpdateUser{}
	c.ModelItem = &models.User{}
	c.ModelItems = &[]models.User{}
	c.ResponseItem = &responses.User{}

	c.BaseUserController.Prepare()
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {object} responses.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	c.BasePost()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	c.BaseGetOne()
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
// @Success 200 {object} responses.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	c.BaseGetAll()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateUser	true		"body for User content"
// @Success 200 {object} responses.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	c.BasePut()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	c.BaseDelete()
}

// Me ...
// @Title Me
// @Description 나의 사용자 정보
// @Success 200 {object} responses.User
// @Failure 403 body is empty
// @router /me [get]
func (c *UserController) Me() {
	user, err := models.GetUserByID(c.CurrentUserMeta.UserID)
	if err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusOK, user)
}

package modules

import (
	"net/http"

	"github.com/jinzhu/copier"
)

//  CircleSetController operations for CircleSet
type CircleSetController struct {
	BaseUserController
	RequestCreateItem *CreateCircleSet
	RequestUpdateItem *UpdateCircleSet
	ModelItem         *CircleSet
	ModelItems        *[]CircleSet
	ResponseItem      *CircleSet
}

func (c *CircleSetController) Prepare() {
	c.RequestCreateItem = &CreateCircleSet{}
	c.RequestUpdateItem = &UpdateCircleSet{}
	c.ModelItem = &CircleSet{}
	c.ModelItems = &[]CircleSet{}
	c.ResponseItem = &CircleSet{}
}

// Post ...
// @Title Post
// @Description create CircleSet
// @Param	body		body 	modules.CircleSet	true		"body for CircleSet content"
// @Success 201 {int} modules.CircleSet
// @Failure 403 body is empty
// @router / [post]
// @Security userAPIKey
func (c *CircleSetController) Post() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	c.SetRequestDataAndValid(c.RequestCreateItem)

	// 2. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	copier.Copy(c.ModelItem, c.RequestCreateItem)

	// 3. DB 입력 단계. Error이면 500
	err := CreateItem(c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계

	// 5. 사용자 응답 단계. 성공 응답 201
	c.Success(http.StatusCreated, c.ModelItem)
}

// GetOne ...
// @Title Get One
// @Description get modules.CircleSet by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} modules.CircleSet
// @Failure 403 :id is empty
// @router /:id [get]
// @Security userAPIKey
func (c *CircleSetController) GetOne() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	// 2. DB 요청 단계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계

	// 5. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ModelItem)
}

// GetAll ...
// @Title Get All
// @Description get CircleSet
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []modules.CircleSet
// @Failure 403
// @router / [get]
// @Security userAPIKey
func (c *CircleSetController) GetAll() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	reqPage := c.GetQueryPage()

	// 2. DB 요청 단계. Error이면 500
	err := GetItems(&c.ModelItems, reqPage)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계

	// 5. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ModelItems)
}

// Put ...
// @Title Put
// @Description update the CircleSet
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	modules.CircleSet	true		"body for CircleSet content"
// @Success 200 {object} modules.CircleSet
// @Failure 403 :id is not int
// @router /:id [put]
// @Security userAPIKey
func (c *CircleSetController) Put() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()
	c.SetRequestDataAndValid(c.RequestUpdateItem)

	// 1-1. 사용자 요청에 대한 DB 데이터 유효성 관계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 2. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	copier.Copy(c.ModelItem, c.RequestUpdateItem)

	// 3. DB 수정 단계. Error이면 500
	err = UpdateItem(c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 5. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ModelItem)
}

// Delete ...
// @Title Delete
// @Description delete the CircleSet
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 204
// @Failure 403 id is empty
// @router /:id [delete]
// @Security userAPIKey
func (c *CircleSetController) Delete() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	// 1-1. 사용자 요청에 대한 DB 데이터 유효성 관계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 2. DB 삭제 단계. Error이면 500
	err = DeleteItem(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 3. 사용자 응답 단계. 성공 응답 204
	c.Success(http.StatusNoContent, nil)
}

type CreateCircleSet struct {
	Name                  string
	Description           string
	Import                string
	IsEnable              bool
	AppVersion            string
	AppTitle              string
	AppDescription        string
	AppContact            string
	AppTermsOfServiceUrl  string
	AppLicense            string
	AppSecurityDefinition string
	RunAppEnvs            string
}

type UpdateCircleSet struct {
	Name                  string
	Description           string
	Import                string
	IsEnable              bool
	AppVersion            string
	AppTitle              string
	AppDescription        string
	AppContact            string
	AppTermsOfServiceUrl  string
	AppLicense            string
	AppSecurityDefinition string
	RunAppEnvs            string
}

func (c *CreateCircleSet) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateCircleSet) Valid() error {
	return validate.Struct(c)
}

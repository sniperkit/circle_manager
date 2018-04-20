package modules

import (
	"net/http"

	"github.com/jinzhu/copier"
)

type ModelItem interface {
}

type ModelItems interface {
}

type BaseCrudController struct {
	BaseController
	RequestCreateItem RequestBody
	RequestUpdateItem RequestBody
	ModelItem         ModelItem
	ModelItems        ModelItems
	ResponseItem      ResponseBody
}

func (c *BaseCrudController) BasePost() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	c.SetRequestDataAndValid(c.RequestCreateItem)

	// 2. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	copier.Copy(c.ModelItem, c.RequestCreateItem)

	// 3. DB 입력 단계. Error이면 500
	err := CreateItem(c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 5. 사용자 응답 단계. 성공 응답 201
	c.Success(http.StatusCreated, c.ResponseItem)
}

func (c *BaseCrudController) BaseGetOne() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	// 2. DB 요청 단계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 3. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 4. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BaseGetAll() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	reqPage := c.GetQueryPage()

	// 2. DB 요청 단계. Error이면 500
	err := GetItems(&c.ModelItems, reqPage)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 5. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BasePut() {
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

	// 4. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 5. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BaseDelete() {
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

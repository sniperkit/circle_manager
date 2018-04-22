package modules

import (
	"net/http"

	"github.com/fatih/structs"

	"github.com/jinzhu/copier"
)

type ModelItem interface {
	GetCreatorID() uint
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
	CurrentCircleUnit *CircleUnit
}

func (c *BaseCrudController) Prepare() {
	if CurCircleSet != nil {
		c.CurrentCircleUnit = CurCircleSet.GetUnit(structs.Name(c.ModelItem))
	}
}

func (c *BaseCrudController) BasePost() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	c.SetRequestDataAndValid(c.RequestCreateItem)

	// 2. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	copier.Copy(c.ModelItem, c.RequestCreateItem)

	if isEnable := checkAble("create", c.CurrentCircleUnit, c.CurrentUserMeta); !isEnable {
		c.ErrorAbort(403, nil)
	}

	// 3. DB 입력 단계. Error이면 500
	err := CreateItem(c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 5. 사용자 응답 단계. 성공 응답 201
	c.SuccessCreate(c.ModelItem, c.ResponseItem)
}

func (c *BaseCrudController) BaseGetOne() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	if isEnable := checkAble("getone", c.CurrentCircleUnit, c.CurrentUserMeta); !isEnable {
		c.ErrorAbort(403, nil)
	}

	// 2. DB 요청 단계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	if isUserData := checkUserData(c.CurrentCircleUnit, c.CurrentUserMeta, c.ModelItem); !isUserData {
		c.ErrorAbort(404, nil)
	}

	// 3. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 4. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BaseGetAll() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	reqPage := c.GetQueryPage()

	if isEnable := checkAble("list", c.CurrentCircleUnit, c.CurrentUserMeta); !isEnable {
		c.ErrorAbort(403, nil)
	}

	// 2. DB 요청 단계. Error이면 500
	err := getItems(c.CurrentCircleUnit, c.CurrentUserMeta, c.ModelItems, reqPage)
	c.CheckRecordNotFoundAndServerError(err)

	// 3. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 4. 사용자 응답 단계. 성공 응답 200
	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BasePut() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()
	c.SetRequestDataAndValid(c.RequestUpdateItem)

	if isEnable := checkAble("update", c.CurrentCircleUnit, c.CurrentUserMeta); !isEnable {
		c.ErrorAbort(403, nil)
	}

	// 1-1. 사용자 요청에 대한 DB 데이터 유효성 관계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	if isUserData := checkUserData(c.CurrentCircleUnit, c.CurrentUserMeta, c.ModelItem); !isUserData {
		c.ErrorAbort(404, nil)
	}

	// 2. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	copier.Copy(c.ModelItem, c.RequestUpdateItem)

	// 3. DB 수정 단계. Error이면 500
	err = UpdateItem(c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 4. 사용자 응답 데이터 가공 단계
	copier.Copy(c.ResponseItem, c.ModelItem)

	// 5. 사용자 응답 단계. 성공 응답 200
	c.SuccessUpdate(c.ModelItem, c.ResponseItem)
}

func (c *BaseCrudController) BaseDelete() {
	// 1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	if isEnable := checkAble("delete", c.CurrentCircleUnit, c.CurrentUserMeta); !isEnable {
		c.ErrorAbort(403, nil)
	}

	// 1-1. 사용자 요청에 대한 DB 데이터 유효성 관계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	if isUserData := checkUserData(c.CurrentCircleUnit, c.CurrentUserMeta, c.ModelItem); !isUserData {
		c.ErrorAbort(404, nil)
	}

	// 2. DB 삭제 단계. Error이면 500
	err = DeleteItem(id, c.ModelItem)
	c.CheckRecordNotFoundAndServerError(err)

	// 3. 사용자 응답 단계. 성공 응답 204
	c.SuccessDelete(c.ModelItem)
}

func getItems(circleUnit *CircleUnit, userMeta *UserMeta, modelItems ModelItems, reqPage *QueryPage) error {
	if circleUnit != nil {
		if circleUnit.OlnyUserData {
			//TODO: 제외되는 UserID 체크(Admin 등)
			//TODO: userMeta가 없을 떄 처리
			return GetItemsOnlyUserData(modelItems, reqPage, userMeta.UserID)
		}
	}
	return GetItems(modelItems, reqPage)
}

func checkUserData(circleUnit *CircleUnit, userMeta *UserMeta, getedModel ModelItem) bool {
	if circleUnit != nil {
		if circleUnit.OlnyUserData {
			if value := getedModel.GetCreatorID(); value == userMeta.UserID {
				return true
			}
			return true
		}
	}
	return true
}

func checkAble(checkType string, circleUnit *CircleUnit, userMeta *UserMeta) bool {
	if circleUnit != nil {
		checkAbleProp := func(isEnable bool, tags string) bool {
			if !isEnable {
				return false
			} else if tags != "" {
				if userMeta == nil {
					return false
				} else {
					//TODO: 태그 처리
				}
			}
			return true
		}
		switch checkType {
		case "create":
			checkAbleProp(circleUnit.IsCreateble, circleUnit.CreatebleTags)
		case "update":
			checkAbleProp(circleUnit.IsUpdateble, circleUnit.UpdatebleTags)
		case "list":
			checkAbleProp(circleUnit.IsListable, circleUnit.ListableTags)
		case "getone":
			checkAbleProp(circleUnit.IsGetable, circleUnit.GetableTags)
		case "delete":
			checkAbleProp(circleUnit.IsDeleteble, circleUnit.DeletebleTags)
		}
	}
	return true
}

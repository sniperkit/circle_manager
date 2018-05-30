package modules

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/fatih/structs"

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
	OldModelItem      ModelItem
	ModelItems        ModelItems
	ResponseItem      ResponseBody
	ResponseItems     ResponseBodies
	CurrentCircleUnit *CircleUnit
}

func (c *BaseCrudController) Prepare() {
	if CurCircleSet != nil {
		c.CurrentCircleUnit = CurCircleSet.GetUnit(structs.Name(c.ModelItem))
	}
}

func (c *BaseCrudController) BasePost() {
	// @step1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	c.SetRequestDataAndValid(c.RequestCreateItem)

	// @step2. API 권한 체크
	c.CheckAble("create")

	// @step3. 사용자 요청 데이터에서 모델 데이터로 가공
	err := copier.Copy(c.ModelItem, c.RequestCreateItem)
	c.Check404And500(err)

	// @step4. DB 입력 단계. Error이면 500.
	err = CreateItem(c.ModelItem)
	c.Check404And500(err)

	// @step5. 사용자 응답 데이터 가공 및 응답
	err = copier.Copy(c.ResponseItem, c.ModelItem)
	c.Check404And500(err)

	c.SuccessCreate()
}

func (c *BaseCrudController) BaseGetOne() {
	// @step1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	// @step2. API 권한 체크
	c.CheckAble("getone")

	// @step3. DB 요청 단계. Error이면 404, 500
	err := GetItemByID(id, c.ModelItem)
	c.Check404And500(err)

	// @step4. 접근 데이터 체크. 접근 할수 없는 데이터는 404
	c.CheckUserData(404)

	// @step5. 사용자 응답 데이터 가공 및 응답
	err = copier.Copy(c.ResponseItem, c.ModelItem)
	c.Check404And500(err)

	c.Success(http.StatusOK, c.ResponseItem)
}

func (c *BaseCrudController) BaseGetAll() {
	// @step1. API 권한 체크
	c.CheckAble("list")

	// @step2. DB 요청 단계. Error이면 500
	err := c.GetItems()
	c.Check404And500(err)

	// @step3. 사용자 응답 데이터 가공 및 응답
	err = copier.Copy(c.ResponseItems, c.ModelItems)
	c.Check404And500(err)

	c.Success(http.StatusOK, c.ResponseItems)
}

func (c *BaseCrudController) BasePut() {
	// @step1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()
	c.SetRequestDataAndValid(c.RequestUpdateItem)

	// @step2. API 권한 체크
	c.CheckAble("update")

	// @step3. 사용자 요청에 대한 DB 데이터 유효성 확인. Error이면 404 or 500
	err := GetItemByID(id, c.ModelItem)
	c.Check404And500(err)

	c.OldModelItem = reflect.New(reflect.TypeOf(c.ModelItem)).Elem().Interface().(ModelItem)
	err = copier.Copy(c.OldModelItem, c.ModelItem)
	c.Check404And500(err)

	// @step4. 접근 데이터 체크. 접근 할수 없는 데이터는 404
	c.CheckUserData(404)

	// @step5. 사용자 요청 데이터에서 DB 데이터로 가공 단계
	err = copier.Copy(c.ModelItem, c.RequestUpdateItem)
	c.Check404And500(err)

	// @step6. DB 수정 단계. Error이면 500
	err = UpdateItem(id, c.ModelItem)
	c.Check404And500(err)

	// @step7. 사용자 응답 데이터 가공 및 응답
	err = copier.Copy(c.ResponseItem, c.ModelItem)
	c.Check404And500(err)

	c.SuccessUpdate()
}

func (c *BaseCrudController) BaseDelete() {
	// @step1. 사용자 요청에 대한 유효성 처리 단계. Error이면 400
	id := c.GetParamID()

	// @step2. API 권한 체크
	c.CheckAble("delete")

	// @step3. 사용자 요청에 대한 DB 데이터 유효성 관계. Error이면 404 or 500
	err := GetItemByID(id, c.ModelItem)
	c.Check404And500(err)

	// @step4. 접근 데이터 체크. 접근 할수 없는 데이터는 404
	c.CheckUserData(404)

	// @step5. DB 삭제 단계. Error이면 500
	err = DeleteItem(id, c.ModelItem)
	c.Check404And500(err)

	// @step6. 사용자 응답
	c.SuccessDelete()
}

func (c *BaseCrudController) GetItems() error {
	reqPage := c.GetQueryPage()

	// TODO: 나중에
	// if c.CurrentCircleUnit != nil {
	// 	if c.CurrentCircleUnit.OlnyUserData {
	// 		//TODO: 제외되는 UserID 체크(Admin 등)
	// 		//TODO: userMeta가 없을 떄 처리
	// 		return GetItemsOnlyUserData(c.ModelItems, reqPage, c.CurrentUserMeta.UserID)
	// 	}
	// }
	return GetItems(c.ModelItems, reqPage)
}

func (c *BaseCrudController) CheckUserData(thenErrorCode int) {
	if c.CurrentCircleUnit == nil {
		return
	}
	if !c.CurrentCircleUnit.OlnyUserData {
		return
	}

	value := extractValueByField("CreatorID", c.ModelItem)
	if userID, ok := value.(uint); ok {
		if userID != c.CurrentUserMeta.UserID {
			c.ErrorAbort(thenErrorCode, nil)
		}
	}
}

func (c *BaseCrudController) CheckAble(checkType string) {
	// TODO: 지금 사용 안함
	// if c.CurrentCircleUnit != nil {
	// 	checkAbleProp := func(userIDs, userTypeIDs, userStatusIDs string) bool {
	// 		if c.CurrentUserMeta == nil && userIDs != "" && userTypeIDs != "" && userStatusIDs != "" {
	// 			return false
	// 		}

	// 		if userIDs != "" {
	// 			//TODO:
	// 		} else if userTypeIDs != "" {
	// 			//TODO:
	// 		} else if userStatusIDs != "" {
	// 			//TODO:
	// 		}
	// 		return false
	// 	}
	// 	switch checkType {
	// 	case "create":
	// 		if !c.CurrentCircleUnit.IsCreateble {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.CreatebleUserExcludeIDs,
	// 			c.CurrentCircleUnit.CreatebleUserExcludeTypeIDs,
	// 			c.CurrentCircleUnit.CreatebleUserExcludeStatusIDs) {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.CreatebleUserIDs,
	// 			c.CurrentCircleUnit.CreatebleUserTypeIDs,
	// 			c.CurrentCircleUnit.CreatebleUserStatusIDs) {
	// 			return
	// 		}
	// 	case "update":
	// 		if !c.CurrentCircleUnit.IsUpdateble {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.UpdatableUserExcludeIDs,
	// 			c.CurrentCircleUnit.UpdatableUserExcludeTypeIDs,
	// 			c.CurrentCircleUnit.UpdatableUserExcludeStatusIDs) {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.UpdatableUserIDs,
	// 			c.CurrentCircleUnit.UpdatableUserTypeIDs,
	// 			c.CurrentCircleUnit.UpdatableUserStatusIDs) {
	// 			return
	// 		}
	// 	case "list":
	// 		if !c.CurrentCircleUnit.IsGetAllable {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.GetAllableUserExcludeIDs,
	// 			c.CurrentCircleUnit.GetAllableUserExcludeTypeIDs,
	// 			c.CurrentCircleUnit.GetAllableUserExcludeStatusIDs) {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.GetAllableUserIDs,
	// 			c.CurrentCircleUnit.GetAllableUserTypeIDs,
	// 			c.CurrentCircleUnit.GetAllableUserStatusIDs) {
	// 			return
	// 		}
	// 	case "getone":
	// 		if !c.CurrentCircleUnit.IsGetOneable {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.GetOneableUserExcludeIDs,
	// 			c.CurrentCircleUnit.GetOneableUserExcludeTypeIDs,
	// 			c.CurrentCircleUnit.GetOneableUserExcludeStatusIDs) {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.GetOneableUserIDs,
	// 			c.CurrentCircleUnit.GetOneableUserTypeIDs,
	// 			c.CurrentCircleUnit.GetOneableUserStatusIDs) {
	// 			return
	// 		}
	// 	case "delete":
	// 		if !c.CurrentCircleUnit.IsDeleteble {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.DeletableUserExcludeIDs,
	// 			c.CurrentCircleUnit.DeletableUserExcludeTypeIDs,
	// 			c.CurrentCircleUnit.DeletableUserExcludeStatusIDs) {
	// 			c.ErrorAbort(404, nil)
	// 		} else if checkAbleProp(c.CurrentCircleUnit.DeletableUserIDs,
	// 			c.CurrentCircleUnit.DeletableUserTypeIDs,
	// 			c.CurrentCircleUnit.DeletableUserStatusIDs) {
	// 			return
	// 		}
	// 	}
	// }
}

// SuccessCreate ...
func (c *BaseCrudController) SuccessCreate() {
	go c.addEvent("create")
	c.Success(http.StatusCreated, c.ResponseItem)
}

// SuccessUpdate ...
func (c *BaseCrudController) SuccessUpdate() {
	go c.addEvent("update")
	c.Success(http.StatusOK, c.ResponseItem)
}

// SuccessDelete ...
func (c *BaseCrudController) SuccessDelete() {
	go c.addEvent("delete")
	c.Success(http.StatusNoContent, nil)
}

func (c *BaseCrudController) addEvent(actionName string) {
	userID := uint(1)
	if c.CurrentUserMeta != nil {
		userID = c.CurrentUserMeta.UserID
	}
	requestData := ""
	UpdatedData := ""
	oldData := ""
	switch actionName {
	case "create":
		requestData = ConvJsonData(c.RequestCreateItem)
		UpdatedData = ConvJsonData(c.ModelItem)
	case "update":
		requestData = ConvJsonData(c.RequestUpdateItem)
		UpdatedData = ConvJsonData(c.ModelItem)
		oldData = ConvJsonData(c.OldModelItem)
	case "delete":
		oldData = ConvJsonData(c.ModelItem)
	}

	resourceID := uint(0)
	if field, ok := structs.New(c.ModelItem).FieldOk("ID"); ok {
		resourceID = field.Value().(uint)
	}

	if _, err := AddCrudEvent(&CrudEvent{
		ActionName:   actionName,
		ActionType:   actionName,
		ResourceID:   resourceID,
		ResourceName: structs.Name(c.ModelItem),
		CreatorID:    userID,
		Where:        "API",
		UpdatedData:  UpdatedData,
		RequestData:  requestData,
		OldData:      oldData,
		ResponseData: ConvJsonData(c.ResponseItem),
	}); err != nil {
		fmt.Println(err)
	}
}

// func EventThenCreate(modelItem ModelItem, currentUserID *uint) {
// 	evnet(structs.Name(modelItem), "add", currentUserID, nil, modelItem)
// }

// func EventThenUpdate(modelItem ModelItem, oldModelItem ModelItem, currentUserID *uint) {
// 	mapUpdateProperties := makeMapUpdateProperties(modelItem, oldModelItem)
// 	evnet(structs.Name(modelItem), "update", currentUserID, mapUpdateProperties, modelItem)
// }

// func makeMapUpdateProperties(modelItem ModelItem, oldModelItem ModelItem) map[string]UpdateProperty {
// 	mapUpdateProperties := map[string]UpdateProperty{}

// 	if modelItem == nil || oldModelItem == nil {
// 		return mapUpdateProperties
// 	}
// 	mapModelItem := structs.Map(modelItem)
// 	mapOldModelItem := structs.Map(oldModelItem)

// 	for key, value := range mapModelItem {
// 		if structs.IsStruct(value) {
// 			continue
// 		}

// 		oldValue := ""
// 		if tempOldValue, ok := mapOldModelItem[key]; ok {
// 			oldValue = convInterface(tempOldValue)
// 		}

// 		mapUpdateProperties[key] = UpdateProperty{
// 			Key:      key,
// 			NewValue: convInterface(value),
// 			OldValue: oldValue,
// 		}
// 	}
// 	return mapUpdateProperties
// }

// func EventThenDelete(modelItem ModelItem, currentUserID *uint) {
// 	evnet(structs.Name(modelItem), "delete", currentUserID, nil, modelItem)
// }

// func evnet(structName string, action string, eventUserID *uint, mapUpdateProperties map[string]UpdateProperty, datas ...interface{}) {
// 	if err := AddActionNotification(
// 		fmt.Sprintf("%s,%s", structName, action),
// 		eventUserID,
// 		mapUpdateProperties,
// 		datas...,
// 	); err != nil {
// 		fmt.Printf("Error : %s\n", err.Error())
// 	}
// }

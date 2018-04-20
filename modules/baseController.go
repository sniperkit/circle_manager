package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fatih/structs"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//BaseController ...
type BaseController struct {
	beego.Controller
	CurrentUserMeta *UserMeta
}

// QueryPage ...
type QueryPage struct {
	Fields     []string
	Sortby     []string
	Order      []string
	Query      map[string]string
	Limit      int
	Offset     int
	Tags       []uint
	Properties map[string][]string
}

// Success ...
func (c *BaseController) Success(statusCode int, data interface{}) {
	ext := c.Ctx.Input.Param(":ext")
	c.Ctx.Output.SetStatus(statusCode)
	if data != nil {
		if ext == "xml" {
			c.Data["xml"] = data
			c.ServeXML()
		} else {
			c.Data["json"] = data
			c.ServeJSON()
		}
	}
}

func getUserIDByUserMeta(userMeta *UserMeta) *uint {
	if userMeta == nil {
		return nil
	}
	return &userMeta.UserID
}

// SuccessCreate ...
func (c *BaseController) SuccessCreate(data interface{}) {
	evnet(structs.Name(data), "add", getUserIDByUserMeta(c.CurrentUserMeta), data)
	c.Success(http.StatusCreated, data)
}

// SuccessUpdate ...
func (c *BaseController) SuccessUpdate(data interface{}) {
	evnet(structs.Name(data), "update", getUserIDByUserMeta(c.CurrentUserMeta), data)
	c.Success(http.StatusOK, data)
}

// SuccessDelete ...
func (c *BaseController) SuccessDelete(data interface{}) {
	evnet(structs.Name(data), "delete", getUserIDByUserMeta(c.CurrentUserMeta), data)
	c.Success(http.StatusNoContent, nil)
}

func evnet(structName string, action string, eventUserID *uint, data interface{}) {
	if err := AddActionNotification(
		fmt.Sprintf("%s,%s", structName, action),
		eventUserID,
		data,
	); err != nil {
		logrus.WithError(err).Error()
	}
}

// ErrorAbort ...
func (c *BaseController) ErrorAbort(code int, err error, withMsg ...interface{}) {
	if err != nil {
		logrus.WithError(err).Error()
		c.CustomAbort(code, err.Error())
	}
	c.CustomAbort(code, "")
}

func (c *BaseController) GetQueryPage() *QueryPage {
	queryPage := &QueryPage{
		Query: make(map[string]string),
		Limit: 10,
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		queryPage.Fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetUint64("limit"); err == nil {
		queryPage.Limit = int(v)
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		queryPage.Offset = int(v)
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		queryPage.Sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		queryPage.Order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.ErrorAbort(400, ErrInvalidQuery)
			}
			k, v := kv[0], kv[1]
			queryPage.Query[k] = v
		}
	}
	return queryPage
}

func (c *BaseController) setRequestDataInterface(reqBody interface{}) error {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		return ErrInvalidRequestBody
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		return err
	}
	return nil
}

func (c *BaseController) setRequestData(reqBody RequestBody) error {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		return ErrInvalidRequestBody
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		return err
	}
	return nil
}

func (c *BaseController) SetRequestDataAndValid(reqBody RequestBody) {
	if len(c.Ctx.Input.RequestBody) == 0 {
		logrus.Error("Body에 정보가 없습니다.")
		c.ErrorAbort(400, ErrInvalidRequestBody)
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
		c.ErrorAbort(400, err)
	}

	// if err := c.setRequestData(reqBody); err != nil {
	// 	c.ErrorAbort(400, err)
	// }
	if err := reqBody.Valid(); err != nil {
		c.ErrorAbort(400, err)
	}
}

func (c *BaseController) GetParamForUintTypeRequired(key string, retID *uint) error {
	idString := c.Ctx.Input.Param(key)
	if idString == "" {
		return ErrInvalidRequestParam
	}
	id, err := strconv.ParseUint(idString, 10, 64)
	if err == nil {
		if id == 0 {
			return ErrInvalidRequestParam
		}
		uintID := uint(id)
		*retID = uintID
		return nil
	}
	return ErrInvalidRequestParam
}

func (c *BaseController) GetParamID() uint {
	id := uint(0)
	if err := c.GetParamForUintTypeRequired(":id", &id); err != nil {
		c.ErrorAbort(400, err)
	}
	return id
}

func (c *BaseController) CheckRecordNotFoundAndServerError(err error) {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		} else {
			c.ErrorAbort(500, err)
		}
	}
}

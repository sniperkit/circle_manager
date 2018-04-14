package modules

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

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

//BaseController ...
type BaseCircleController struct {
	BaseController
}

func (c *BaseCircleController) BasePost(reqBody RequestBody, object interface{}) {
	c.SetRequestDataAndValid(reqBody)

	copier.Copy(object, reqBody)

	if err := CreateItem(object); err != nil {
		c.ErrorAbort(500, err)
	}
	c.Success(http.StatusCreated, object)
}

func (c *BaseCircleController) BaseGetOne(object interface{}) {
	id := c.GetParamID()

	if err := GetItemByID(id, object); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		}
		c.ErrorAbort(500, err)
	}
	c.Success(http.StatusOK, object)
}

func (c *BaseCircleController) BaseGetAll(objects interface{}) {
	reqPage, err := c.GetQueryPage()
	if err != nil {
		c.ErrorAbort(400, err)
	}

	if err := GetItems(&objects, reqPage); err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusOK, objects)
}

func (c *BaseCircleController) BasePut(reqBody RequestBody, object interface{}) {
	id := c.GetParamID()
	c.SetRequestDataAndValid(reqBody)

	if err := GetItemByID(id, object); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		}
		c.ErrorAbort(500, err)
	}

	copier.Copy(object, reqBody)

	if err := UpdateItem(object); err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusOK, object)
}

func (c *BaseCircleController) BaseDelete(object interface{}) {
	id := c.GetParamID()

	if err := GetItemByID(id, object); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		}
		c.ErrorAbort(500, err)
	}

	if err := DeleteItem(id, object); err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusNoContent, nil)
}

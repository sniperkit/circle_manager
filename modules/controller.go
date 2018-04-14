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

func (c *BaseCircleController) basePost(reqBody RequestBody, object interface{}) {
	c.setRequestDataAndValid(reqBody)

	copier.Copy(object, reqBody)

	if err := CreateItem(object); err != nil {
		c.ErrorAbort(500, err)
	}
	c.Success(http.StatusCreated, object)
}

func (c *BaseCircleController) baseGetOne(object interface{}) {
	id := c.GetParamID()

	if err := GetItemByID(id, object); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ErrorAbort(404, err)
		}
		c.ErrorAbort(500, err)
	}
	c.Success(http.StatusOK, object)
}

func (c *BaseCircleController) baseGetAll(objects interface{}) {
	reqPage, err := c.getQueryPage()
	if err != nil {
		c.ErrorAbort(400, err)
	}

	if err := GetItems(&objects, reqPage); err != nil {
		c.ErrorAbort(500, err)
	}

	c.Success(http.StatusOK, objects)
}

func (c *BaseCircleController) basePut(reqBody RequestBody, object interface{}) {
	id := c.GetParamID()
	c.setRequestDataAndValid(reqBody)

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

func (c *BaseCircleController) baseDelete(object interface{}) {
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

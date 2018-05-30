package main

const (
	adminTemplate = `addResourceAndMenu(&models.{{.Name}}{}, "{{.MenuName}}", "{{.MenuGroup}}", anyoneAllow, -1)
	`
	routerTemplate = `beego.NSNamespace("/{{.GetURL}}",
			beego.NSInclude(
				&controllers.{{.Name}}Controller{},
			),
		),`
	MODEL_TEMPLATE = `package models

	import (
		"time"
	
		"github.com/jungju/circle_manager/modules"
	)
	
	// gen:qs
	type {{.Name}} struct {
		ID          uint      ` + "`description:\"\"`" + `
		CreatedAt   time.Time ` + "`description:\"등록일\"`" + `
		UpdatedAt   time.Time ` + "`description:\"수정일\"`" + `
		Name        string    ` + "`description:\"이름\"`" + `
		Description string    ` + "`description:\"설명\" sql:\"type:text\"`" + `
		CreatorID   uint      ` + "`description:\"작성자\"`" + `
		{{range $i, $property := .Properties}}{{.Name}} {{.GetTypeInModel}} ` + "`description:\"{{.Description}}\"`" + `
		{{end}}
	}
	
	func init() {
		registModel(&{{.Name}}{})
	}
	
	func Add{{.Name}}({{.GetVariableName}} *{{.Name}}) (id uint, err error) {
		err = {{.GetVariableName}}.Create(gGormDB)
		id = {{.GetVariableName}}.ID
		return
	}
	
	func Get{{.Name}}ByID(id uint) ({{.GetVariableName}} *{{.Name}}, err error) {
		{{.GetVariableName}} = &{{.Name}}{
			ID: id,
		}
		err = New{{.Name}}QuerySet(gGormDB).
			One({{.GetVariableName}})
		return
	}
	
	func GetAll{{.Name}}(queryPage *modules.QueryPage) ({{.GetVariableName}}s []{{.Name}}, err error) {
		err = New{{.Name}}QuerySet(gGormDB).
			All(&{{.GetVariableName}}s)
		return
	}
	
	func Update{{.Name}}ByID({{.GetVariableName}} *{{.Name}}) (err error) {
		err = {{.GetVariableName}}.Update(gGormDB,
			{{.Name}}DBSchema.Name,
			{{.Name}}DBSchema.Description,
			{{$name := .Name}}{{range $i, $property := .Properties}}{{$name}}DBSchema.{{.Name}},
			{{end}}
		)
		return
	}
	
	func Delete{{.Name}}(id uint) (err error) {
		{{.GetVariableName}} := &{{.Name}}{
			ID: id,
		}
		err = {{.GetVariableName}}.Delete(gGormDB)
		return
	}
	`
	CONTROLLER_TEMPLATE = `package controllers

	import (
		"github.com/{{.Import}}/models"
		"github.com/{{.Import}}/requests"
		"github.com/{{.Import}}/responses"
		"github.com/jungju/circle_manager/modules"
	)
	
	//  {{.Name}}Controller operations for {{.Name}}
	type {{.Name}}Controller struct {
		modules.BaseUserController
	}
	
	func (c *{{.Name}}Controller) Prepare() {
		c.RequestCreateItem = &requests.Create{{.Name}}{}
		c.RequestUpdateItem = &requests.Update{{.Name}}{}
		c.ModelItem = &models.{{.Name}}{}
		c.ModelItems = &[]models.{{.Name}}{}
		c.ResponseItem = &responses.{{.Name}}{}
	}
	
	// Post ...
	// @Title Post
	// @Description create {{.Name}}
	// @Param	body		body 	models.{{.Name}}	true		"body for {{.Name}} content"
	// @Success 201 {int} responses.{{.Name}}
	// @Failure 403 body is empty
	// @router / [post]
	// @Security userAPIKey
	func (c *{{.Name}}Controller) Post() {
		c.BasePost()
	}
	
	// GetOne ...
	// @Title Get One
	// @Description get {{.Name}} by id
	// @Param	id		path 	string	true		"The key for staticblock"
	// @Success 200 {object} responses.{{.Name}}
	// @Failure 403 :id is empty
	// @router /:id [get]
	// @Security userAPIKey
	func (c *{{.Name}}Controller) GetOne() {
		c.BaseGetOne()
	}
	
	// GetAll ...
	// @Title Get All
	// @Description get {{.Name}}
	// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
	// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
	// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
	// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
	// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
	// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
	// @Success 200 {object} []responses.{{.Name}}
	// @Failure 403
	// @router / [get]
	// @Security userAPIKey
	func (c *{{.Name}}Controller) GetAll() {
		c.BaseGetAll()
	}
	
	// Put ...
	// @Title Put
	// @Description update the {{.Name}}
	// @Param	id		path 	string	true		"The id you want to update"
	// @Param	body		body 	models.{{.Name}}	true		"body for {{.Name}} content"
	// @Success 200 {object} responses.{{.Name}}
	// @Failure 403 :id is not int
	// @router /:id [put]
	// @Security userAPIKey
	func (c *{{.Name}}Controller) Put() {
		c.BasePut()
	}
	
	// Delete ...
	// @Title Delete
	// @Description delete the {{.Name}}
	// @Param	id		path 	string	true		"The id you want to delete"
	// @Success 204
	// @Failure 403 id is empty
	// @router /:id [delete]
	// @Security userAPIKey
	func (c *{{.Name}}Controller) Delete() {
		c.BaseDelete()
	}
	`
	REQUEST_TEMPLATE = `package requests

	import "time"
	
	type Create{{.Name}} struct {
	  {{range $i, $property := .Properties}}{{.Name}} {{.Type}}
		{{end}}
	}
	
	type Update{{.Name}} struct {
	  {{range $i, $property := .Properties}}{{.Name}} {{.Type}}
		{{end}}
	}
	
	func (c *Create{{.Name}}) Valid() error {
	  return validate.Struct(c)
	}
	
	func (c *Update{{.Name}}) Valid() error {
	  return validate.Struct(c)
	}
	`
	RESPONSE_TEMPLATE = `package responses

	import "time"
	
	type {{.Name}} struct {
	  ID          uint      
		CreatedAt   time.Time 
		UpdatedAt   time.Time 
		Name        string    
		Description string    
	  {{range $i, $property := .Properties}}{{.Name}} {{.Type}}
		{{end}}
	}
	
	`
)

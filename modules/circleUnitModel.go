package modules

import (
	"fmt"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/jinzhu/inflection"
)

// gen:qs
type CircleUnit struct {
	ID                             uint                 `description:""`
	CreatedAt                      time.Time            `description:"등록일"`
	UpdatedAt                      time.Time            `description:"수정일"`
	Name                           string               `description:"이름"`
	Description                    string               `description:"설명" sql:"type:text"`
	CreatorID                      uint                 `description:"작성자"`
	CircleSet                      CircleSet            `description:""`
	CircleSetID                    uint                 `description:""`
	Properties                     []CircleUnitProperty `description:""`
	Import                         string               `description:""`
	URL                            string               `description:""` //DEPRECATE
	MenuName                       string               `description:""`
	MenuGroup                      string               `description:""`
	IsEnable                       bool                 `description:"사용여부"`
	IsManual                       bool                 `description:""`
	IsSystem                       bool                 `description:""`
	IsCreateble                    bool                 `description:""`
	CreatebleUserIDs               string               `description:""`
	CreatebleUserTypeIDs           string               `description:""`
	CreatebleUserStatusIDs         string               `description:""`
	CreatebleUserExcludeIDs        string               `description:""`
	CreatebleUserExcludeTypeIDs    string               `description:""`
	CreatebleUserExcludeStatusIDs  string               `description:""`
	IsGetAllable                   bool                 `description:""`
	GetAllableUserIDs              string               `description:""`
	GetAllableUserTypeIDs          string               `description:""`
	GetAllableUserStatusIDs        string               `description:""`
	GetAllableUserExcludeIDs       string               `description:""`
	GetAllableUserExcludeTypeIDs   string               `description:""`
	GetAllableUserExcludeStatusIDs string               `description:""`
	IsGetOneable                   bool                 `description:""`
	GetOneableUserIDs              string               `description:""`
	GetOneableUserTypeIDs          string               `description:""`
	GetOneableUserStatusIDs        string               `description:""`
	GetOneableUserExcludeIDs       string               `description:""`
	GetOneableUserExcludeTypeIDs   string               `description:""`
	GetOneableUserExcludeStatusIDs string               `description:""`
	IsDeleteble                    bool                 `description:""`
	DeletableUserIDs               string               `description:""`
	DeletableUserTypeIDs           string               `description:""`
	DeletableUserStatusIDs         string               `description:""`
	DeletableUserExcludeIDs        string               `description:""`
	DeletableUserExcludeTypeIDs    string               `description:""`
	DeletableUserExcludeStatusIDs  string               `description:""`
	IsUpdateble                    bool                 `description:""`
	UpdatableUserIDs               string               `description:""`
	UpdatableUserTypeIDs           string               `description:""`
	UpdatableUserStatusIDs         string               `description:""`
	UpdatableUserExcludeIDs        string               `description:""`
	UpdatableUserExcludeTypeIDs    string               `description:""`
	UpdatableUserExcludeStatusIDs  string               `description:""`
	OlnyUserData                   bool                 `description:""`
	//AllDataTags   string               `description:""`

}

func (c *CircleUnit) GetCreatorID() uint {
	return c.CreatorID
}

func (m *CircleUnit) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func (c CircleUnit) GetVariableName() string {
	return lowerFirst(c.Name)
}

func (c CircleUnit) GetURL() string {
	return MakeFirstLowerCase(inflection.Plural(c.Name))
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func (c CircleUnit) GetControllerName() string {
	return fmt.Sprintf("%sController", c.Name)
}

func AddCircleUnit(circleUnit *CircleUnit) (id uint, err error) {
	err = circleUnit.Create(gGormDB)
	id = circleUnit.ID
	return
}

func GetCircleUnitByID(id uint) (circleUnit *CircleUnit, err error) {
	circleUnit = &CircleUnit{
		ID: id,
	}
	err = NewCircleUnitQuerySet(gGormDB).One(circleUnit)
	return
}

func GetAllCircleUnit(queryPage *QueryPage) (circleUnits []CircleUnit, err error) {
	err = NewCircleUnitQuerySet(gGormDB).All(&circleUnits)
	return
}

func UpdateCircleUnitByID(circleUnit *CircleUnit) (err error) {
	err = circleUnit.Update(gGormDB,
		CircleUnitDBSchema.Description,
	)
	return
}

func DeleteCircleUnit(id uint) (err error) {
	circleUnit := &CircleUnit{
		ID: id,
	}
	err = circleUnit.Delete(gGormDB)
	return
}

func SaveCircleUnit(nameAndValue map[string]interface{}) error {
	//TODO: ....
	return nil
}

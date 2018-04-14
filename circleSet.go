package main

import (
	"time"

	stringcase "github.com/reiver/go-stringcase"
)

type CircleSet struct {
	ID                    uint         `description:""`
	CreatedAt             time.Time    `description:"등록일"`
	UpdatedAt             time.Time    `description:"수정일"`
	Name                  string       `description:"이름"`
	Description           string       `description:"설명" sql:"type:text"`
	Import                string       `description:""`
	Units                 []CircleUnit `description:""`
	IsEnable              bool         `description:"사용여부"`
	AppVersion            string       `description:""`
	AppTitle              string       `description:""`
	AppDescription        string       `description:""`
	AppContact            string       `description:""`
	AppTermsOfServiceUrl  string       `description:""`
	AppLicense            string       `description:""`
	AppSecurityDefinition string       `description:""`
}

type CircleUnit struct {
	ID             uint                 `description:""`
	CreatedAt      time.Time            `description:"등록일"`
	UpdatedAt      time.Time            `description:"수정일"`
	Name           string               `description:"이름"`
	Description    string               `description:"설명" sql:"type:text"`
	CircleSet      CircleSet            ``
	CircleSetID    uint                 ``
	ControllerName string               ``
	Properties     []CircleUnitProperty `description:""`
	Import         string               `description:""`
	URL            string               `description:""`
	MenuName       string               `description:""`
	MenuGroup      string               `description:""`
	IsEnable       bool                 `description:"사용여부"`
	IsManual       bool                 `description:""`
	IsSystem       bool                 `description:""`
}

type CircleUnitProperty struct {
	ID           uint       `description:""`
	CreatedAt    time.Time  `description:"등록일"`
	UpdatedAt    time.Time  `description:"수정일"`
	Name         string     `description:"이름"`
	Description  string     `description:"설명" sql:"type:text"`
	CircleUnit   CircleUnit `description:""`
	CircleUnitID uint       `description:""`
	Type         string     `description:""`
	Nullable     bool       `description:""`
	IsEnable     bool       `description:"사용여부"`
	IsManual     bool       `description:""`
	IsSystem     bool       `description:""`
}

func (c CircleUnit) GetVariableName() string {
	return stringcase.ToCamelCase(c.Name)
}

func SaveCircleUnit(nameAndValue map[string]interface{}) error {
	//TODO: ....
	return nil
}

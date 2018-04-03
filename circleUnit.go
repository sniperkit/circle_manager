package circle_manager

import "time"

type CircleUnit struct {
	ID           uint                 `description:""`
	CreatedAt    time.Time            `description:"등록일"`
	UpdatedAt    time.Time            `description:"수정일"`
	Name         string               `description:"이름"`
	Description  string               `description:"설명" sql:"type:text"`
	CircleSet    CircleSet            ``
	CircleSetID  uint                 ``
	Properties   []CircleUnitProperty `description:""`
	VariableName string               `description:""`
	Import       string               `description:""`
	Url          string               `description:""`
	MenuName     string               `description:""`
	MenuGroup    string               `description:""`
	IsEnable     bool                 `description:"사용여부"`
}

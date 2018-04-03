package circle_manager

import "time"

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
}

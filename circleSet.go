package circle_manager

import "time"

type CircleSet struct {
	ID          uint         `description:""`
	CreatedAt   time.Time    `description:"등록일"`
	UpdatedAt   time.Time    `description:"수정일"`
	Name        string       `description:"이름"`
	Description string       `description:"설명" sql:"type:text"`
	Import      string       `description:""`
	Units       []CircleUnit `description:""`
	IsEnable    bool         `description:"사용여부"`
}

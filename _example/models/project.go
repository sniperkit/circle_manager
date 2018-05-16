package models

	import (
		"time"
	
		"github.com/jungju/circle_manager/modules"
	)
	
	// gen:qs
	type Project struct {
		ID          uint      description:""`   ss
		CreatedAt   time.Time description:"등록일"`
		UpdatedAt   time.Time description:"수정일"`
		Name        string    description:"이름"`
		Description string    description:"설명" sql:"type:text"`
		CreatorID   uint      description:"작성자"`
		Status string description:""`
		
	}
	
	func init() {
		registModel(&Project{})
	}
	
	func (m *Project) GetCreatorID() uint {
		return m.CreatorID
	}
	
	func (m *Project) SetCreatorID(creatorID uint) {
		m.CreatorID = creatorID
	}
	
	func AddProject(project *Project) (id uint, err error) {
		err = project.Create(gGormDB)
		id = project.ID
		return
	}
	
	func GetProjectByID(id uint) (project *Project, err error) {
		project = &Project{
			ID: id,
		}
		err = NewProjectQuerySet(gGormDB).
			One(project)
		return
	}
	
	func GetAllProject(queryPage *modules.QueryPage) (projects []Project, err error) {
		err = NewProjectQuerySet(gGormDB).
			All(&projects)
		return
	}
	
	func UpdateProjectByID(project *Project) (err error) {
		err = project.Update(gGormDB,
			ProjectDBSchema.Name,
			ProjectDBSchema.Description,
			ProjectDBSchema.Status,
			
		)
		return
	}
	
	func DeleteProject(id uint) (err error) {
		project := &Project{
			ID: id,
		}
		err = project.Delete(gGormDB)
		return
	}
	
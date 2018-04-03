package circle_manager

import "github.com/jinzhu/gorm"

type CircleManager struct {
	db               *gorm.DB
	TemplatePath     string
	SourceRootPath   string
	ModelsPath       string
	RoutersPath      string
	ControllersPath  string
	QORAdminPath     string
	RequestsBodyPath string
	ResponseBodyPath string
	BuildPath        string
}

func New(db *gorm.DB) (*CircleManager, error) {
	cm := &CircleManager{
		db: db,
	}
	err := cm.db.AutoMigrate(
		&CircleSet{},
		&CircleUnit{},
		&CircleUnitProperty{},
	).Error
	if err != nil {
		return nil, err
	}

	if cm.TemplatePath == "" {
		cm.TemplatePath = "templates"
	}
	if cm.SourceRootPath == "" {
		cm.TemplatePath = ""
	}
	if cm.ModelsPath == "" {
		cm.TemplatePath = "models"
	}
	if cm.RoutersPath == "" {
		cm.TemplatePath = "routers"
	}
	if cm.ControllersPath == "" {
		cm.TemplatePath = "contorllers"
	}
	if cm.QORAdminPath == "" {
		cm.TemplatePath = "admin"
	}
	if cm.RequestsBodyPath == "" {
		cm.TemplatePath = "requests"
	}
	if cm.ResponseBodyPath == "" {
		cm.TemplatePath = "responses"
	}

	return cm, nil
}

func (cm *CircleManager) GeneateSource(circleIDUint uint) error {
	cs, err := getCircleSetByID(cm.db, circleIDUint)
	if err != nil {
		return err
	}

	if err := cm.runGen(cs); err != nil {
		return err
	}
	return nil
}

func (cm *CircleManager) GeneateSourceBySet(cs *CircleSet) error {
	if err := cm.runGen(cs); err != nil {
		return err
	}
	return nil
}

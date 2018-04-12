package models

import "time"

// gen:qs
type Employee struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	OriginName  string    `description:""`
}

func init() {
	registModel(&Employee{})
}

func AddEmployee(employee *Employee) (id uint, err error) {
	err = employee.Create(gGormDB)
	id = employee.ID
	return
}

func GetEmployeeByID(id uint) (employee *Employee, err error) {
	employee = &Employee{
		ID: id,
	}
	err = NewEmployeeQuerySet(gGormDB).
		One(employee)
	return
}

func GetAllEmployee(queryPage *QueryPage) (employees []Employee, err error) {
	err = NewEmployeeQuerySet(gGormDB).
		All(&employees)
	return
}

func UpdateEmployeeByID(employee *Employee) (err error) {
	err = employee.Update(gGormDB,
		EmployeeDBSchema.Name,
		EmployeeDBSchema.Description,
		EmployeeDBSchema.OriginName,
	)
	return
}

func DeleteEmployee(id uint) (err error) {
	employee := &Employee{
		ID: id,
	}
	err = employee.Delete(gGormDB)
	return
}

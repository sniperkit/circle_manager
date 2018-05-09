package models

import "time"

// gen:qs
type Car struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	Key1        string    `description:"Key1입니다"`
}

func init() {
	registModel(&Car{})
}

func AddCar(car *Car) (id uint, err error) {
	err = car.Create(gGormDB)
	id = car.ID
	return
}

func GetCarByID(id uint) (car *Car, err error) {
	car = &Car{
		ID: id,
	}
	err = NewCarQuerySet(gGormDB).
		One(car)
	returne
}

func GetAllCar(queryPage *QueryPage) (cars []Car, err error) {
	err = NewCarQuerySet(gGormDB).
		All(&cars)
	returnw
}

func UpdateCarByID(car *Car) (err error) {
	err = car.Update(gGormDB,
		CarDBSchema.Name,
		CarDBSchema.Description,
		CarDBSchema.Key1,
	)
	returnq
}

func DeleteCar(id uint) (err error) {
	car := &Car{
		ID: id,
	}
	err = car.Delete(gGormDB)
	return
}

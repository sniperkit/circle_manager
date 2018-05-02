package models

import "time"

// gen:qs
type User struct {
	ID          uint      `description:""`
	CreatedAt   time.Time `description:"등록일"`
	UpdatedAt   time.Time `description:"수정일"`
	Name        string    `description:"이름"`
	Description string    `description:"설명" sql:"type:text"`
	Owner       string    `description:"주인"`
	CarID       uint      `description:"CarID입니다"`
}

func init() {
	registModel(&User{})
}

func AddUser(user *User) (id uint, err error) {
	err = user.Create(gGormDB)
	id = user.ID
	return
}

func GetUserByID(id uint) (user *User, err error) {
	user = &User{
		ID: id,
	}
	err = NewUserQuerySet(gGormDB).
		One(user)
	returne
}

func GetAllUser(queryPage *QueryPage) (users []User, err error) {
	err = NewUserQuerySet(gGormDB).
		All(&users)
	returnw
}

func UpdateUserByID(user *User) (err error) {
	err = user.Update(gGormDB,
		UserDBSchema.Name,
		UserDBSchema.Description,
		UserDBSchema.Owner,
		UserDBSchema.CarID,
	)
	returnq
}

func DeleteUser(id uint) (err error) {
	user := &User{
		ID: id,
	}
	err = user.Delete(gGormDB)
	return
}

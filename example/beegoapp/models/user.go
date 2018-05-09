package models

import (
	"time"

	"github.com/jungju/circle_manager/modules"
)

// User struct represent user model.
// gen:qs
type User struct {
	ID                 uint      `description:""`
	CreatedAt          time.Time `description:"등록일"`
	UpdatedAt          time.Time `description:"수정일"`
	Name               string    `description:"이름"`
	Description        string    `description:"설명" sql:"type:text"`
	CreatorID          uint      `description:"작성자"`
	Username           string    `description:"사용자 아이디"`
	Password           string    `description:"비밀번호"`
	EncryptedPassword  string    `description:"암호화된 비밀번호"`
	Email              string    `description:"사용자 이메일"`
	Mobile             string    `description:"사용자 핸드폰번호"`
	PosibleSendSMS     bool      `description:""`
	PosibleSendEmail   bool      `description:""`
	PosibleSendWeb     bool      `description:""`
	PosibleSendWebhook bool      `description:""`
}

func init() {
	registModel(&User{})
}

func (m *User) GetCreatorID() uint {
	return m.CreatorID
}

func (m *User) SetCreatorID(creatorID uint) {
	m.CreatorID = creatorID
}

func (u User) DisplayName() string {
	return u.Email
}

// AddUser insert a new User into database and returns
// last inserted ID on success.
func AddUser(user *User) (id uint, err error) {
	err = user.Create(gGormDB)
	id = user.ID
	return
}

// GetUserByID retrieves User by ID. Returns error if
// ID doesn't exist
func GetUserByID(id uint) (user *User, err error) {
	user = &User{
		ID: id,
	}
	err = NewUserQuerySet(gGormDB).
		One(user)
	return
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(queryPage *modules.QueryPage) (users []User, err error) {
	err = NewUserQuerySet(gGormDB).
		w(queryPage.ParsedQueryPage(gGormDB)).
		All(&users)
	return
}

// UpdateUser updates User by ID and returns error if
// the record to be updated doesn't exist
func UpdateUserByID(user *User) (err error) {
	err = user.Update(gGormDB,
		UserDBSchema.Username,
		UserDBSchema.Email,
		UserDBSchema.Name,
		UserDBSchema.Description,
		UserDBSchema.Mobile,
	)
	return
}

// DeleteUser deletes User by ID and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id uint) (err error) {
	user := &User{
		ID: id,
	}
	err = user.Delete(gGormDB)
	return
}

func GetUserByUsername(username string) (user *User, err error) {
	user = &User{}
	err = NewUserQuerySet(gGormDB).
		UsernameEq(username).
		One(user)
	return
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUserOnly() (users []User, err error) {
	err = NewUserQuerySet(gGormDB).
		All(&users)
	return
}

func GetOnlyUserByID(id uint) (user *User, err error) {
	user = &User{
		ID: id,
	}

	err = NewUserQuerySet(gGormDB).
		One(user)
	return
}

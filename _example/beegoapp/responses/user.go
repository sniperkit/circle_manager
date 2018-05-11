package responses

import "time"

type User struct {
	ID                 uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string
	Description        string
	CreatorID          uint
	Username           string
	Password           string
	EncryptedPassword  string
	Email              string
	Mobile             string
	PosibleSendSMS     bool
	PosibleSendEmail   bool
	PosibleSendWeb     bool
	PosibleSendWebhook bool
}

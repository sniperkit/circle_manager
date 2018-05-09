package requests

type CreateUser struct {
	Username    string `validate:"required,min=2,max=36"`
	Password    string `validate:"required,min=6,max=36"`
	Email       string ``
	Name        string `validate:"required,min=1,max=36"`
	Description string ``
	Mobile      string ``
}
type UpdateUser struct {
	Username    string `validate:"required,min=2,max=36"`
	Password    string `validate:"required,min=6,max=36"`
	Email       string ``
	Name        string `validate:"required,min=1,max=36"`
	Description string ``
	Mobile      string ``
}

func (c *CreateUser) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateUser) Valid() error {
	return validate.Struct(c)
}

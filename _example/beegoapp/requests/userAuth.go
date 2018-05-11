package requests

type LoginUser struct {
	Username string
	Password string
}

type RegistUser struct {
	Username string
	Password string
}

type CheckUsername struct {
	Username string
}

func (c *LoginUser) Valid() error {
	return validate.Struct(c)
}

func (c *RegistUser) Valid() error {
	return validate.Struct(c)
}

func (c *CheckUsername) Valid() error {
	return validate.Struct(c)
}

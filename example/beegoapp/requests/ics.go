package requests

type CreateIcs struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

type UpdateIcs struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

func (c *CreateIcs) Valid() error {
  return validate.Struct(c)
}

func (c *UpdateIcs) Valid() error {
  return validate.Struct(c)
}

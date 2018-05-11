package requests

type CreateSprint struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

type UpdateSprint struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

func (c *CreateSprint) Valid() error {
  return validate.Struct(c)
}

func (c *UpdateSprint) Valid() error {
  return validate.Struct(c)
}

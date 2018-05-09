package requests

type CreateTrello struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

type UpdateTrello struct {
  Name        string  `validate:"required,min=2,max=36"`
  Description string
}

func (c *CreateTrello) Valid() error {
  return validate.Struct(c)
}

func (c *UpdateTrello) Valid() error {
  return validate.Struct(c)
}

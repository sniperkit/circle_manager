package requests

type CreateCar struct {
	Key1 string
}

type UpdateCar struct {
	Key1 string
}

func (c *CreateCar) Valid() error {
	return validate.Struct(c)
}

func (c *UpdateCar) Valid() error {
	return validate.Struct(c)
}

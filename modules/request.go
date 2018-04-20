package modules

import validator "gopkg.in/go-playground/validator.v9"

func init() {
	validate = validator.New()
}

// RequestBody ...
type RequestBody interface {
	Valid() error
}

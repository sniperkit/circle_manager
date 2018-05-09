package errors

import "errors"

var (
	ErrDataNotFound        = errors.New("ErrDataNotFound")
	ErrDulicatedObject     = errors.New("ErrDulicatedObject")
	ErrInvalidQuery        = errors.New("ErrInvalidQuery")
	ErrUnauthorized        = errors.New("ErrUnauthorized")
	ErrInvalidRequestBody  = errors.New("ErrInvalidRequestBody")
	ErrInvalidRequestParam = errors.New("ErrInvalidRequestParam")
	ErrInvalidUsername     = errors.New("ErrInvalidUsername")
	ErrInvalidToken        = errors.New("ErrInvalidToken")
	ErrUnknown             = errors.New("Unknown")
)

func New(errMsg string) error {
	return errors.New(errMsg)
}

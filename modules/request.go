package modules

// RequestBody ...
type RequestBody interface {
	Valid() error
}

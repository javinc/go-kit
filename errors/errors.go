package errors

type err struct {
	svc string
	msg string
	err error
	pan bool
}

// error interface implemetation
func (i *err) Error() string {
	return i.svc + ": " + i.msg + "(" + i.err.Error() + ")"
}

// New init normal error
func New(name, msg string, e error) error {
	return create(name, msg, e, true)
}

// NewPanic init panic error
func NewPanic(name, msg string, e error) error {
	return create(name, msg, e, false)
}

// create new error
func create(name, msg string, e error, pan bool) error {
	return &err{name, msg, e, pan}
}

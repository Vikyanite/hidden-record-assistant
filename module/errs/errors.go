package errs

import "errors"

var (
	ErrNeedAdmin      = errors.New("need admin permission")
	ErrCantFindClient = errors.New("can't find lol client")
)

func InternalError(err error) error {
	return errors.New("internal error: " + err.Error())
}

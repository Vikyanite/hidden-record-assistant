package errs

import "github.com/pkg/errors"

var (
	ErrNeedAdmin      = errors.New("need admin permission")
	ErrCantFindClient = errors.New("can't find lol client")
	ErrNotRunning     = errors.New("network.Client is not running")
)

func InternalError(err error) error {
	return errors.New("internal error: " + err.Error())
}

func UnknownError(err error) error {
	return errors.New("internal error: " + err.Error())
}

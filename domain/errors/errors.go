package errors

import "errors"

var (
	NotFound        = errors.New("not found")
	InvalidArgument = errors.New("invalid argument")
	Internal        = errors.New("internal error")
)
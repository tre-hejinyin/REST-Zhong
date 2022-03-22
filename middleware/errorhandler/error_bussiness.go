package errorhandler

import (
	"errors"
)

var (
	ErrorEmailOrPassword = errors.New("email or password error")
)

const (
	CodeEmailOrPasswordError = 1001 + iota // email or password is wrong
)

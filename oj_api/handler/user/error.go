package user

import (
	"errors"
)

var (
	PasswdIncorrect = errors.New("password incorrect")
	NameNotExist = errors.New("can't find the user name")
	
)

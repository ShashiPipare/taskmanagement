package user

import "errors"

var (
	ErrMissingRequiredParams = errors.New("required keys missing")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrIncorrectPassword     = errors.New("incorrect password")
	ErrAuthUser              = errors.New("error in authenticating user")
	ErrEmailParamMandatory   = errors.New("missing mandatory key : email_id")
)

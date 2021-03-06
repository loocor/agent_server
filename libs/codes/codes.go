package codes

import (
	"google.golang.org/grpc/codes"
)

const (
	// Common
	Unknown         codes.Code = 1001
	InvalidArgument codes.Code = 1002
	DBInsertError   codes.Code = 1011
	DBUpdateError   codes.Code = 1012
	DBDeleteError   codes.Code = 1013

	// Reg
	PhoneAlreadyExists    codes.Code = 2001
	IDNumberAlreadyExists codes.Code = 2002

	// Login
	PasswordWrong codes.Code = 2011
	UserNotFound  codes.Code = 2012
	UserFreeze    codes.Code = 2013
	UserRemoved   codes.Code = 2014
)

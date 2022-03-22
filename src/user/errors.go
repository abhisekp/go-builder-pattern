package user

import (
	"errors"
	"fmt"

	errorConstants "builder-pattern/src/constants/errors"
)

type BuilderError struct {
	error
	Code int64
}

func (be BuilderError) String(outputFormat string) string {
	switch outputFormat {
	case "json":
		return fmt.Sprintf("{\"code\": %d, \"message\": \"%s\"}", be.Code, be.Error())
	default:
		return fmt.Sprintf("Error Code: %d\nError Message: %s\n", be.Code, be.Error())
	}
}

var (
	ErrNameEmpty = BuilderError{
		error: errors.New("name is empty"),
		Code:  errorConstants.ErrNameEmpty,
	}

	ErrNameInvalid = BuilderError{
		error: errors.New("name is invalid"),
		Code:  errorConstants.ErrNameInvalid,
	}

	ErrGenderInvalid = BuilderError{
		error: errors.New("gender is invalid"),
		Code:  errorConstants.ErrGenderInvalid,
	}

	ErrAgeInvalid = BuilderError{
		error: errors.New("age is invalid"),
		Code:  errorConstants.ErrAgeInvalid,
	}

	ErrRoleInvalid = BuilderError{
		error: errors.New("role is invalid"),
		Code:  errorConstants.ErrRoleInvalid,
	}

	ErrAddressInvalid = BuilderError{
		error: errors.New("address is invalid"),
		Code:  errorConstants.ErrAddressInvalid,
	}
)

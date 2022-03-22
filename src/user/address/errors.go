package address

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
	ErrLine1Empty = BuilderError{
		error: errors.New("line 1 is empty"),
		Code:  errorConstants.ErrLine1Empty,
	}

	ErrLine2Empty = BuilderError{
		error: errors.New("line 2 is empty"),
		Code:  errorConstants.ErrLine2Empty,
	}

	ErrCityEmpty = BuilderError{
		error: errors.New("city is empty"),
		Code:  errorConstants.ErrCityEmpty,
	}

	ErrStateEmpty = BuilderError{
		error: errors.New("state is empty"),
		Code:  errorConstants.ErrStateEmpty,
	}

	ErrCountryEmpty = BuilderError{
		error: errors.New("country is empty"),
		Code:  errorConstants.ErrCountryEmpty,
	}

	ErrPincodeEmpty = BuilderError{
		error: errors.New("pincode is empty"),
		Code:  errorConstants.ErrPincodeEmpty,
	}

	ErrStateInvalid = BuilderError{
		error: errors.New("state is invalid"),
		Code:  errorConstants.ErrStateInvalid,
	}

	ErrCountryInvalid = BuilderError{
		error: errors.New("country is invalid"),
		Code:  errorConstants.ErrCountryInvalid,
	}

	ErrPincodeInvalid = BuilderError{
		error: errors.New("pincode is invalid"),
		Code:  errorConstants.ErrPincodeInvalid,
	}
)

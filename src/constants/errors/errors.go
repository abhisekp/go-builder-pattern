package errors

// Address Error Constants
const (
	ErrLine1Empty = (iota+1)*1000 + 1
	ErrLine2Empty
	ErrCityEmpty
	ErrStateEmpty
	ErrCountryEmpty
	ErrPincodeEmpty
	ErrStateInvalid
	ErrCountryInvalid
	ErrPincodeInvalid
)

// User Error Constants
const (
	ErrNameEmpty = (iota+1)*2000 + 1
	ErrNameInvalid
	ErrGenderInvalid
	ErrAgeInvalid

	ErrRoleInvalid
	ErrAddressInvalid
)

package address

import "sync"

type Builder struct {
	sync.Mutex
	address Address
	error   *BuilderError
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetLine1(line1 string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetLine1(line1)
	return b
}

func (b *Builder) SetLine2(line2 string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetLine2(line2)
	return b
}

func (b *Builder) SetCity(city string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetCity(city)
	return b
}

func (b *Builder) SetState(state string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetState(state)
	return b
}

func (b *Builder) SetPincode(pincode string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetPincode(pincode)
	return b
}

func (b *Builder) SetCountry(country string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.address.SetCountry(country)
	return b
}

func (b *Builder) Build() *Address {
	b.Lock()
	defer b.Unlock()

	if b.address.Line1() == "" {
		b.error = &ErrLine1Empty
		return nil
	}

	if b.address.Pincode() == "" {
		b.error = &ErrPincodeEmpty
		return nil
	}

	if b.address.City() == "" {
		b.error = &ErrCityEmpty
		return nil
	}

	if b.address.State() == "" {
		b.error = &ErrStateEmpty
		return nil
	}

	if b.address.Country() == "" {
		b.error = &ErrCountryEmpty
		return nil
	}

	return &b.address
}

func (b *Builder) Error() *BuilderError {
	b.Lock()
	defer b.Unlock()

	return b.error
}

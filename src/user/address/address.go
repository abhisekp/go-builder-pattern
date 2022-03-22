package address

import "sync"

type Address struct {
	sync.Mutex
	line1   string
	line2   string
	city    string
	state   string
	pincode string
	country string
}

func (a *Address) SetLine1(line1 string) *Address {
	a.Lock()
	defer a.Unlock()

	a.line1 = line1
	return a
}

func (a *Address) SetLine2(line2 string) *Address {
	a.Lock()
	defer a.Unlock()

	a.line2 = line2
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.Lock()
	defer a.Unlock()

	a.city = city
	return a
}

func (a *Address) SetState(state string) *Address {
	a.Lock()
	defer a.Unlock()

	a.state = state
	return a
}

func (a *Address) SetPincode(pincode string) *Address {
	a.Lock()
	defer a.Unlock()

	a.pincode = pincode
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.Lock()
	defer a.Unlock()

	a.country = country
	return a
}

// \/ Getters \/ //

func (a *Address) Line1() string {
	a.Lock()
	defer a.Unlock()

	return a.line1
}

func (a *Address) Line2() string {
	a.Lock()
	defer a.Unlock()

	return a.line2
}

func (a *Address) City() string {
	a.Lock()
	defer a.Unlock()

	return a.city
}

func (a *Address) State() string {
	a.Lock()
	defer a.Unlock()

	return a.state
}

func (a *Address) Pincode() string {
	a.Lock()
	defer a.Unlock()

	return a.pincode
}

func (a *Address) Country() string {
	a.Lock()
	defer a.Unlock()

	return a.country
}

func (a *Address) String() string {
	msg := ""

	if a.line1 != "" {
		msg += a.line1
	}

	if a.line2 != "" {
		msg += ", " + a.line2
	}

	if a.city != "" {
		msg += ", " + a.city
	}

	if a.state != "" {
		msg += ", " + a.state
	}

	if a.pincode != "" {
		msg += ", " + a.pincode
	}

	if a.country != "" {
		msg += ", " + a.country
	}

	return msg
}

func (a *Address) IsEmpty() bool {
	return a.line1 == "" && a.line2 == "" && a.city == "" && a.state == "" && a.pincode == "" && a.country == ""
}

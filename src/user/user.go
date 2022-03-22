package user

import (
	"fmt"
	"strings"
	"sync"

	Address "builder-pattern/src/user/address"
	. "builder-pattern/src/user/gender"
	"builder-pattern/src/user/role"
)

type User struct {
	sync.Mutex
	name   string
	age    int64
	gender Gender
	// addressBuilder Address.Builder
	address *Address.Address
	roles   []role.Role
}

func (u *User) WithName(name string) *User {
	u.Lock()
	defer u.Unlock()

	u.name = name
	return u
}

func (u *User) WithAge(age int64) *User {
	u.Lock()
	defer u.Unlock()

	u.age = age
	return u
}

func (u *User) WithGender(gender Gender) *User {
	u.Lock()
	defer u.Unlock()

	if gender == Male || gender == Female {
		u.gender = Gender(strings.Title(string(gender)))
	}

	return u
}

func (u *User) withAddress(address *Address.Address) *User {
	u.Lock()
	defer u.Unlock()

	u.address = address
	return u
}

func (u *User) WithAddressBuilder(addressBuilder *Address.Builder) *User {
	u.Lock()
	defer u.Unlock()

	address := addressBuilder.Build()
	if address != nil {
		u.address = address
	}

	return u
}

func (u *User) WithRole(role role.Role) *User {
	u.Lock()
	defer u.Unlock()

	u.roles = append(u.roles, role)
	return u
}

// \/ Getters \/ //

func (u *User) Name() string {
	u.Lock()
	defer u.Unlock()

	return u.name
}

func (u *User) Age() int64 {
	u.Lock()
	defer u.Unlock()

	return u.age
}

func (u *User) Gender() Gender {
	u.Lock()
	defer u.Unlock()

	return u.gender
}

func (u *User) Address() *Address.Address {
	u.Lock()
	defer u.Unlock()

	return u.address
}

func (u *User) Roles() []role.Role {
	u.Lock()
	defer u.Unlock()

	return u.roles
}

func (u *User) String(outputFormat string) string {
	msg := ""

	name := u.Name()
	if name != "" {
		msg += fmt.Sprintf("Name: %s\n", name)
	}

	if u.Age() != 0 {
		msg += fmt.Sprintf("Age: %d\n", u.Age())
	}

	address := u.Address()
	if address != nil && !address.IsEmpty() {
		msg += fmt.Sprintf("Address: %s\n", address.String())
	}

	// roles := u.Roles()
	// if len(roles) > 0 {
	// 	msg += fmt.Sprintf("Roles: %s\n", strings.Join(role.RolesToString(roles), ", "))
	// }

	if u.Gender() != "" {
		msg += fmt.Sprintf("Gender: %s\n", u.Gender())
	}

	return msg
}

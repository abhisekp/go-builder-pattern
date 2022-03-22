package user

import (
	"sync"

	Address "builder-pattern/src/user/address"
	"builder-pattern/src/user/gender"
	userRole "builder-pattern/src/user/role"
)

type Builder struct {
	sync.Mutex
	user  User
	error *BuilderError
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetName(name string) *Builder {
	b.Lock()
	defer b.Unlock()

	b.user.WithName(name)
	return b
}

func (b *Builder) SetAge(age int64) *Builder {
	b.Lock()
	defer b.Unlock()

	b.user.WithAge(age)
	return b
}

func (b *Builder) SetAddress(address *Address.Address) *Builder {
	b.Lock()
	defer b.Unlock()

	b.user.withAddress(address)
	return b
}

func (b *Builder) SetRoles(roles []userRole.Role) *Builder {
	b.Lock()
	defer b.Unlock()

	for _, role := range roles {
		b.user.WithRole(role)
	}

	return b
}

func (b *Builder) SetGender(gender gender.Gender) *Builder {
	b.Lock()
	defer b.Unlock()

	b.user.WithGender(gender)
	return b
}

func (b *Builder) Build() *User {
	user := &b.user

	if user.name == "" {
		b.error = &ErrNameEmpty
		return nil
	}

	return user
}

func (b *Builder) Error() *BuilderError {
	b.Lock()
	defer b.Unlock()

	return b.error
}

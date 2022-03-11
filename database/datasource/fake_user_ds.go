package datasource

import (
	"github.com/PepperTalk/api/database/model"
	"strings"
)

type FakeUserDS struct {
	Users *[]*model.User
}

func (f FakeUserDS) GetByFirstname(firstname string, dest *[]*model.User) {
	result := model.Filter(*f.Users, func(user model.User) bool {
		return user.Firstname == firstname
	})
	*dest = result
}

func (f FakeUserDS) GetByEmail(email string, dest *[]*model.User) {
	result := model.Filter(*f.Users, func(user model.User) bool {
		return user.Email == email
	})
	*dest = result
}

func (f FakeUserDS) GetByUsernameLike(username string, dest *[]*model.User) {
	result := model.Filter(*f.Users, func(user model.User) bool {
		return strings.Contains(user.Username, username)
	})
	*dest = result
}

func (f FakeUserDS) Create(user *model.User) error {
	*f.Users = append(*f.Users, user)
	return nil
}

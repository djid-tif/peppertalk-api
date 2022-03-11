package datasource

import "github.com/PepperTalk/api/database/model"

type IUserDatasource interface {
	GetByFirstname(firstname string, dest *[]*model.User)
	GetByUsernameLike(username string, dest *[]*model.User)
	GetByEmail(email string, dest *[]*model.User)
	Create(user *model.User) error
}

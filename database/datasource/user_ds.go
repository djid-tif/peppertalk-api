package datasource

import (
	"github.com/PepperTalk/api/database/model"
	"gorm.io/gorm"
)

type DatabaseDS struct {
	DB *gorm.DB
}

func (d DatabaseDS) GetByFirstname(firstname string, dest *[]*model.User) {
	d.DB.Where("firstname = ?", firstname).Find(dest)
}
func (d DatabaseDS) GetByEmail(email string, dest *[]*model.User) {
	d.DB.Where("email = ?", email).Find(dest)
}

func (d DatabaseDS) GetByUsernameLike(username string, dest *[]*model.User) {
	d.DB.Where("username LIKE ?", "%"+username+"%").Find(dest)
}

func (d DatabaseDS) Create(user *model.User) error {
	return d.DB.Create(&user).Error
}

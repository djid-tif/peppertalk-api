package model

import (
	"math/rand"

	"github.com/lib/pq"
	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&UserSettings{})
}

// UserSettings represents the settings of a user, which is represented by the UserId
type UserSettings struct {
	UserID              uint           `gorm:"primarykey;autoIncrement;unique"`
	Lang                string         `gorm:"type:char(3)"`
	NotificationChannel pq.StringArray `gorm:"type:char(32)[]"`
}

func UserSettingsGenerator() interface{} {
	return &UserSettings{
		UserID:              uint(rand.Intn(1234567)),
		Lang:                []string{"fr", "en", "zh", "jp"}[rand.Intn(3)],
		NotificationChannel: []string{[]string{"push-mobile", "push-desktop", "email", "sms"}[rand.Intn(3)]},
	}
}

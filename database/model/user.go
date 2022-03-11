package model

import (
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"goyave.dev/goyave/v3/database"
	"time"
)

// A model is a structure reflecting a database table structure. An instance of a model
// is a single database record. Each model is defined in its own file inside the database/models directory.
// Models are usually just normal Golang structs, basic Go types, or pointers of them.
// "sql.Scanner" and "driver.Valuer" interfaces are also supported.

// Learn more here: https://goyave.dev/guide/basics/database.html#models

func init() {
	// All models should be registered in an "init()" function inside their model file.
	database.RegisterModel(&User{})
}

// User represents a user.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid"` // ;default:uuid_generate_v4()
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Email     string       `gorm:"type:text;uniqueIndex"`
	Username  string       `gorm:"type:text;uniqueIndex"`
	Firstname string       `gorm:"type:text;"`
	Lastname  string       `gorm:"type:text;"`
	Password  string       `gorm:"type:char(60);"` // 60 for bcrypt
}

func Filter(arr []*User, test func(User) bool) (ret []*User) {
	for _, s := range arr {
		if test(*s) {
			ret = append(ret, s)
		}
	}
	return
}

// You may need to test features interacting with your database.
// Goyave provides a handy way to generate and save records in your database: factories.
// Factories need a generator function. These functions generate a single random record.
//
// "database.Generator" is an alias for "func() interface{}"
//
// Learn more here: https://goyave.dev/guide/advanced/testing.html#database-testing

// UserGenerator generator function for the User model.
// Generate users using the following:
//  database.NewFactory(model.UserGenerator).Generate(5)
func UserGenerator() interface{} {
	user := &User{}
	user.Firstname = faker.FirstName()
	user.Lastname = faker.LastName()
	user.Username = user.Firstname + " " + user.Lastname
	user.Password = faker.Password()

	faker.SetGenerateUniqueValues(true)
	user.ID = uuid.MustParse(faker.UUIDDigit())
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false)
	return user
}

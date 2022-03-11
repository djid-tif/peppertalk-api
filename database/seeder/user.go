package seeder

import (
	"github.com/PepperTalk/api/database/model"
	"github.com/bxcodec/faker/v3"
	"goyave.dev/goyave/v3/database"
)

// Seeders are functions which create a number of random records in the database
// in order to create a full and realistic test environment.
//
// Each seeder should have its own file.
// A seeder's responsibilities are limited to a single table or model.
// For example, the "seeder.User" should only seed the "users" table.
// Moreover, seeders should have the same name as the model they are using.
//
// Learn more here: https://goyave.dev/guide/advanced/testing.html#seeders

// User seeder for users. Generate and save 10 users in the database.
func User() {

	database.NewFactory(model.UserGenerator).Save(1)

	/*user := &model.User{}
	user.Firstname = faker.Username()
	user.Lastname = faker.Username()
	user.Username = user.Firstname + "." + user.Lastname
	user.Password = faker.Password()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	database.Conn().Create(user)*/

	// As user generator makes unique emails,
	// forget generated unique emails.
	// See https://github.com/bxcodec/faker/blob/master/SingleFakeData.md#unique-values
	faker.ResetUnique()
}
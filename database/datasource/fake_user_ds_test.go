package datasource

import (
	"fmt"
	"github.com/PepperTalk/api/database/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
The only purpose of this file was to understand testing in general and should not be used as a real
test file since it's testing a faked Datasource
*/

func generateUsers(dest *[]*model.User) {
	*dest = append(*dest, &model.User{
		ID:        uuid.New(),
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Password:  "#johnIsZeBest69",
		Username:  "John.Doe",
	},
		&model.User{
			ID:        uuid.New(),
			Firstname: "Johnny",
			Lastname:  "Doe",
			Email:     "johnny.doe@gmail.com",
			Password:  "#johnIsZeBest70",
			Username:  "Johnny.Doe2",
		},
	)
}

func TestFakeUserDS_Create(t *testing.T) {
	ds := FakeUserDS{Users: &[]*model.User{}}
	_ = ds.Create(&model.User{
		ID:        uuid.New(),
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@gmail.com",
		Password:  "#johnIsZeBest69",
		Username:  "John" + "." + "",
	})
	fmt.Println(ds.Users)
	assert.NotEmpty(t, ds.Users)
}

func TestFakeUserDS_GetByFirstname(t *testing.T) {
	users := &[]*model.User{}
	generateUsers(users)
	ds := FakeUserDS{Users: users}
	result := &[]*model.User{}
	ds.GetByFirstname("John", result)
	assert.NotEmpty(t, result, *result)
	assert.Equal(t, "John.Doe", (*result)[0].Username)
	assert.Equal(t, "#johnIsZeBest69", (*result)[0].Password)
}

func TestDatabaseDS_GetByUsernameLike(t *testing.T) {
	users := &[]*model.User{}
	generateUsers(users)
	ds := FakeUserDS{Users: users}
	result := &[]*model.User{}
	ds.GetByUsernameLike("ohn", result)
	assert.Equal(t, 2, len(*result))
}

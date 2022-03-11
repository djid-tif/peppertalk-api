package auth

import (
	"github.com/PepperTalk/api/database/datasource"
	"github.com/PepperTalk/api/database/model"
	"github.com/PepperTalk/api/test"
	"github.com/stretchr/testify/assert"
	"goyave.dev/goyave/v3"
	"testing"
)

func setRegisterFields(req *goyave.Request) {
	req.Data = map[string]interface{}{
		"email":     "john.doe@gmail.com",
		"firstname": "John",
		"lastname":  "Doe",
		"password":  "#johnIsZeBest69",
	}
}

func Test_Register(t *testing.T) {
	request, response := test.GetReqAndRes()
	setRegisterFields(request)
	ds := datasource.FakeUserDS{Users: &[]*model.User{}}
	Register(response, request, ds)
	correspondingUsers := &[]*model.User{}
	ds.GetByFirstname("John", correspondingUsers)
	assert.NotEmpty(t, correspondingUsers, ds.Users)
}

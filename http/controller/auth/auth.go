package auth

import (
	"github.com/PepperTalk/api/database/datasource"
	"github.com/PepperTalk/api/database/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
	"net/http"
	"strconv"
)

func LoginHandler(response *goyave.Response, _ *goyave.Request) {
	_ = response.String(http.StatusOK, "Login")
}

func RegisterHandler(response *goyave.Response, request *goyave.Request) {
	db := datasource.DatabaseDS{DB: database.Conn()}
	Register(response, request, db)
}

func Register(response *goyave.Response, request *goyave.Request, db datasource.IUserDatasource) {
	user := &model.User{
		ID:        uuid.New(),
		Firstname: request.String("firstname"),
		Lastname:  request.String("lastname"),
		Email:     request.String("email"),
		Password:  request.String("password"),
		Username:  request.String("firstname") + "." + request.String("lastname"),
	}

	users := &[]*model.User{}

	// Check if email already exists in DB
	db.GetByEmail(user.Firstname, users)
	if len(*users) != 0 {
		_ = response.JSON(http.StatusBadRequest, "this email is already used !")
		return
	}

	users = &[]*model.User{}

	// Check if username already exists in DB
	db.GetByUsernameLike(user.Firstname, users)
	if len(*users) != 0 {
		user.Username += strconv.Itoa(len(*users))
	}

	if !isValidPassWord(user.Password) {
		_ = response.JSON(http.StatusBadRequest, "respect the password rules ! (len >= 6, 1 Upper, 1 Lower, 1 Number, 1 Special)")
		return
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(bytes)

	if err := db.Create(user); err != nil {
		_ = response.Error(err)
	} else {
		_ = response.JSON(http.StatusCreated, user)
	}
}

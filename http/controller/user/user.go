package user

import (
	"github.com/PepperTalk/api/http/data"
	"goyave.dev/goyave/v3"
	"net/http"
)

func UsernameHandler(response *goyave.Response, request *goyave.Request) {
	user, err := data.ExtractUserFromRequest(request, "id")
	if err != nil {
		response.JSON(http.StatusBadRequest, data.ResponseFrom(err.Error(), nil))
		return
	}
	response.JSON(http.StatusOK, data.ResponseFrom("", data.MapFrom("username", user.Username)))
}

func ProfileHandler(response *goyave.Response, request *goyave.Request) {
	// user := request.User.(*model.User) <=== Good
	if !request.Has("id") {
		response.JSON(http.StatusBadRequest, data.ResponseFrom("no id provided", nil)) // temp
		return
	}
	user, found := data.LoadUserById(request.String("id")) // temp
	if !found {
		response.JSON(http.StatusBadRequest, data.ResponseFrom("user not found", nil)) // temp
		return
	}
	/*user, err := data.ExtractUserFromRequest(request, "id")
	if err != nil {
		response.JSON(http.StatusBadRequest, data.ResponseFrom(err.Error(), nil))
		return
	}*/
	response.JSON(http.StatusOK, data.ResponseFrom("", data.MapFrom("firstname", user.Firstname, "lastname", user.Lastname, "username", user.Username, "email", user.Email)))
}

func AvatarHandler(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "Avatar")
}

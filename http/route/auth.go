package route

import (
	"github.com/PepperTalk/api/http/controller/auth"
	"goyave.dev/goyave/v3"
)

func authHandler(router *goyave.Router) {

	router.Post("/login", auth.LoginHandler)
	router.Post("/register", auth.RegisterHandler).Validate(auth.RegisterRequest)

}

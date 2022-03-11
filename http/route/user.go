package route

import (
	"github.com/PepperTalk/api/http/controller/user"
	"goyave.dev/goyave/v3"
)

func userHandler(router *goyave.Router) {

	router.Get("/{id}/username", user.UsernameHandler).Middleware() // TODO: auth required
	router.Post("/profile", user.ProfileHandler).Middleware()       // TODO: self-auth required
	router.Get("/{id}/avatar", user.AvatarHandler)                  // TODO: auth required

}

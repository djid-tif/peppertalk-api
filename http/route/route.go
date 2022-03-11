package route

import (
	"github.com/PepperTalk/api/http/controller/hello"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/cors"
	"net/http"
)

// Routing is an essential part of any Goyave application.
// Routes definition is the action of associating a URI, sometimes having
// parameters, with a handler which will process the request and respond to it.

// Routes are defined in routes registrer functions.
// The main route registrer is passed to "goyave.Start()" and is executed
// automatically with a newly created root-level router.

const allMethods = http.MethodGet + "|" + http.MethodPost + "|" + http.MethodPut + "|" + http.MethodPatch + "|" + http.MethodDelete + "|" + http.MethodOptions + "|" + http.MethodHead

// Register all the application routes. This is the main route registrer.
func Register(router *goyave.Router) { // Applying default CORS settings (allow all methods and all origins)
	// Learn more about CORS options here: https://goyave.dev/guide/advanced/cors.html
	router.CORS(cors.Default())

	// Register your routes here

	// Route without validation
	router.Get("/hello", hello.SayHi)
	authHandler(router.Subrouter("/auth"))

	// authenticator := auth.Middleware(&model.User{}, &auth.BasicAuthenticator{})
	// router.Middleware(authenticator)

	// Route with validation
	router.Post("/echo", hello.Echo)
	userHandler(router.Subrouter("/user"))
}

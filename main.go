package main

import (
	"os"

	"github.com/PepperTalk/api/http/route"
	_ "github.com/PepperTalk/api/http/validation"

	"goyave.dev/goyave/v3"
	// Import the appropriate GORM dialect for the database.
	_ "goyave.dev/goyave/v3/database/dialect/postgres"
)

func main() {
	// This is the entry point of the application.

	/*if !config.IsLoaded() {
		if err := config.Load(); err != nil {
			goyave.ErrLogger.Println(err)
			os.Exit(1)
		}
	}

	// enable default uuid generation method
	database.Conn().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA ?;", config.GetString("database.name"))*/

	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/envy"

	"github.com/invitation/actions"
)

var logFile *os.File

// main is the starting point to your Buffalo application.
// you can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	logFile, err := os.OpenFile("/var/log/invitation-factory.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	if envy.Get("GO_ENV", "development") == "test" {
		//logFile.Close()
		//logFile, _ = os.OpenFile(envy.GoPath()+"/log/invitation-factory.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	app := actions.App()
	log.Println("Invitation-Factory started.")
	if err := app.Serve(); err != nil {
		log.Println(err)
		logFile.Close()
		return
	}
	log.Println("Invitation-Factory stopped.")
	logFile.Close()
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heaving lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/

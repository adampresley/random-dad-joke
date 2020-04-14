/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/adampresley/random-dad-joke/api/jokes"
	"github.com/adampresley/random-dad-joke/api/version"
	"github.com/adampresley/random-dad-joke/assets"
)

/*
IApplication defines an interface for running the main application
*/
type IApplication interface {
	Start() chan os.Signal
	Stop()
}

/*
Application provides an implementation of the main application
*/
type Application struct {
	Config          *viper.Viper
	HTTPServer      *echo.Echo
	Logger          *logrus.Entry
	ShutdownContext context.Context

	/*
	 * Controllers
	 */
	JokeController    jokes.IJokeController
	VersionController *version.VersionController

	/*
	 * Services
	 */
	JokeService jokes.JokeServicer
}

/*
NewApplication is the factory method to create a new Application
*/
func NewApplication(shutdownContext context.Context, logger *logrus.Entry, config *viper.Viper) *Application {
	result := &Application{
		Config:          config,
		Logger:          logger,
		ShutdownContext: shutdownContext,
	}

	result.setupServices()
	result.setupControllers()
	result.setupHandlers()

	return result
}

/*
handleMainPage is what serves the primary HTML container. Modify this to include
new scripts, CSS, or change the title
*/
func (a *Application) handleMainPage(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />

	<title>Random Dad Joke</title>

	<link rel="stylesheet" type="text/css" href="/app/assets/bootstrap/css/bootstrap.min.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/random-dad-joke/css/styles.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/vue-loading-overlay/vue-loading-overlay.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/fontawesome-5.11.2/css/all.css" />
</head>

<body>
	<div id="app"></div>

	<script src="/app/assets/babel/babel.min.js"></script>
	<script src="/app/assets/moment/moment.min.js"></script>
	<script src="/app/assets/vue/vue-2.6.10.js"></script>
	<script src="/app/assets/vue-router/vue-router-3.1.3.min.js"></script>
	<script src="/app/assets/vue-resource/vue-resource-1.5.1.min.js"></script>
	<script src="/app/assets/jquery/jquery-3.4.1.min.js"></script>
	<script src="/app/assets/popper/popper.min.js"></script>
	<script src="/app/assets/bootstrap/js/bootstrap.min.js"></script>
	<script src="/app/assets/vue-loading-overlay/vue-loading-overlay.js"></script>

	<script src="/app/main.js" type="module"></script>
	</body>
</html>
`)
}

/*
setupControllers is where you will initialize controllers that handle
your API routes
*/
func (a *Application) setupControllers() {
	a.JokeController = jokes.NewJokeController(jokes.JokeControllerConfig{
		Config:      a.Config,
		JokeService: a.JokeService,
		Logger:      a.Logger,
	})

	a.VersionController = version.NewVersionController(&version.VersionControllerConfig{
		Config: a.Config,
	})
}

/*
setupHandlers is where API routes are managed. Here you wire up a URL
route to a controller function
*/
func (a *Application) setupHandlers() {
	a.HTTPServer = echo.New()
	a.HTTPServer.HideBanner = true
	a.HTTPServer.Use(middleware.CORS())

	api := a.HTTPServer.Group("/api")

	a.HTTPServer.GET("/app/*", echo.WrapHandler(http.FileServer(assets.Assets)))
	a.HTTPServer.GET("/", a.handleMainPage)

	api.GET("/version", a.VersionController.GetVersion)
	api.GET("/joke/random", a.JokeController.GetRandomJoke)
}

/*
setupServices is where you will initialize an services that are
dependencies to controllers
*/
func (a *Application) setupServices() {
	a.JokeService = jokes.NewJokeService(jokes.JokeServiceConfig{
		Config: a.Config,
	})
}

/*
Start begins serving the API application
*/
func (a *Application) Start() chan os.Signal {
	go func() {
		var err error
		var host string

		if a.Config.GetString("PORT") != "" {
			host = fmt.Sprintf(":%s", a.Config.GetString("PORT"))
		} else {
			host = a.Config.GetString("server.host")
		}

		a.Logger.WithFields(logrus.Fields{
			"host":          host,
			"serverVersion": a.Config.GetString("server.version"),
			"debug":         a.Config.GetBool("server.debug"),
			"logLevel":      a.Config.GetString("fireplace.loglevel"),
		}).Infof("Starting")

		if err = a.HTTPServer.Start(host); err != nil {
			if err != http.ErrServerClosed {
				a.Logger.WithError(err).Fatalf("Unable to start application")
			} else {
				a.Logger.Infof("Shutting down the server...")
			}
		}
	}()

	/*
	 * Setup shutdown handler
	 */
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	return quit
}

/*
Stop halt API server execution. It waits for 10 seconds, and if
the server has not stopped a panic is thrown
*/
func (a *Application) Stop() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = a.HTTPServer.Shutdown(ctx); err != nil {
		a.Logger.WithError(err).Errorf("There was an error shutting down the server")
	}
}

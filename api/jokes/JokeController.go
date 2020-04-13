/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// IJokeController describes endpoints related to jokes
type IJokeController interface {
	GetRandomJoke(ctx echo.Context) error
}

// JokeControllerConfig is used to configure a JokeController
type JokeControllerConfig struct {
	Config      *viper.Viper
	JokeService JokeServicer
	Logger      *logrus.Entry
}

// JokeController implements endpoints related to jokes
type JokeController struct {
	Config      *viper.Viper
	JokeService JokeServicer
	Logger      *logrus.Entry
}

// NewJokeController creates a new JokeController
func NewJokeController(config JokeControllerConfig) *JokeController {
	return &JokeController{
		Config:      config.Config,
		JokeService: config.JokeService,
		Logger:      config.Logger,
	}
}

// GetRandomJoke returns a string with a random dad joke
func (c *JokeController) GetRandomJoke(ctx echo.Context) error {
	var err error
	var result *Joke

	if result, err = c.JokeService.GetRandomJoke(); err != nil {
		c.Logger.WithError(err).Error("Error getting random joke")
		return ctx.String(http.StatusInternalServerError, "Error getting random dad joke")
	}

	return ctx.JSON(http.StatusOK, result)
}

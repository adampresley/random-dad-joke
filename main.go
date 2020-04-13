/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

//go:generate go run -tags=dev assets_generate.go

package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/adampresley/random-dad-joke/configuration"
)

func main() {
	var err error
	var loglevel logrus.Level

	config := configuration.NewConfig("1.0.0")
	logger := logrus.New().WithField("who", "RandomDadJoke")

	if loglevel, err = logrus.ParseLevel(config.GetString("server.loglevel")); err != nil {
		logger.WithError(err).Fatal("Invalid log level")
	}

	logger.Logger.SetLevel(loglevel)

	shutdownContext, cancelFunc := context.WithCancel(context.Background())

	application := NewApplication(shutdownContext, logger, config)
	quit := application.Start()

	<-quit
	cancelFunc()

	application.Stop()
	logger.Info("Application stopped")
}

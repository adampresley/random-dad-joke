/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package configuration

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func NewConfig(serverVersion string) *viper.Viper {
	var err error

	result := viper.New()
	result.Set("server.version", serverVersion)

	result.SetDefault("server.host", "0.0.0.0:8080")
	result.SetDefault("server.loglevel", true)

	pflag.String("server.host", "0.0.0.0:8080", "host and IP to bind to")
	pflag.Parse()
	result.BindPFlags(pflag.CommandLine)

	result.BindEnv("PORT")

	result.SetConfigName("config")
	result.SetConfigType("yaml")
	result.AddConfigPath("/opt/starter")
	result.AddConfigPath("C:\\starter")
	result.AddConfigPath("$HOME/.starter")
	result.AddConfigPath(".")

	if err = result.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading configuration file")
			panic(err)
		}
	}

	return result
}

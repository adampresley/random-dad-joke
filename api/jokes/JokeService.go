/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/adampresley/random-dad-joke/api/httpclient"
)

// JokeServicer describes methods for getting dad jokes
type JokeServicer interface {
	GetRandomJoke() (*Joke, error)
}

// JokeServiceConfig is a configuration object for making a new JokeService
type JokeServiceConfig struct {
	Config     *viper.Viper
	HttpClient httpclient.HttpClient
}

// JokeService provides methods for getting dad jokes
type JokeService struct {
	Config     *viper.Viper
	HttpClient httpclient.HttpClient
}

// NewJokeService creates a new JokeService
func NewJokeService(config JokeServiceConfig) *JokeService {
	return &JokeService{
		Config:     config.Config,
		HttpClient: config.HttpClient,
	}
}

// GetRandomJoke retrieves a random dad joke
func (s *JokeService) GetRandomJoke() (*Joke, error) {
	var err error
	var request *http.Request
	var response *http.Response
	var result *Joke

	if request, err = http.NewRequest("GET", "https://icanhazdadjoke.com", nil); err != nil {
		return result, fmt.Errorf("error creating request in GetRandomJoke: %w", err)
	}

	request.Header.Add("Accept", "application/json")

	if response, err = s.HttpClient.Do(request); err != nil {
		return result, fmt.Errorf("error getting joke in GetRandomJoke: %w", err)
	}

	decoder := json.NewDecoder(response.Body)

	if err = decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("error decoding body in GetRandomJoke: %w", err)
	}

	return result, nil
}

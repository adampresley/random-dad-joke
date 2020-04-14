/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/adampresley/random-dad-joke/api/jokes"
)

func TestJokeController_GetRandomJoke(t *testing.T) {
	logger := logrus.New().WithField("who", "JokeControllerTests")

	type fields struct {
		Config      *viper.Viper
		JokeService jokes.JokeServicer
		Logger      *logrus.Entry
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name       string
		fields     fields
		wantStatus int
	}{
		{
			name: "Returns a joke",
			fields: fields{
				Config: viper.New(),
				JokeService: &jokes.MockJokeService{
					GetRandomJokeFunc: func() (*jokes.Joke, error) {
						return &jokes.Joke{
							ID:     "abc",
							Joke:   "This is a funny",
							Status: 200,
						}, nil
					},
				},
				Logger: logger,
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "Returns an error when JokeService fails",
			fields: fields{
				Config: viper.New(),
				JokeService: &jokes.MockJokeService{
					GetRandomJokeFunc: func() (*jokes.Joke, error) {
						return &jokes.Joke{}, fmt.Errorf("Error getting joke")
					},
				},
				Logger: logger,
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		e := echo.New()

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)

		t.Run(tt.name, func(t *testing.T) {
			c := &jokes.JokeController{
				Config:      tt.fields.Config,
				JokeService: tt.fields.JokeService,
				Logger:      tt.fields.Logger,
			}

			if err := c.GetRandomJoke(ctx); err != nil {
				t.Errorf("GetRandomJoke() error = %v", err)
			}

			if tt.wantStatus != recorder.Code {
				t.Errorf("GetRandomJoke() expected %d, got %d", tt.wantStatus, recorder.Code)
			}
		})
	}
}

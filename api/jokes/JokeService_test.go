/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/spf13/viper"

	"github.com/adampresley/random-dad-joke/api/httpclient"
	"github.com/adampresley/random-dad-joke/api/jokes"
)

func TestJokeService_GetRandomJoke(t *testing.T) {
	config := viper.New()
	joke1 := []byte(`{"id": "bprz5wXSSvc", "joke": "Why did the scarecrow win an award? Because he was outstanding in his field.", "status": 200}`)

	type fields struct {
		Config     *viper.Viper
		HttpClient httpclient.HttpClient
	}

	tests := []struct {
		name    string
		fields  fields
		want    *jokes.Joke
		wantErr bool
	}{
		{
			name: "Returns a joke",
			fields: fields{
				Config: config,
				HttpClient: &httpclient.MockHttpClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(bytes.NewReader(joke1)),
						}, nil
					},
				},
			},
			want: &jokes.Joke{
				ID:     "bprz5wXSSvc",
				Joke:   "Why did the scarecrow win an award? Because he was outstanding in his field.",
				Status: 200,
			},
			wantErr: false,
		},
		{
			name: "Returns an error when the service fails",
			fields: fields{
				Config: config,
				HttpClient: &httpclient.MockHttpClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("Error getting joke")
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &jokes.JokeService{
				Config:     tt.fields.Config,
				HttpClient: tt.fields.HttpClient,
			}

			got, err := s.GetRandomJoke()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandomJoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomJoke() got = %v, want %v", got, tt.want)
			}
		})
	}
}

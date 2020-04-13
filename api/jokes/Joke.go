/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type JokeCollection []*Joke

/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package jokes

type MockJokeService struct {
	GetRandomJokeFunc func() (*Joke, error)
}

func (m *MockJokeService) GetRandomJoke() (*Joke, error) {
	return m.GetRandomJokeFunc()
}

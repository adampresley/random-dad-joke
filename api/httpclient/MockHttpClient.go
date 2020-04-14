/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package httpclient

import (
	"net/http"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

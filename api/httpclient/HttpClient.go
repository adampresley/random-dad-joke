/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

package httpclient

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

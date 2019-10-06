package config

import (
	"net/http"
	"net/url"
)

type HTTPRequest interface {
	GetQueryParams() url.Values
	GetMethod() string
	GetPath() string
	GetBody() []byte
	GetHeaderParams() http.Header
}

type httpRequest struct {
	headers     http.Header
	body        []byte
	path        string
	method      string
	queryParams url.Values
}

func (req *httpRequest) GetQueryParams() url.Values {
	return req.queryParams
}

func (req *httpRequest) GetMethod() string {
	return req.method
}

func (req *httpRequest) GetPath() string {
	return req.path
}

func (req *httpRequest) GetBody() []byte {
	return req.body
}

func (req *httpRequest) SetHeaderParam(key string, values ...string) error {
	req.headers.Del(key)
	for _, val := range values {
		req.headers.Add(key, val)
	}
	return nil
}

func (req *httpRequest) GetHeaderParams() http.Header {
	return req.headers
}

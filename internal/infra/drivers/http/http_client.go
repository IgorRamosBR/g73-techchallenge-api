package http

import (
	"bytes"
	httpClient "net/http"
)

type HttpClient interface {
	DoPost(url string, body []byte) (*httpClient.Response, error)
	DoGet(url string) (*httpClient.Response, error)
}

type client struct {
	client *httpClient.Client
}

func NewHttpClient() HttpClient {
	return client{
		client: &httpClient.Client{},
	}
}

func (c client) DoPost(url string, body []byte) (*httpClient.Response, error) {
	return httpClient.Post(url, "application/json", bytes.NewBuffer(body))
}

func (c client) DoGet(url string) (*httpClient.Response, error) {
	return httpClient.Get(url)
}

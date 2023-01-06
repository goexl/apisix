package apisix

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

// Client 客户端
type Client struct {
	endpoint string
	apiKey   string

	http   *resty.Client
	logger simaqian.Logger
}

func newClient(endpoint string, apiKey string, http *resty.Client, logger simaqian.Logger) *Client {
	return &Client{
		endpoint: endpoint,
		apiKey:   apiKey,

		http:   http,
		logger: logger,
	}
}

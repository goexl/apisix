package apisix

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = New

// Creator 客户端创建者
type Creator struct {
	http   *resty.Client
	logger simaqian.Logger
}

func New() *Creator {
	return &Creator{
		http:   resty.New(),
		logger: simaqian.Default(),
	}
}

func (c *Creator) HttpClient(client *resty.Client) *Creator {
	c.http = client

	return c
}

func (c *Creator) Logger(logger simaqian.Logger) *Creator {
	c.logger = logger

	return c
}

func (c *Creator) Create(endpoint string, apiKey string) *Client {
	return newClient(endpoint, apiKey, c.http, c.logger)
}

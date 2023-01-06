package apisix

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = New

type builder struct {
	endpoint string
	apiKey   string
	http     *resty.Client
	logger   simaqian.Logger
}

func New(endpoint string, token string) *builder {
	return &builder{
		endpoint: endpoint,
		apiKey:   token,

		http:   resty.New(),
		logger: simaqian.Default(),
	}
}

func (b *builder) HttpClient(client *resty.Client) *builder {
	b.http = client

	return b
}

func (b *builder) Logger(logger simaqian.Logger) *builder {
	b.logger = logger

	return b
}

func (b *builder) Build() *Client {
	return newClient(b.endpoint, b.apiKey, b.http, b.logger)
}

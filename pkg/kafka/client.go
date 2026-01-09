package kafka

import "github.com/samber/do/v2"

type Client struct {
}

func NewKafka(injector do.Injector) (*Client, error) {
	return &Client{}, nil
}

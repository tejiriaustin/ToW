package services

import (
	"errors"
	"github.com/tejiriaustin/ToW/payment"
)

type (
	Container struct {
		AccountsService  AccountServiceInterface
		AdminService     AdminServiceInterface
		PaymentProcessor payment.ProcessorProvider
		TokenProvider    TokenProvider
	}
	Pager struct {
		Page    int64
		PerPage int64
	}
	Options func(c *Container) error
)

func New(opts ...Options) *Container {
	c := &Container{}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return c
		}
	}
	return c
}

func WithPaymentProcessor(provider payment.ProcessorProvider) Options {
	return func(c *Container) error {
		if provider == nil {
			return errors.New("provider is nil")
		}
		c.PaymentProcessor = provider
		return nil
	}
}

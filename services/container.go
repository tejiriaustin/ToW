package services

import (
	"errors"
	"github.com/tejiriaustin/ToW/payment"
	"reflect"
)

type (
	Container struct {
		AccountsService  AccountServiceInterface
		WalletService    WalletServiceInterface
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
	c := &Container{
		AccountsService: NewAccountService(),
		WalletService:   NewWalletService(),
		AdminService:    NewAdminService(),
		TokenProvider:   NewTokenProvider(),
	}

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

func Contains[T any](s T, c []T) bool {
	for _, t := range c {
		if reflect.DeepEqual(t, s) {
			return true
		}
	}
	return false
}

func StringContains(s string, t []string) bool {
	return Contains[string](s, t)
}

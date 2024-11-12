package container

import (
	"github.com/murasame29/go-httpserver-template/internal/router"
	"go.uber.org/dig"
)

var container *dig.Container

type provideArg struct {
	constructor any
	opts        []dig.ProvideOption
}

// NewContainer は、digを用いて依存注入を行う
func NewContainer() error {
	var noOpts []dig.ProvideOption
	container = dig.New()

	args := []provideArg{
		{constructor: router.NewEcho, opts: noOpts},
		// {constructor: db.NewRepository, opts: as[dai.DataAccessInterfce]()},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func as[T any]() []dig.ProvideOption {
	return []dig.ProvideOption{dig.As(new(T))}
}

// Invoke は、 *dig.ContainerのInvokeをwrapしてる関数
func Invoke[T any](opts ...dig.InvokeOption) (T, error) {
	var r T
	if err := container.Invoke(func(t T) error {
		r = t
		return nil
	}, opts...); err != nil {
		return r, err
	}
	return r, nil
}

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
	container = dig.New()

	args := []provideArg{
		{constructor: router.NewEcho, opts: []dig.ProvideOption{}},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
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

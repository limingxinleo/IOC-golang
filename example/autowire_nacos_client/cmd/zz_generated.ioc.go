//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package main

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &App{},
		Factory: func() interface{} {
			return &App{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*App)
			return param.Init(impl)
		},
	})
}

type paramInterface interface {
	Init(impl *App) (*App, error)
}

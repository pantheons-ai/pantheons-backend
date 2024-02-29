package rest

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf


}

type ServiceContext struct {
	Config Config
}

func NewServiceContext(c Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
	}
}
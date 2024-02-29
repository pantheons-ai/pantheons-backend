package rest

import (
	"log"

	"github.com/pantheons-ai/pantheon/pkg/ethc"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ChainConfig ethc.ClientConfig
}

type ServiceContext struct {
	Config Config
	ChainClient *ethc.Client
}

func NewServiceContext(c Config) *ServiceContext {
	chainClient, err := ethc.NewChainClient(c.ChainConfig)
	if err != nil {
		log.Fatalf("new client of chain error: %v", err)
	}
	
	return &ServiceContext{
		Config:  c,
		ChainClient: chainClient,
	}
}
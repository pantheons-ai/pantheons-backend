package rest

import (
	"log"

	"github.com/pantheons-ai/pantheon/pkg/ethc"
	"github.com/pantheons-ai/pantheon/pkg/ipfs"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ChainConfig ethc.Config
	IPFSConfig  ipfs.Config
}

type ServiceContext struct {
	Config      Config
	ChainClient *ethc.Client
	IpfsClient  *ipfs.Client
}

func NewServiceContext(c Config) *ServiceContext {
	chainClient, err := ethc.NewClient(c.ChainConfig)
	if err != nil {
		log.Fatalf("new client for chain error: %v", err)
	}
	ipfsClient, err := ipfs.NewClient(c.IPFSConfig)
	if err != nil {
		log.Fatalf("new client for ipfs error: %v", err)
	}
	return &ServiceContext{
		Config:      c,
		IpfsClient:  ipfsClient,
		ChainClient: chainClient,
	}
}

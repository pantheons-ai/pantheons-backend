package ipfs

import (
	"fmt"

	"github.com/ipfs/kubo/client/rpc"
	// path "github.com/ipfs/boxo/coreiface/path"
)

type Config struct {
	GatewayURL string // the url of ipfs gateway
}

type Client struct {
	config Config
	client *rpc.HttpApi
	folders map[string]string
}

func NewClient(c Config) (*Client, error) {
	client, err := rpc.NewPathApi(c.GatewayURL)
	if err != nil {
		return nil, fmt.Errorf("new ipfs client error: %v", err)
	}
	return &Client{config: c, client:  client}, nil
}

func (c *Client) CreateFolderForIP(ipAddr string) error {

	c.client.Name().Publish()


	return nil
}
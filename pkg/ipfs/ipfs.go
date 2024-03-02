package ipfs

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/ipfs/boxo/path"
	cidx "github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
)

type Config struct {
	GatewayURL string // the url of ipfs gateway
}

type Client struct {
	config  Config
	client  *rpc.HttpApi
	folders map[string]string
}

func NewClient(c Config) (*Client, error) {
	client, err := rpc.NewPathApi(c.GatewayURL)
	if err != nil {
		return nil, fmt.Errorf("new ipfs client error: %v", err)
	}
	return &Client{config: c, client: client}, nil
}

func (c *Client) PublishIP(ctx context.Context, ipAddr string) (string, error) {
	folderPath := filepath.Join(c.config.GatewayURL, ipAddr)
	ipPath, err := path.NewPath(folderPath)
	if err != nil {
		return "", fmt.Errorf("new path of ip %s error: %v", ipAddr, err)
	}
	ipnsName, err := c.client.Name().Publish(ctx, ipPath)
	if err != nil {
		return "", fmt.Errorf("publish ip %s to ipns error: %s", ipAddr, err)
	}
	return ipnsName.Cid().String(), nil
}

func (c *Client) GetContent(ctx context.Context, cidStr string) ([]byte, error) {
	cid, err := cidx.Decode(cidStr)
	if err != nil {
		return nil, fmt.Errorf("cid %s is not valid: %v", cidStr, err)
	}
	reader, err := c.client.Object().Get(ctx, path.FromCid(cid))
	if err != nil {
		return nil, fmt.Errorf("get object of cid %s error %v", cidStr, err)
	}
	return reader.RawData(), nil
}

func (c *Client) PutContent(ctx context.Context, filename string, content io.Reader) (string, error) {

	c.client.Object().Put(ctx, content)
	return "", nil
}

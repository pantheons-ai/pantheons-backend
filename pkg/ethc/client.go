package ethc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pantheons-ai/pantheon/pkg/contract"
)

type ClientConfig struct {
	RPCURL          string `json:"rpc_url"`
	PrivateKey      string `json:"private_key"`
}

type Client struct {
	config ClientConfig
	client *ethclient.Client
	privateKey *ecdsa.PrivateKey
	pantheons map[string]*contract.Pantheon
}

func NewChainClient(c ClientConfig) (*Client, error) {
	client, err := ethclient.Dial(c.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}
	// load private key
	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to load private key: %v", err)
	}
	return &Client{config: c, client: client, privateKey: privateKey}, nil
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.client.NetworkID(ctx)
}

func (c *Client) NewKeyedTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := c.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id error: %v", err)
	}
	return bind.NewKeyedTransactorWithChainID(c.privateKey, chainID)
}

func (c *Client) NewPantheon(address string) error {
	if _, ok := c.pantheons[address]; ok {
		return nil
	}

	contractAddress := common.HexToAddress(address)
	instance, err := contract.NewPantheon(contractAddress, c.client)
	if err != nil {
		return fmt.Errorf("Failed to create a instence of Pantheon contract: %v", err)
	}
	c.pantheons[address] = instance
	return nil
}
package ethc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pantheons-ai/pantheon/pkg/contract"
)

type Config struct {
	RPCURL     string `json:"rpc_url"`
	PrivateKey string `json:"private_key"`
}

type Client struct {
	config     Config
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	pantheons  map[string]*contract.Pantheon
	founders   map[string]common.Address
}

func NewClient(c Config) (*Client, error) {
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

func (c *Client) NewPantheon(ipaddr string) error {
	if _, ok := c.pantheons[ipaddr]; ok {
		return nil
	}

	contractAddress := common.HexToAddress(ipaddr)
	instance, err := contract.NewPantheon(contractAddress, c.client)
	if err != nil {
		return fmt.Errorf("Failed to create a instence of Pantheon contract: %v", err)
	}
	c.pantheons[ipaddr] = instance
	return nil
}

func (c *Client) AddUserToWhiteList(ipaddr, useraddr, cidStr string) error {
	contract, ok := c.pantheons[ipaddr]
	if !ok {
		return fmt.Errorf("ip %s not exists in chain", ipaddr)
	}
	founder, ok := c.founders[ipaddr]
	if !ok {
		return fmt.Errorf("not find founder for ip %s", ipaddr)
	}

	opt := &bind.TransactOpts{From: founder}

	userAddr := common.HexToAddress(useraddr)
	_, err := contract.AddToWhitelist(opt, userAddr)
	if err != nil {
		return fmt.Errorf("add user %s to ip %s error: %v", useraddr, ipaddr)
	}
	return nil
}

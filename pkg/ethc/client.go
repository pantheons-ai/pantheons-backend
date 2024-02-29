package ethc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
)

type ClientConfig struct {
	RPCURL          string `json:"rpc_url"`
	PrivateKey      string `json:"private_key"`
	ContractAddress string `json:"contract_address"`
}

type ChainClient struct {

}

func NewChainClient(c ClientConfig) error {
	client, err := ethclient.Dial(c.RPCURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

}
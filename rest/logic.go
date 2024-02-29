package rest

import (
	"context"
	"fmt"
)

type Logic struct {
	ctx context.Context
	svcCtx *ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *ServiceContext) *Logic {


	return &Logic{}
}

func (l *Logic) CreateIPUniversal(req *CreateIPReq) (string, error) {
	// 1. create a instence of Pantheon contract on chain
	err := l.svcCtx.ChainClient.NewPantheon(req.ContractAddress)
	if err != nil {
		return "", fmt.Errorf("init IP universal error: %v", err)
	}

	// 2. create a folder to store all files of this IP
	

	

	return "", nil
}

func (l *Logic) AddUserToWhiteList() error {

	return nil
}
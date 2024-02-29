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
	if err := l.svcCtx.ChainClient.NewPantheon(req.Address); err != nil {
		return "", fmt.Errorf("init IP universal error: %v", err)
	}

	return "", nil
}
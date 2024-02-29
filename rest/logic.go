package rest

import "context"

type Logic struct {
	ctx context.Context
	svcCtx *ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *ServiceContext) *Logic {

	
	return &Logic{}
}
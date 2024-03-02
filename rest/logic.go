package rest

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type Logic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *ServiceContext
}

func NewLogic(ctx context.Context, svcCtx *ServiceContext) *Logic {
	return &Logic{ctx: ctx, svcCtx: svcCtx, logger: logx.WithContext(ctx)}
}

func (l *Logic) CreateIPUniversal(req *CreateIPReq) (*CreateIPResp, error) {
	// 1. create a instence of Pantheon contract on chain
	err := l.svcCtx.ChainClient.NewPantheon(req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("init IP universal error: %v", err)
	}

	// 2. create a folder to store all files of this IP
	folderCIDOfIP, err := l.svcCtx.IpfsClient.PublishIP(l.ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("create namespace for ip %s on ipfs error: %v", req.Name, err)
	}
	return &CreateIPResp{FolderCID: folderCIDOfIP}, nil
}

func (l *Logic) AddUserToWhiteList(req *AddUserReq) error {
	return l.svcCtx.ChainClient.AddUserToWhiteList(req.ContractAddress, req.UserAddress)
}

func (l *Logic) PushlishContent(req *PublishContentReq, form *multipart.Form) (*PublishContentResp, error) {
	// 1. upload file to the folder of ip to ipfs
	content := strings.NewReader(strings.Join(form.Value[req.FileName], "/n"))
	cid, err := l.svcCtx.IpfsClient.PutContent(l.ctx, req.FileName, content)
	if err != nil {
		return nil, fmt.Errorf("push file %s to ipfs error: %v", req.FileName, err)
	}
	// 2. add cid of file and the user and ip address info on chain
	err = l.svcCtx.ChainClient.AddCID(req.ContractAddress, req.UserAddress, cid)
	if err != nil {
		return nil, fmt.Errorf("fail to pushlish file %s error: %s", req.FileName, err)
	}
	return &PublishContentResp{Cid: cid}, nil
}
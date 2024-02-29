package rest

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

const VersionPrefix = "/v1"

func RegisterHandlers(server *rest.Server, svcCtx *ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method: http.MethodPost,
			Path: "/ip-universal",
			Handler: CreateIPHandle(svcCtx),
		},
		{
			Method: http.MethodPost,
			Path: "/create-user",
			Handler: CreateUserHandle(svcCtx),
		},

		{
			Method: http.MethodPost,
			Path: "/upload-files",
			Handler: UploadFilesHandle(svcCtx),
		},
		
	}, rest.WithPrefix(VersionPrefix),)
}
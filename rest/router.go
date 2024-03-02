package rest

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

const VersionPrefix = "/v1"

func RegisterHandlers(server *rest.Server, svcCtx *ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/ip-universal",
			Handler: CreateIPHandle(svcCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/add-user",
			Handler: AddUserHandle(svcCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/push-content",
			Handler: PublishContentHandle(svcCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/ip-contents",
			Handler: GetIPContentsHandle(svcCtx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/user-contents",
			Handler: GetUserContentsHandler(svcCtx),
		},
	}, rest.WithPrefix(VersionPrefix))
}

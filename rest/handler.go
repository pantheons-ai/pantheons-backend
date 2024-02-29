package rest

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateIPHandle(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req CreateIPReq
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}

		logic := NewLogic(ctx, svcCtx)
		resp, err := logic.CreateIPUniversal(&req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}
		httpx.OkJsonCtx(ctx, w, resp)
	}
}

func AddUserHandle(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req CreateIPReq
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}

		logic := NewLogic(ctx, svcCtx)
		resp, err := logic.CreateIPUniversal(&req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}
		httpx.OkJsonCtx(ctx, w, resp)
	}
}

func UploadFilesHandle(svc *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
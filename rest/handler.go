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
		var req AddUserReq
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}

		logic := NewLogic(ctx, svcCtx)
		err = logic.AddUserToWhiteList(&req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}
		httpx.Ok(w)
	}
}

func PublishContentHandle(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req PublishContentReq
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}

		logic := NewLogic(ctx, svcCtx)
		resp, err := logic.PushlishContent(&req, r.MultipartForm)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
			return
		}
		httpx.OkJsonCtx(ctx, w, resp)
	}
}

func GetIPContentsHandle(svcCtx *ServiceContext) http.HandlerFunc {
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

func GetUserContentsHandler(svcCtx *ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

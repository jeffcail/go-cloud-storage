package handler

import (
	"net/http"

	"github.com/jeffcail/cloud-storage/server/internal/logic"
	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareInput
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadPrepare(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

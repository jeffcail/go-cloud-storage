// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/jeffcail/cloud-storage/server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send",
				Handler: MailCodeSendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/file/detail",
				Handler: ShareFileDetailHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/file/list",
					Handler: UserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/name/update",
					Handler: UserFileNameUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/floder/create",
					Handler: UserFloderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file/delete",
					Handler: UserFileDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/file/move",
					Handler: UserFileMoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/create",
					Handler: ShareCrateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/save",
					Handler: ShareSaveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/refresh/authorization",
					Handler: RefreshAuthorizationHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/prepare",
					Handler: FileUploadPrepareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/chunk",
					Handler: FileUploadChunkHandler(serverCtx),
				},
			}...,
		),
	)
}

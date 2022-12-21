package logic

import (
	"context"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveInput, userIdentity string) (resp *types.UserRepositorySaveOut, err error) {
	ur := &models.StorageUserRepository{
		Ext:                req.Ext,
		Identity:           utils.GenerateUuid(),
		Name:               req.Name,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		UserIdentity:       userIdentity,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp = &types.UserRepositorySaveOut{Message: "success"}
	return
}

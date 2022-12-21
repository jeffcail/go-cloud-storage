package logic

import (
	"context"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareCrateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareCrateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareCrateLogic {
	return &ShareCrateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareCrateLogic) ShareCrate(req *types.ShareCrateInput, userIdentity string) (resp *types.ShareCrateOut, err error) {
	uuid := utils.GenerateUuid()
	data := &models.StorageShare{
		ExpiredTime:        req.ExpiredTime,
		RepositoryIdentity: req.RepositoryIdentity,
		UserIdentity:       userIdentity,
		Identity:           uuid,
		ClickNum:           0,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareCrateOut{Identity: uuid}
	return
}

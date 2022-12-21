package logic

import (
	"context"
	"errors"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareSaveLogic {
	return &ShareSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareSaveLogic) ShareSave(req *types.ShareSaveInput, userIdentity string) (resp *types.ShareSaveOut, err error) {
	rp := new(models.StorageRepositoryPool)
	has, err := l.svcCtx.Engine.
		Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("资源不存在")
	}

	ur := &models.StorageUserRepository{
		Ext:                rp.Ext,
		Identity:           utils.GenerateUuid(),
		Name:               rp.Name,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		UserIdentity:       userIdentity,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	resp = &types.ShareSaveOut{Identity: ur.Identity}
	return
}

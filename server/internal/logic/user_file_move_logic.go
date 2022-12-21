package logic

import (
	"context"
	"errors"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveInput, userIdentity string) (resp *types.UserFileMoveOut, err error) {
	parentData := new(models.StorageUserRepository)
	has, err := l.svcCtx.Engine.
		Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).
		Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}
	_, err = l.svcCtx.Engine.
		Where("identity = ?", req.Identity).
		Update(models.StorageUserRepository{ParentId: parentData.Id})
	if err != nil {
		return nil, err
	}
	resp = &types.UserFileMoveOut{Message: "success"}

	return
}

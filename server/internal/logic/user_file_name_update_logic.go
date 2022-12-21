package logic

import (
	"context"
	"errors"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateInput, userIdentity string) (resp *types.UserFileNameUpdateOut, err error) {
	// 检测当前名称在该层级下是否存在
	count, err := l.svcCtx.Engine.Where("name = ? AND parent_id = (SELECT parent_id FROM storage_user_repository ur WHERE "+
		"ur.identity = ?)", req.Name, req.Identity).Count(new(models.StorageUserRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("文件名称已存在")
	}

	data := &models.StorageUserRepository{Name: req.Name}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return nil, err
	}
	resp = &types.UserFileNameUpdateOut{Message: "success"}
	return
}

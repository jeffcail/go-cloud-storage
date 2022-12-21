package logic

import (
	"context"
	"errors"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailInput) (resp *types.UserDetailOut, err error) {
	u := new(models.StorageUser)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(u)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}
	resp = &types.UserDetailOut{
		Name:  u.Name,
		Email: u.Email,
	}
	return
}

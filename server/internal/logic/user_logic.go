package logic

import (
	"context"
	"errors"

	"github.com/jeffcail/cloud-storage/server/core"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginInput) (resp *types.LoginOut, err error) {
	user := new(models.StorageUser)
	pass := utils.Sha256(req.Password)
	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, pass).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name, int64(core.CodeExpire))
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(user.Id, user.Identity, user.Name, int64(core.TokenExpire))
	if err != nil {
		return nil, err
	}

	resp = &types.LoginOut{
		Id:           user.Id,
		Name:         user.Name,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}

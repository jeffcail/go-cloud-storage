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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterInput) (resp *types.UserRegisterOut, err error) {
	u := new(models.StorageUser)
	res, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("验证码失效")
	}
	if res != req.Code {
		return nil, errors.New("验证码错误")
	}
	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(u)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}
	u.Name = req.Name
	u.Email = req.Email
	u.Identity = utils.GenerateUuid()
	u.Password = utils.Sha256(req.Password)
	i, err := l.svcCtx.Engine.Insert(u)
	if err != nil {
		return nil, err
	}
	if i != 1 {
		return nil, errors.New("注册失败")
	}
	l.svcCtx.RDB.Del(l.ctx, req.Email)

	resp = &types.UserRegisterOut{Message: "注册成功,请去登陆"}

	return
}

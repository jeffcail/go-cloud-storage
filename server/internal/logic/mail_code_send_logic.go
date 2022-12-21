package logic

import (
	"context"
	"errors"
	"time"

	"github.com/jeffcail/cloud-storage/server/core"

	"github.com/jeffcail/cloud-storage/server/models"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendInput) (resp *types.MailCodeSendOut, err error) {
	u := new(models.StorageUser)
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(u)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("该邮箱已被注册")
	}

	// 随机code
	code := utils.RandCode()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(core.CodeExpire))

	err = utils.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	resp = &types.MailCodeSendOut{Code: code}

	return
}

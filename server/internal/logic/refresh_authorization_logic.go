package logic

import (
	"context"

	"github.com/jeffcail/cloud-storage/server/core"
	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationInput, authorization string) (resp *types.RefreshAuthorizationOut, err error) {
	userClaim, err := utils.ParseToken(authorization)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(userClaim.ID, userClaim.Name, userClaim.Identity, int64(core.CodeExpire))
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(userClaim.ID, userClaim.Name, userClaim.Identity, int64(core.TokenExpire))
	if err != nil {
		return nil, err
	}

	resp = &types.RefreshAuthorizationOut{
		Id:           userClaim.ID,
		Name:         userClaim.Name,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}

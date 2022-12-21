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

type UserFloderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFloderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFloderCreateLogic {
	return &UserFloderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFloderCreateLogic) UserFloderCreate(req *types.UserFloderCreateInput, userIdentity string) (resp *types.UserFloderCreateOut, err error) {
	// 检测当前名称在该层级下是否存在
	count, err := l.svcCtx.Engine.Where("`name` = ? AND `parent_id` = ?)", req.Name, req.ParentId).Count(new(models.StorageUserRepository))

	//if err != nil {
	//	return nil, err
	//}
	if count > 0 {
		return nil, errors.New("文件名称已存在")
	}

	data := &models.StorageUserRepository{
		Identity:     utils.GenerateUuid(),
		Name:         req.Name,
		ParentId:     req.ParentId,
		UserIdentity: userIdentity,
	}
	if _, err = l.svcCtx.Engine.Insert(data); err != nil {
		return nil, err
	}
	resp = &types.UserFloderCreateOut{Identity: userIdentity}

	return
}

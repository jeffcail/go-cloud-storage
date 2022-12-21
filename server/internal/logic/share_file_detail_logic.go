package logic

import (
	"context"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileDetailLogic {
	return &ShareFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileDetailLogic) ShareFileDetail(req *types.ShareFileDetailInput) (resp *types.ShareFileDetailOut, err error) {
	_, err = l.svcCtx.Engine.Exec("UPDATE storage_share SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		return nil, err
	}
	resp = new(types.ShareFileDetailOut)
	_, err = l.svcCtx.Engine.Table("storage_share").
		Select("storage_share.repository_identity, storage_repository_pool.name, storage_repository_pool.ext, storage_repository_pool.path, storage_repository_pool.size").
		Join("LEFT", "storage_repository_pool", "storage_share.repository_identity = storage_repository_pool.identity").
		Where("storage_share.identity = ?", req.Identity).
		Get(resp)
	return
}

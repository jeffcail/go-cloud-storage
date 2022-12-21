package logic

import (
	"context"
	"time"

	"github.com/jeffcail/cloud-storage/server/core"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListInput, userIdentity string) (resp *types.UserFileListOut, err error) {
	uf := make([]*types.UserFile, 0)
	var cnt int64
	resp = new(types.UserFileListOut)
	size := req.PageSize
	if size == 0 {
		size = core.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	cnt, err = l.svcCtx.Engine.Table("storage_user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("storage_user_repository.id, storage_user_repository.identity, storage_user_repository.repository_identity, "+
			"storage_user_repository.ext, storage_user_repository.name, storage_repository_pool.path, storage_repository_pool.size").
		Join("LEFT", "storage_repository_pool", "storage_user_repository.repository_identity "+
			"= storage_repository_pool.identity").
		Where("storage_user_repository.deleted_at = ? OR storage_user_repository.deleted_at IS NULL", time.Time{}.Format("2006-01-02 15:04:05")).
		Limit(size, offset).
		FindAndCount(&uf)
	if err != nil {
		return nil, err
	}

	resp.List = uf
	resp.Count = cnt
	return
}

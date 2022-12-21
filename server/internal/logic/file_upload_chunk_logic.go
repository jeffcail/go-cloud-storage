package logic

import (
	"context"

	"github.com/jeffcail/cloud-storage/server/models"
	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/internal/svc"
	"github.com/jeffcail/cloud-storage/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkLogic {
	return &FileUploadChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkLogic) FileUploadChunk(req *types.FileUploadChunkInput) (resp *types.FileUploadChunkOut, err error) {
	rp := &models.StorageRepositoryPool{
		Ext:      req.Ext,
		Hash:     req.Hash,
		Identity: utils.GenerateUuid(),
		Name:     req.Name,
		Path:     req.Path,
		Size:     req.Size,
	}
	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadChunkOut)
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}

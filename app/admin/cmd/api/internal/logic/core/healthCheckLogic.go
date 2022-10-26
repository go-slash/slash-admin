package core

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"slash-admin/app/admin/cmd/api/internal/svc"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() error {
	// todo: add your logic here and delete this line

	return nil
}

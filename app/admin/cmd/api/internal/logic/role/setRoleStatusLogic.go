package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"
	"slash-admin/pkg/message"
	pType "slash-admin/pkg/types"

	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleStatusLogic {
	return &SetRoleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoleStatusLogic) SetRoleStatus(req *types.SetStatusReq) (resp *types.SimpleMsgResp, err error) {
	if req.Status != 0 {
		req.Status = 1
	}
	_, err = l.svcCtx.EntClient.SysRole.UpdateOneID(req.ID).SetStatus(pType.Status(req.Status)).Save(l.ctx)

	if err != nil {
		logx.Errorw("update role status failed, please check the role id", logx.Field("roleId", req.ID))
		return nil, errorx.NewApiInternalServerError(message.DatabaseError)
	}
	logx.Infow("update role status successfully", logx.Field("roleID", req.ID), logx.Field("status", req.Status))
	return &types.SimpleMsgResp{Msg: message.UpdateSuccess}, nil
}

package notifymanagelogic

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/internal/repo/relationDB"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"

	"gitee.com/i-Things/core/service/syssvr/internal/svc"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyChannelCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyChannelCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyChannelCreateLogic {
	return &NotifyChannelCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyChannelCreateLogic) NotifyChannelCreate(in *sys.NotifyChannel) (*sys.WithID, error) {
	po := utils.Copy[relationDB.SysNotifyChannel](in)
	po.ID = 0
	if !utils.SliceIn(po.Type, def.NotifyTypeSms, def.NotifyTypeEmail, def.NotifyTypeDingTalk,
		def.NotifyTypeDingWebhook, def.NotifyTypeWx, def.NotifyTypeMessage, def.NotifyTypeWxEWebhook) {
		return nil, errors.Parameter.AddMsg("type not support")
	}
	err := relationDB.NewNotifyChannelRepo(l.ctx).Insert(l.ctx, po)
	return &sys.WithID{Id: po.ID}, err
}
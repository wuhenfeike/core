package self

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/utils"

	"gitee.com/i-Things/core/service/apisvr/internal/svc"
	"gitee.com/i-Things/core/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileIndexLogic {
	return &ProfileIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileIndexLogic) ProfileIndex(req *types.UserProfileIndexReq) (resp *types.UserProfileIndexResp, err error) {
	ret, err := l.svcCtx.UserRpc.UserProfileIndex(l.ctx, utils.Copy[sys.UserProfileIndexReq](req))

	return utils.Copy[types.UserProfileIndexResp](ret), nil
}
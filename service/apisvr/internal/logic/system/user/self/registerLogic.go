package self

import (
	"context"
	"gitee.com/unitedrhino/core/service/apisvr/internal/logic/system/user"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"

	"gitee.com/unitedrhino/core/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/core/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) error {
	_, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &sys.UserRegisterReq{
		RegType:  req.RegType,
		Account:  req.Account,
		Code:     req.Code,
		CodeID:   req.CodeID,
		Password: req.Password,
		Expand:   req.Expand,
		Info:     user.UserInfoToRpc(req.Info),
	})
	return err
}

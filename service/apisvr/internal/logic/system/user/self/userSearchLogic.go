package self

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"

	"gitee.com/i-Things/core/service/apisvr/internal/svc"
	"gitee.com/i-Things/core/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSearchLogic {
	return &UserSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSearchLogic) UserSearch(req *types.UserSearchReq) (resp *types.UserSearchResp, err error) {
	info, err := l.svcCtx.UserRpc.UserInfoIndex(l.ctx, &sys.UserInfoIndexReq{
		Page: &sys.PageInfo{
			Page: 1,
			Size: 1,
		},
		Account: req.Account,
	})
	if err != nil {
		return nil, err
	}
	if len(info.List) == 0 {
		return nil, errors.NotFind
	}
	return utils.Copy[types.UserSearchResp](info.List[0]), nil
}
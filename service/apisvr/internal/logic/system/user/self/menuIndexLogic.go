package self

import (
	"context"
	"gitee.com/unitedrhino/core/service/apisvr/internal/logic/system"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/share/ctxs"

	"gitee.com/unitedrhino/core/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/core/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuIndexLogic {
	return &MenuIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuIndexLogic) MenuIndex(req *types.UserMenuIndexReq) (resp *types.TenantAppMenuIndexResp, err error) {
	uc := ctxs.GetUserCtx(l.ctx)
	var menuIDs []int64
	if !uc.IsSuperAdmin {
		ids, err := l.svcCtx.RoleRpc.RoleMenuIndex(l.ctx, &sys.RoleMenuIndexReq{Ids: uc.RoleIDs, AppCode: uc.AppCode, ModuleCode: req.ModuleCode})
		if err != nil {
			return nil, err
		}
		menuIDs = ids.MenuIDs
		if len(menuIDs) == 0 {
			return nil, nil
		}
	}

	ret, err := l.svcCtx.TenantRpc.TenantAppMenuIndex(l.ctx, &sys.TenantAppMenuIndexReq{Code: uc.TenantCode, IsCommon: req.IsCommon, AppCode: uc.AppCode, ModuleCode: req.ModuleCode, MenuIDs: menuIDs, IsRetTree: true})
	if err != nil {
		return nil, err
	}
	return &types.TenantAppMenuIndexResp{List: system.ToTenantAppMenusApi(ret.List)}, nil
}

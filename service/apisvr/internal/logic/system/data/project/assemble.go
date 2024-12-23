package project

import (
	"context"
	"gitee.com/unitedrhino/core/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/core/service/apisvr/internal/types"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/share/utils"
)

func ToProjectApis(ctx context.Context, svcCtx *svc.ServiceContext, in []*sys.DataProject) (ret []*types.DataProject) {
	if in == nil {
		return
	}
	for _, v := range in {
		ui := &types.UserCore{}
		if svcCtx != nil {
			u, err := svcCtx.UserCache.GetData(ctx, v.TargetID)
			if err != nil {
				continue
			}
			ui = utils.Copy[types.UserCore](u)
		}
		ret = append(ret, &types.DataProject{ProjectID: v.ProjectID, AuthType: v.AuthType, TargetID: v.TargetID, UpdatedTime: v.UpdatedTime, User: ui})
	}
	return
}

func ToProjectPbs(in []*types.DataProject) (ret []*sys.DataProject) {
	if in == nil {
		return
	}
	for _, v := range in {
		ret = append(ret, &sys.DataProject{ProjectID: v.ProjectID, AuthType: v.AuthType})
	}
	return
}

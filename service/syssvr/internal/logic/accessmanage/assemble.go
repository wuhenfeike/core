package accessmanagelogic

import (
	"gitee.com/unitedrhino/core/service/syssvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/share/utils"
)

func ToApiInfoPo(in *sys.ApiInfo) *relationDB.SysApiInfo {
	return utils.Copy[relationDB.SysApiInfo](in)
}

func ToApiInfoPb(in *relationDB.SysApiInfo) *sys.ApiInfo {
	return utils.Copy[sys.ApiInfo](in)
}

func ToAccessPo(in *sys.AccessInfo) *relationDB.SysAccessInfo {
	return utils.Copy[relationDB.SysAccessInfo](in)
}

func ToAccessPb(in *relationDB.SysAccessInfo) *sys.AccessInfo {
	return utils.Copy[sys.AccessInfo](in)
}

// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: sys.proto

package tenantmanage

import (
	"context"

	"gitee.com/unitedrhino/core/service/syssvr/internal/svc"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessInfo                            = sys.AccessInfo
	AccessInfoIndexReq                    = sys.AccessInfoIndexReq
	AccessInfoIndexResp                   = sys.AccessInfoIndexResp
	AccessInfoMultiImportReq              = sys.AccessInfoMultiImportReq
	AccessInfoMultiImportResp             = sys.AccessInfoMultiImportResp
	ApiInfo                               = sys.ApiInfo
	ApiInfoIndexReq                       = sys.ApiInfoIndexReq
	ApiInfoIndexResp                      = sys.ApiInfoIndexResp
	AppInfo                               = sys.AppInfo
	AppInfoIndexReq                       = sys.AppInfoIndexReq
	AppInfoIndexResp                      = sys.AppInfoIndexResp
	AppModuleIndexReq                     = sys.AppModuleIndexReq
	AppModuleIndexResp                    = sys.AppModuleIndexResp
	AppModuleMultiUpdateReq               = sys.AppModuleMultiUpdateReq
	AreaInfo                              = sys.AreaInfo
	AreaInfoIndexReq                      = sys.AreaInfoIndexReq
	AreaInfoIndexResp                     = sys.AreaInfoIndexResp
	AreaInfoReadReq                       = sys.AreaInfoReadReq
	AreaProfile                           = sys.AreaProfile
	AreaProfileIndexReq                   = sys.AreaProfileIndexReq
	AreaProfileIndexResp                  = sys.AreaProfileIndexResp
	AreaProfileReadReq                    = sys.AreaProfileReadReq
	AreaWithID                            = sys.AreaWithID
	AuthApiInfo                           = sys.AuthApiInfo
	CompareInt64                          = sys.CompareInt64
	CompareString                         = sys.CompareString
	ConfigResp                            = sys.ConfigResp
	DataArea                              = sys.DataArea
	DataAreaIndexReq                      = sys.DataAreaIndexReq
	DataAreaIndexResp                     = sys.DataAreaIndexResp
	DataAreaMultiDeleteReq                = sys.DataAreaMultiDeleteReq
	DataAreaMultiUpdateReq                = sys.DataAreaMultiUpdateReq
	DataProject                           = sys.DataProject
	DataProjectDeleteReq                  = sys.DataProjectDeleteReq
	DataProjectIndexReq                   = sys.DataProjectIndexReq
	DataProjectIndexResp                  = sys.DataProjectIndexResp
	DataProjectMultiDeleteReq             = sys.DataProjectMultiDeleteReq
	DataProjectMultiSaveReq               = sys.DataProjectMultiSaveReq
	DataProjectSaveReq                    = sys.DataProjectSaveReq
	DateRange                             = sys.DateRange
	DeptInfo                              = sys.DeptInfo
	DeptInfoIndexReq                      = sys.DeptInfoIndexReq
	DeptInfoIndexResp                     = sys.DeptInfoIndexResp
	DeptInfoReadReq                       = sys.DeptInfoReadReq
	DeptRoleIndexReq                      = sys.DeptRoleIndexReq
	DeptRoleIndexResp                     = sys.DeptRoleIndexResp
	DeptRoleMultiSaveReq                  = sys.DeptRoleMultiSaveReq
	DeptSyncJob                           = sys.DeptSyncJob
	DeptSyncJobExecuteReq                 = sys.DeptSyncJobExecuteReq
	DeptSyncJobExecuteResp                = sys.DeptSyncJobExecuteResp
	DeptSyncJobIndexReq                   = sys.DeptSyncJobIndexReq
	DeptSyncJobIndexResp                  = sys.DeptSyncJobIndexResp
	DeptSyncJobReadReq                    = sys.DeptSyncJobReadReq
	DeptUser                              = sys.DeptUser
	DeptUserIndexReq                      = sys.DeptUserIndexReq
	DeptUserIndexResp                     = sys.DeptUserIndexResp
	DeptUserMultiCreateReq                = sys.DeptUserMultiCreateReq
	DeptUserMultiDeleteReq                = sys.DeptUserMultiDeleteReq
	DictDetail                            = sys.DictDetail
	DictDetailIndexReq                    = sys.DictDetailIndexReq
	DictDetailIndexResp                   = sys.DictDetailIndexResp
	DictDetailMultiCreateReq              = sys.DictDetailMultiCreateReq
	DictDetailReadReq                     = sys.DictDetailReadReq
	DictInfo                              = sys.DictInfo
	DictInfoIndexReq                      = sys.DictInfoIndexReq
	DictInfoIndexResp                     = sys.DictInfoIndexResp
	DictInfoReadReq                       = sys.DictInfoReadReq
	Empty                                 = sys.Empty
	IDList                                = sys.IDList
	JwtToken                              = sys.JwtToken
	LoginLogCreateReq                     = sys.LoginLogCreateReq
	LoginLogIndexReq                      = sys.LoginLogIndexReq
	LoginLogIndexResp                     = sys.LoginLogIndexResp
	LoginLogInfo                          = sys.LoginLogInfo
	Map                                   = sys.Map
	MenuInfo                              = sys.MenuInfo
	MenuInfoIndexReq                      = sys.MenuInfoIndexReq
	MenuInfoIndexResp                     = sys.MenuInfoIndexResp
	MenuMultiExportReq                    = sys.MenuMultiExportReq
	MenuMultiExportResp                   = sys.MenuMultiExportResp
	MenuMultiImportReq                    = sys.MenuMultiImportReq
	MenuMultiImportResp                   = sys.MenuMultiImportResp
	MessageInfo                           = sys.MessageInfo
	MessageInfoIndexReq                   = sys.MessageInfoIndexReq
	MessageInfoIndexResp                  = sys.MessageInfoIndexResp
	MessageInfoSendReq                    = sys.MessageInfoSendReq
	ModuleInfo                            = sys.ModuleInfo
	ModuleInfoIndexReq                    = sys.ModuleInfoIndexReq
	ModuleInfoIndexResp                   = sys.ModuleInfoIndexResp
	NotifyChannel                         = sys.NotifyChannel
	NotifyChannelIndexReq                 = sys.NotifyChannelIndexReq
	NotifyChannelIndexResp                = sys.NotifyChannelIndexResp
	NotifyConfig                          = sys.NotifyConfig
	NotifyConfigIndexReq                  = sys.NotifyConfigIndexReq
	NotifyConfigIndexResp                 = sys.NotifyConfigIndexResp
	NotifyConfigSendReq                   = sys.NotifyConfigSendReq
	NotifyConfigTemplate                  = sys.NotifyConfigTemplate
	NotifyConfigTemplateDeleteReq         = sys.NotifyConfigTemplateDeleteReq
	NotifyConfigTemplateIndexReq          = sys.NotifyConfigTemplateIndexReq
	NotifyConfigTemplateIndexResp         = sys.NotifyConfigTemplateIndexResp
	NotifyTemplate                        = sys.NotifyTemplate
	NotifyTemplateIndexReq                = sys.NotifyTemplateIndexReq
	NotifyTemplateIndexResp               = sys.NotifyTemplateIndexResp
	OpenAccess                            = sys.OpenAccess
	OpenAccessIndexReq                    = sys.OpenAccessIndexReq
	OpenAccessIndexResp                   = sys.OpenAccessIndexResp
	OperLogCreateReq                      = sys.OperLogCreateReq
	OperLogIndexReq                       = sys.OperLogIndexReq
	OperLogIndexResp                      = sys.OperLogIndexResp
	OperLogInfo                           = sys.OperLogInfo
	OpsFeedback                           = sys.OpsFeedback
	OpsFeedbackIndexReq                   = sys.OpsFeedbackIndexReq
	OpsFeedbackIndexResp                  = sys.OpsFeedbackIndexResp
	OpsWorkOrder                          = sys.OpsWorkOrder
	OpsWorkOrderIndexReq                  = sys.OpsWorkOrderIndexReq
	OpsWorkOrderIndexResp                 = sys.OpsWorkOrderIndexResp
	PageInfo                              = sys.PageInfo
	PageInfo_OrderBy                      = sys.PageInfo_OrderBy
	Point                                 = sys.Point
	ProjectAuth                           = sys.ProjectAuth
	ProjectInfo                           = sys.ProjectInfo
	ProjectInfoIndexReq                   = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp                  = sys.ProjectInfoIndexResp
	ProjectProfile                        = sys.ProjectProfile
	ProjectProfileIndexReq                = sys.ProjectProfileIndexReq
	ProjectProfileIndexResp               = sys.ProjectProfileIndexResp
	ProjectProfileReadReq                 = sys.ProjectProfileReadReq
	ProjectWithID                         = sys.ProjectWithID
	QRCodeReadReq                         = sys.QRCodeReadReq
	QRCodeReadResp                        = sys.QRCodeReadResp
	RoleAccessIndexReq                    = sys.RoleAccessIndexReq
	RoleAccessIndexResp                   = sys.RoleAccessIndexResp
	RoleAccessMultiUpdateReq              = sys.RoleAccessMultiUpdateReq
	RoleApiAuthReq                        = sys.RoleApiAuthReq
	RoleApiAuthResp                       = sys.RoleApiAuthResp
	RoleAppIndexReq                       = sys.RoleAppIndexReq
	RoleAppIndexResp                      = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq                 = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq                      = sys.RoleAppUpdateReq
	RoleInfo                              = sys.RoleInfo
	RoleInfoIndexReq                      = sys.RoleInfoIndexReq
	RoleInfoIndexResp                     = sys.RoleInfoIndexResp
	RoleMenuIndexReq                      = sys.RoleMenuIndexReq
	RoleMenuIndexResp                     = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq                = sys.RoleMenuMultiUpdateReq
	RoleModuleIndexReq                    = sys.RoleModuleIndexReq
	RoleModuleIndexResp                   = sys.RoleModuleIndexResp
	RoleModuleMultiUpdateReq              = sys.RoleModuleMultiUpdateReq
	SendOption                            = sys.SendOption
	ServiceInfo                           = sys.ServiceInfo
	SlotInfo                              = sys.SlotInfo
	SlotInfoIndexReq                      = sys.SlotInfoIndexReq
	SlotInfoIndexResp                     = sys.SlotInfoIndexResp
	TenantAccessIndexReq                  = sys.TenantAccessIndexReq
	TenantAccessIndexResp                 = sys.TenantAccessIndexResp
	TenantAccessMultiSaveReq              = sys.TenantAccessMultiSaveReq
	TenantAgreement                       = sys.TenantAgreement
	TenantAgreementIndexReq               = sys.TenantAgreementIndexReq
	TenantAgreementIndexResp              = sys.TenantAgreementIndexResp
	TenantAppIndexReq                     = sys.TenantAppIndexReq
	TenantAppIndexResp                    = sys.TenantAppIndexResp
	TenantAppInfo                         = sys.TenantAppInfo
	TenantAppMenu                         = sys.TenantAppMenu
	TenantAppMenuIndexReq                 = sys.TenantAppMenuIndexReq
	TenantAppMenuIndexResp                = sys.TenantAppMenuIndexResp
	TenantAppModule                       = sys.TenantAppModule
	TenantAppMultiUpdateReq               = sys.TenantAppMultiUpdateReq
	TenantAppWithIDOrCode                 = sys.TenantAppWithIDOrCode
	TenantConfig                          = sys.TenantConfig
	TenantConfigRegisterAutoCreateArea    = sys.TenantConfigRegisterAutoCreateArea
	TenantConfigRegisterAutoCreateProject = sys.TenantConfigRegisterAutoCreateProject
	TenantInfo                            = sys.TenantInfo
	TenantInfoCreateReq                   = sys.TenantInfoCreateReq
	TenantInfoIndexReq                    = sys.TenantInfoIndexReq
	TenantInfoIndexResp                   = sys.TenantInfoIndexResp
	TenantModuleCreateReq                 = sys.TenantModuleCreateReq
	TenantModuleIndexReq                  = sys.TenantModuleIndexReq
	TenantModuleIndexResp                 = sys.TenantModuleIndexResp
	TenantModuleWithIDOrCode              = sys.TenantModuleWithIDOrCode
	TenantOpenCheckTokenReq               = sys.TenantOpenCheckTokenReq
	TenantOpenCheckTokenResp              = sys.TenantOpenCheckTokenResp
	TenantOpenWebHook                     = sys.TenantOpenWebHook
	ThirdApp                              = sys.ThirdApp
	ThirdAppConfig                        = sys.ThirdAppConfig
	ThirdDeptInfoIndexReq                 = sys.ThirdDeptInfoIndexReq
	ThirdDeptInfoReadReq                  = sys.ThirdDeptInfoReadReq
	ThirdEmail                            = sys.ThirdEmail
	ThirdSms                              = sys.ThirdSms
	TimeRange                             = sys.TimeRange
	UserAreaApplyCreateReq                = sys.UserAreaApplyCreateReq
	UserAreaApplyDealReq                  = sys.UserAreaApplyDealReq
	UserAreaApplyIndexReq                 = sys.UserAreaApplyIndexReq
	UserAreaApplyIndexResp                = sys.UserAreaApplyIndexResp
	UserAreaApplyInfo                     = sys.UserAreaApplyInfo
	UserBindAccountReq                    = sys.UserBindAccountReq
	UserCaptchaReq                        = sys.UserCaptchaReq
	UserCaptchaResp                       = sys.UserCaptchaResp
	UserChangePwdReq                      = sys.UserChangePwdReq
	UserCheckTokenReq                     = sys.UserCheckTokenReq
	UserCheckTokenResp                    = sys.UserCheckTokenResp
	UserCodeToUserIDReq                   = sys.UserCodeToUserIDReq
	UserCodeToUserIDResp                  = sys.UserCodeToUserIDResp
	UserCreateResp                        = sys.UserCreateResp
	UserDeptIndexReq                      = sys.UserDeptIndexReq
	UserDeptIndexResp                     = sys.UserDeptIndexResp
	UserDeptMultiSaveReq                  = sys.UserDeptMultiSaveReq
	UserForgetPwdReq                      = sys.UserForgetPwdReq
	UserInfo                              = sys.UserInfo
	UserInfoCreateReq                     = sys.UserInfoCreateReq
	UserInfoDeleteReq                     = sys.UserInfoDeleteReq
	UserInfoIndexReq                      = sys.UserInfoIndexReq
	UserInfoIndexResp                     = sys.UserInfoIndexResp
	UserInfoReadReq                       = sys.UserInfoReadReq
	UserInfoUpdateReq                     = sys.UserInfoUpdateReq
	UserLoginReq                          = sys.UserLoginReq
	UserLoginResp                         = sys.UserLoginResp
	UserMessage                           = sys.UserMessage
	UserMessageIndexReq                   = sys.UserMessageIndexReq
	UserMessageIndexResp                  = sys.UserMessageIndexResp
	UserMessageStatistics                 = sys.UserMessageStatistics
	UserMessageStatisticsResp             = sys.UserMessageStatisticsResp
	UserProfile                           = sys.UserProfile
	UserProfileIndexReq                   = sys.UserProfileIndexReq
	UserProfileIndexResp                  = sys.UserProfileIndexResp
	UserRegisterReq                       = sys.UserRegisterReq
	UserRegisterResp                      = sys.UserRegisterResp
	UserRoleIndexReq                      = sys.UserRoleIndexReq
	UserRoleIndexResp                     = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq                = sys.UserRoleMultiUpdateReq
	WeatherAir                            = sys.WeatherAir
	WeatherReadReq                        = sys.WeatherReadReq
	WeatherReadResp                       = sys.WeatherReadResp
	WithAppCodeID                         = sys.WithAppCodeID
	WithCode                              = sys.WithCode
	WithID                                = sys.WithID
	WithIDCode                            = sys.WithIDCode

	TenantManage interface {
		// 新增区域
		TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error)
		// 更新区域
		TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Empty, error)
		// 删除区域
		TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Empty, error)
		// 获取租户信息详情
		TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error)
		// 获取租户信息列表
		TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error)
		TenantConfigUpdate(ctx context.Context, in *TenantConfig, opts ...grpc.CallOption) (*Empty, error)
		TenantConfigRead(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantConfig, error)
		TenantAccessMultiDelete(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		TenantAccessMultiCreate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		TenantAccessMultiUpdate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		TenantAccessIndex(ctx context.Context, in *TenantAccessIndexReq, opts ...grpc.CallOption) (*TenantAccessIndexResp, error)
		TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error)
		TenantAppCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error)
		TenantAppRead(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*TenantAppInfo, error)
		TenantAppUpdate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error)
		TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Empty, error)
		TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error)
		TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Empty, error)
		TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error)
		TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Empty, error)
		TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error)
		TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error)
		TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Empty, error)
		TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Empty, error)
		TenantOpenCheckToken(ctx context.Context, in *TenantOpenCheckTokenReq, opts ...grpc.CallOption) (*TenantOpenCheckTokenResp, error)
		TenantOpenWebHook(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantOpenWebHook, error)
		TenantAgreementIndex(ctx context.Context, in *TenantAgreementIndexReq, opts ...grpc.CallOption) (*TenantAgreementIndexResp, error)
		TenantAgreementUpdate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*Empty, error)
		TenantAgreementCreate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*WithID, error)
		TenantAgreementRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantAgreement, error)
		TenantAgreementDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultTenantManage struct {
		cli zrpc.Client
	}

	directTenantManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.TenantManageServer
	}
)

func NewTenantManage(cli zrpc.Client) TenantManage {
	return &defaultTenantManage{
		cli: cli,
	}
}

func NewDirectTenantManage(svcCtx *svc.ServiceContext, svr sys.TenantManageServer) TenantManage {
	return &directTenantManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增区域
func (m *defaultTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoCreate(ctx, in, opts...)
}

// 新增区域
func (d *directTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.TenantInfoCreate(ctx, in)
}

// 更新区域
func (m *defaultTenantManage) TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoUpdate(ctx, in, opts...)
}

// 更新区域
func (d *directTenantManage) TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantInfoUpdate(ctx, in)
}

// 删除区域
func (m *defaultTenantManage) TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoDelete(ctx, in, opts...)
}

// 删除区域
func (d *directTenantManage) TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantInfoDelete(ctx, in)
}

// 获取租户信息详情
func (m *defaultTenantManage) TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoRead(ctx, in, opts...)
}

// 获取租户信息详情
func (d *directTenantManage) TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
	return d.svr.TenantInfoRead(ctx, in)
}

// 获取租户信息列表
func (m *defaultTenantManage) TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoIndex(ctx, in, opts...)
}

// 获取租户信息列表
func (d *directTenantManage) TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error) {
	return d.svr.TenantInfoIndex(ctx, in)
}

func (m *defaultTenantManage) TenantConfigUpdate(ctx context.Context, in *TenantConfig, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantConfigUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantConfigUpdate(ctx context.Context, in *TenantConfig, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantConfigUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantConfigRead(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantConfig, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantConfigRead(ctx, in, opts...)
}

func (d *directTenantManage) TenantConfigRead(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantConfig, error) {
	return d.svr.TenantConfigRead(ctx, in)
}

func (m *defaultTenantManage) TenantAccessMultiDelete(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAccessMultiDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAccessMultiDelete(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAccessMultiDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAccessMultiCreate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAccessMultiCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAccessMultiCreate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAccessMultiCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAccessMultiUpdate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAccessMultiUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAccessMultiUpdate(ctx context.Context, in *TenantAccessMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAccessMultiUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAccessIndex(ctx context.Context, in *TenantAccessIndexReq, opts ...grpc.CallOption) (*TenantAccessIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAccessIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAccessIndex(ctx context.Context, in *TenantAccessIndexReq, opts ...grpc.CallOption) (*TenantAccessIndexResp, error) {
	return d.svr.TenantAccessIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error) {
	return d.svr.TenantAppIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppRead(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*TenantAppInfo, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppRead(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppRead(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*TenantAppInfo, error) {
	return d.svr.TenantAppRead(ctx, in)
}

func (m *defaultTenantManage) TenantAppUpdate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppUpdate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleMultiCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppModuleMultiCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppModuleCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error) {
	return d.svr.TenantAppModuleIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppModuleDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.TenantAppMenuCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error) {
	return d.svr.TenantAppMenuIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppMenuUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAppMenuDelete(ctx, in)
}

func (m *defaultTenantManage) TenantOpenCheckToken(ctx context.Context, in *TenantOpenCheckTokenReq, opts ...grpc.CallOption) (*TenantOpenCheckTokenResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantOpenCheckToken(ctx, in, opts...)
}

func (d *directTenantManage) TenantOpenCheckToken(ctx context.Context, in *TenantOpenCheckTokenReq, opts ...grpc.CallOption) (*TenantOpenCheckTokenResp, error) {
	return d.svr.TenantOpenCheckToken(ctx, in)
}

func (m *defaultTenantManage) TenantOpenWebHook(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantOpenWebHook, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantOpenWebHook(ctx, in, opts...)
}

func (d *directTenantManage) TenantOpenWebHook(ctx context.Context, in *WithCode, opts ...grpc.CallOption) (*TenantOpenWebHook, error) {
	return d.svr.TenantOpenWebHook(ctx, in)
}

func (m *defaultTenantManage) TenantAgreementIndex(ctx context.Context, in *TenantAgreementIndexReq, opts ...grpc.CallOption) (*TenantAgreementIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAgreementIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAgreementIndex(ctx context.Context, in *TenantAgreementIndexReq, opts ...grpc.CallOption) (*TenantAgreementIndexResp, error) {
	return d.svr.TenantAgreementIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAgreementUpdate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAgreementUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAgreementUpdate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAgreementUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAgreementCreate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAgreementCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAgreementCreate(ctx context.Context, in *TenantAgreement, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.TenantAgreementCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAgreementRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantAgreement, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAgreementRead(ctx, in, opts...)
}

func (d *directTenantManage) TenantAgreementRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantAgreement, error) {
	return d.svr.TenantAgreementRead(ctx, in)
}

func (m *defaultTenantManage) TenantAgreementDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAgreementDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAgreementDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.TenantAgreementDelete(ctx, in)
}

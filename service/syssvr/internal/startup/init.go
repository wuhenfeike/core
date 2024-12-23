package startup

import (
	"context"
	"gitee.com/unitedrhino/core/service/syssvr/domain/module"
	"gitee.com/unitedrhino/core/service/syssvr/internal/logic"
	areamanagelogic "gitee.com/unitedrhino/core/service/syssvr/internal/logic/areamanage"
	modulemanagelogic "gitee.com/unitedrhino/core/service/syssvr/internal/logic/modulemanage"
	projectmanagelogic "gitee.com/unitedrhino/core/service/syssvr/internal/logic/projectmanage"
	tenantmanagelogic "gitee.com/unitedrhino/core/service/syssvr/internal/logic/tenantmanage"
	usermanagelogic "gitee.com/unitedrhino/core/service/syssvr/internal/logic/usermanage"
	"gitee.com/unitedrhino/core/service/syssvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/core/service/syssvr/internal/svc"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/share/caches"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/domain/tenant"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/share/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Init(svcCtx *svc.ServiceContext) {
	ctx := context.Background()
	utils.Go(ctx, func() {
		list, err := relationDB.NewTenantInfoRepo(ctx).FindByFilter(ctx, relationDB.TenantInfoFilter{}, nil)
		logx.Must(err)
		err = caches.InitTenant(ctx, logic.ToTenantInfoCaches(list)...)
		logx.Must(err)
	})
	InitCache(svcCtx)
	TableInit(svcCtx)
}

func TableInit(svcCtx *svc.ServiceContext) {
	if !relationDB.NeedInitColumn {
		return
	}
	root := "./etc/init/module/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		moduleCode, _ := strings.CutSuffix(info.Name(), ".json")
		_, err = modulemanagelogic.NewModuleMenuMultiImportLogic(ctxs.WithRoot(context.TODO()), svcCtx).ModuleMenuMultiImport(&sys.MenuMultiImportReq{
			ModuleCode: moduleCode,
			Mode:       module.MenuImportModeAll,
			Menu:       string(body),
		})
		return err
	})
	logx.Error(err)
	return
}

func InitCache(svcCtx *svc.ServiceContext) {
	{
		tenantCache, err := caches.NewCache(caches.CacheConfig[tenant.Info, string]{
			KeyType:   eventBus.ServerCacheKeySysTenantInfo,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key string) (*tenant.Info, error) {
				db := relationDB.NewTenantInfoRepo(ctx)
				if key == "" {
					key = ctxs.GetUserCtxNoNil(ctx).TenantCode
				}
				pi, err := db.FindOneByFilter(ctx, relationDB.TenantInfoFilter{
					Codes: []string{key}})
				pb := logic.ToTenantInfoCache(pi)
				return pb, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.TenantCache = tenantCache
	}
	{
		tenantCache, err := caches.NewCache(caches.CacheConfig[sys.TenantConfig, string]{
			KeyType:   eventBus.ServerCacheKeySysTenantConfig,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key string) (*sys.TenantConfig, error) {
				db := relationDB.NewTenantConfigRepo(ctx)
				if key == "" {
					key = ctxs.GetUserCtxNoNil(ctx).TenantCode
				}
				pi, err := db.FindOneByFilter(ctx, relationDB.TenantConfigFilter{
					TenantCode: key})
				pb := tenantmanagelogic.ToTenantConfigPb(ctx, svcCtx, pi)
				return pb, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.TenantConfigCache = tenantCache
	}
	{
		userCache, err := caches.NewCache(caches.CacheConfig[sys.UserInfo, int64]{
			KeyType:   eventBus.ServerCacheKeySysUserInfo,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key int64) (*sys.UserInfo, error) {
				db := relationDB.NewUserInfoRepo(ctx)
				if key == 0 {
					key = ctxs.GetUserCtxNoNil(ctx).UserID
				}
				pi, err := db.FindOne(ctx, key)
				pb := usermanagelogic.UserInfoToPb(ctx, pi, svcCtx)
				return pb, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.UserCache = userCache
	}
	{
		AreaCache, err := caches.NewCache(caches.CacheConfig[sys.AreaInfo, int64]{
			KeyType:   eventBus.ServerCacheKeySysAreaInfo,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key int64) (*sys.AreaInfo, error) {
				db := relationDB.NewAreaInfoRepo(ctx)
				if key == 0 {
					key = ctxs.GetUserCtxNoNil(ctx).UserID
				}
				pi, err := db.FindOne(ctx, key, nil)
				pb := areamanagelogic.TransPoToPb(ctx, pi, svcCtx)
				return pb, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.AreaCache = AreaCache
	}
	{
		projectCache, err := caches.NewCache(caches.CacheConfig[sys.ProjectInfo, int64]{
			KeyType:   eventBus.ServerCacheKeySysProjectInfo,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key int64) (*sys.ProjectInfo, error) {
				db := relationDB.NewProjectInfoRepo(ctx)
				if key == 0 {
					key = ctxs.GetUserCtxNoNil(ctx).ProjectID
				}
				pi, err := db.FindOne(ctx, key)
				pb := projectmanagelogic.ProjectInfoToPb(ctx, svcCtx, pi)
				return pb, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.ProjectCache = projectCache
	}

	{
		c, err := caches.NewCache(caches.CacheConfig[relationDB.SysApiInfo, string]{
			KeyType:   eventBus.ServerCacheKeySysAccessApi,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key string) (*relationDB.SysApiInfo, error) {
				method, path, _ := strings.Cut(key, ":")
				db := relationDB.NewApiInfoRepo(ctx)
				pi, err := db.FindOneByFilter(ctx, relationDB.ApiInfoFilter{
					Route:      path,
					Method:     method,
					WithAccess: true,
				})
				return pi, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.ApiCache = c
	}

	{
		c, err := caches.NewCache(caches.CacheConfig[map[int64]struct{}, string]{
			KeyType:   eventBus.ServerCacheKeySysRoleAccess,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key string) (*map[int64]struct{}, error) {
				db := relationDB.NewRoleAccessRepo(ctx)
				pi, err := db.FindByFilter(ctx, relationDB.RoleAccessFilter{
					AccessCodes: []string{key},
				}, nil)
				if err != nil {
					return nil, err
				}
				var ret = make(map[int64]struct{})
				for _, v := range pi {
					ret[v.RoleID] = struct{}{}
				}
				return &ret, err
			},
			ExpireTime: 20 * time.Minute,
		})
		logx.Must(err)
		svcCtx.RoleAccessCache = c
	}
}

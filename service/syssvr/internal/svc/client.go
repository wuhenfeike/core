package svc

import (
	"context"
	"gitee.com/unitedrhino/core/service/syssvr/internal/config"
	"gitee.com/unitedrhino/core/service/syssvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/share/clients/dingClient"
	"gitee.com/unitedrhino/share/clients/wxClient"
	"gitee.com/unitedrhino/share/conf"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
)

type Clients struct {
	WxOfficial  *wxClient.WxOfficialAccount
	MiniProgram *wxClient.MiniProgram
	DingMini    *dingClient.DingTalk
	Config      *relationDB.SysTenantApp
}
type ClientsManage struct {
	Config config.Config
	sf     syncx.SingleFlight
}

var (
	tc = sync.Map{}
)

func NewClients(c config.Config) *ClientsManage {
	return &ClientsManage{Config: c, sf: syncx.NewSingleFlight()}
}

func (c *ClientsManage) ClearClients(ctx context.Context, appCode string) error {
	uc := ctxs.GetUserCtx(ctx)
	if appCode == "" {
		appCode = uc.AppCode
	}
	var tenantCode = uc.TenantCode
	tc.Delete(tenantCode + appCode)
	return nil
}

func (c *ClientsManage) GetClients(ctx context.Context, appCode string) (Clients, error) {
	uc := ctxs.GetUserCtx(ctx)
	if appCode == "" {
		appCode = uc.AppCode
	}
	var tenantCode = uc.TenantCode
	var key = tenantCode + ":" + appCode
	val, ok := tc.Load(tenantCode + appCode)
	if ok {
		return val.(Clients), nil
	}
	cli, err := c.sf.Do(key, func() (any, error) {
		//如果缓存里没有,需要查库
		cfg, err := relationDB.NewTenantAppRepo(ctx).FindOneByFilter(ctx, relationDB.TenantAppFilter{TenantCode: tenantCode, AppCodes: []string{appCode}})
		if err != nil {
			if !errors.Cmp(err, errors.NotFind) {
				return Clients{}, err
			}
			return Clients{}, errors.Parameter.AddMsg("未配置应用")
		}
		var cli Clients
		cli.Config = cfg
		if cfg.DingMini != nil && cfg.DingMini.AppSecret != "" {
			cli.DingMini, err = dingClient.NewDingTalkClient(&conf.ThirdConf{
				AppID:     cfg.DingMini.AppID,
				AppKey:    cfg.DingMini.AppKey,
				AppSecret: cfg.DingMini.AppSecret,
			})
			if err != nil {
				return Clients{}, err
			}
		}
		if cfg.WxMini != nil && cfg.WxMini.AppSecret != "" {
			cli.MiniProgram, err = wxClient.NewWxMiniProgram(ctx, &conf.ThirdConf{
				AppID:     cfg.WxMini.AppID,
				AppKey:    cfg.WxMini.AppKey,
				AppSecret: cfg.WxMini.AppSecret,
			}, c.Config.CacheRedis)
			if err != nil {
				return Clients{}, err
			}
		}
		if cfg.WxOpen != nil && cfg.WxOpen.AppSecret != "" {
			cli.WxOfficial, err = wxClient.NewWxOfficialAccount(ctx, &conf.ThirdConf{
				AppID:     cfg.WxOpen.AppID,
				AppKey:    cfg.WxOpen.AppKey,
				AppSecret: cfg.WxOpen.AppSecret,
			}, c.Config.CacheRedis)
			if err != nil {
				return Clients{}, err
			}
		}
		tc.Store(tenantCode, cli)
		return cli, nil
	})

	return cli.(Clients), err
}

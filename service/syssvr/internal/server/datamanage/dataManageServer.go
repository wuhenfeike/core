// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: sys.proto

package server

import (
	"context"

	"gitee.com/unitedrhino/core/service/syssvr/internal/logic/datamanage"
	"gitee.com/unitedrhino/core/service/syssvr/internal/svc"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
)

type DataManageServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedDataManageServer
}

func NewDataManageServer(svcCtx *svc.ServiceContext) *DataManageServer {
	return &DataManageServer{
		svcCtx: svcCtx,
	}
}

func (s *DataManageServer) DataProjectCreate(ctx context.Context, in *sys.DataProjectSaveReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataProjectCreateLogic(ctx, s.svcCtx)
	return l.DataProjectCreate(in)
}

func (s *DataManageServer) DataProjectMultiCreate(ctx context.Context, in *sys.DataProjectMultiSaveReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataProjectMultiCreateLogic(ctx, s.svcCtx)
	return l.DataProjectMultiCreate(in)
}

func (s *DataManageServer) DataProjectDelete(ctx context.Context, in *sys.DataProjectDeleteReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataProjectDeleteLogic(ctx, s.svcCtx)
	return l.DataProjectDelete(in)
}

func (s *DataManageServer) DataProjectMultiDelete(ctx context.Context, in *sys.DataProjectMultiDeleteReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataProjectMultiDeleteLogic(ctx, s.svcCtx)
	return l.DataProjectMultiDelete(in)
}

func (s *DataManageServer) DataProjectIndex(ctx context.Context, in *sys.DataProjectIndexReq) (*sys.DataProjectIndexResp, error) {
	l := datamanagelogic.NewDataProjectIndexLogic(ctx, s.svcCtx)
	return l.DataProjectIndex(in)
}

func (s *DataManageServer) DataAreaMultiUpdate(ctx context.Context, in *sys.DataAreaMultiUpdateReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataAreaMultiUpdateLogic(ctx, s.svcCtx)
	return l.DataAreaMultiUpdate(in)
}

func (s *DataManageServer) DataAreaIndex(ctx context.Context, in *sys.DataAreaIndexReq) (*sys.DataAreaIndexResp, error) {
	l := datamanagelogic.NewDataAreaIndexLogic(ctx, s.svcCtx)
	return l.DataAreaIndex(in)
}

func (s *DataManageServer) DataAreaMultiDelete(ctx context.Context, in *sys.DataAreaMultiDeleteReq) (*sys.Empty, error) {
	l := datamanagelogic.NewDataAreaMultiDeleteLogic(ctx, s.svcCtx)
	return l.DataAreaMultiDelete(in)
}

func (s *DataManageServer) UserAreaApplyIndex(ctx context.Context, in *sys.UserAreaApplyIndexReq) (*sys.UserAreaApplyIndexResp, error) {
	l := datamanagelogic.NewUserAreaApplyIndexLogic(ctx, s.svcCtx)
	return l.UserAreaApplyIndex(in)
}

func (s *DataManageServer) UserAreaApplyDeal(ctx context.Context, in *sys.UserAreaApplyDealReq) (*sys.Empty, error) {
	l := datamanagelogic.NewUserAreaApplyDealLogic(ctx, s.svcCtx)
	return l.UserAreaApplyDeal(in)
}

func (s *DataManageServer) DataOpenAccessIndex(ctx context.Context, in *sys.OpenAccessIndexReq) (*sys.OpenAccessIndexResp, error) {
	l := datamanagelogic.NewDataOpenAccessIndexLogic(ctx, s.svcCtx)
	return l.DataOpenAccessIndex(in)
}

func (s *DataManageServer) DataOpenAccessRead(ctx context.Context, in *sys.WithID) (*sys.OpenAccess, error) {
	l := datamanagelogic.NewDataOpenAccessReadLogic(ctx, s.svcCtx)
	return l.DataOpenAccessRead(in)
}

func (s *DataManageServer) DataOpenAccessUpdate(ctx context.Context, in *sys.OpenAccess) (*sys.Empty, error) {
	l := datamanagelogic.NewDataOpenAccessUpdateLogic(ctx, s.svcCtx)
	return l.DataOpenAccessUpdate(in)
}

func (s *DataManageServer) DataOpenAccessCreate(ctx context.Context, in *sys.OpenAccess) (*sys.WithID, error) {
	l := datamanagelogic.NewDataOpenAccessCreateLogic(ctx, s.svcCtx)
	return l.DataOpenAccessCreate(in)
}

func (s *DataManageServer) DataOpenAccessDelete(ctx context.Context, in *sys.WithID) (*sys.Empty, error) {
	l := datamanagelogic.NewDataOpenAccessDeleteLogic(ctx, s.svcCtx)
	return l.DataOpenAccessDelete(in)
}

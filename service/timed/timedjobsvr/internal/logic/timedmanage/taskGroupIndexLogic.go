package timedmanagelogic

import (
	"context"
	"gitee.com/unitedrhino/core/service/timed/internal/repo/relationDB"

	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/internal/svc"
	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/pb/timedjob"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskGroupIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskGroupIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskGroupIndexLogic {
	return &TaskGroupIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskGroupIndexLogic) TaskGroupIndex(in *timedjob.TaskGroupIndexReq) (*timedjob.TaskGroupIndexResp, error) {
	f := relationDB.TaskGroupFilter{}
	repo := relationDB.NewTaskGroupRepo(l.ctx)
	list, err := repo.FindByFilter(l.ctx, f, ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	total, err := repo.CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	return &timedjob.TaskGroupIndexResp{Total: total, List: ToTaskGroupPbs(list)}, nil
}

package self

import (
	"net/http"

	"gitee.com/i-Things/share/result"

	"gitee.com/i-Things/core/service/apisvr/internal/logic/system/user/self"
	"gitee.com/i-Things/core/service/apisvr/internal/svc"
)

func CancelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := self.NewCancelLogic(r.Context(), svcCtx)
		err := l.Cancel()
		result.Http(w, r, nil, err)
	}
}
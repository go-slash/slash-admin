package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/menu"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route delete /menu menu DeleteMenu
//
// "Delete menu information | 删除菜单信息"
//
// "Delete menu information | 删除菜单信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func DeleteMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := menu.NewDeleteMenuLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMenu(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

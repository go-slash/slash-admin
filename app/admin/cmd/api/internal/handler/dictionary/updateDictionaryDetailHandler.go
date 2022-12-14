package dictionary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/dictionary"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route put /dict/detail dictionary UpdateDictionaryDetail
//
// "Update dictionary KV information | 更新字典键值信息"
//
// "Update dictionary KV information | 更新字典键值信息"
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateDictionaryDetailReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateDictionaryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDictionaryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictionary.NewUpdateDictionaryDetailLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDictionaryDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

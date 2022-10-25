package token

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/token"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route POST /token/list token getTokenList
// Get Token list | 获取token列表
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TokenListReq
// Responses:
//   200: TokenListResp
//   401: SimpleMsg
//   500: SimpleMsg

func GetTokenListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TokenListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewGetTokenListLogic(r.Context(), svcCtx)
		resp, err := l.GetTokenList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"slash-admin/app/admin/cmd/api/internal/logic/user"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
)

// swagger:route post /user/profile user UpdateUserProfile
//
//  Update user's profile | 更新用户个人信息
//
//  Update user's profile | 更新用户个人信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateProfileReq
//
//
// Responses:
//  200: SimpleMsgResp
//
//

func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

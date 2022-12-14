// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "slash-admin/app/admin/cmd/api/internal/handler/api"
	authority "slash-admin/app/admin/cmd/api/internal/handler/authority"
	captcha "slash-admin/app/admin/cmd/api/internal/handler/captcha"
	core "slash-admin/app/admin/cmd/api/internal/handler/core"
	dictionary "slash-admin/app/admin/cmd/api/internal/handler/dictionary"
	menu "slash-admin/app/admin/cmd/api/internal/handler/menu"
	oauth "slash-admin/app/admin/cmd/api/internal/handler/oauth"
	role "slash-admin/app/admin/cmd/api/internal/handler/role"
	token "slash-admin/app/admin/cmd/api/internal/handler/token"
	user "slash-admin/app/admin/cmd/api/internal/handler/user"
	"slash-admin/app/admin/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/core/health",
				Handler: core.HealthCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/database",
				Handler: core.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/role/:id",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/role/:id",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/status/:id",
					Handler: role.SetRoleStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/change-password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu",
					Handler: menu.CreateOrUpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/menu",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/list",
					Handler: menu.GetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/role",
					Handler: menu.GetMenuByRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/param",
					Handler: menu.CreateOrUpdateMenuParamHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/param/list",
					Handler: menu.GetMenuParamListByMenuIdHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/menu/param",
					Handler: menu.DeleteMenuParamHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api",
					Handler: api.CreateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api",
					Handler: api.UpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/list",
					Handler: api.GetApiListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/authority/api",
					Handler: authority.CreateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/authority/api",
					Handler: authority.UpdateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/role",
					Handler: authority.GetApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu",
					Handler: authority.CreateMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/authority/menu",
					Handler: authority.UpdateMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/role",
					Handler: authority.GetMenuAuthorityHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dict",
					Handler: dictionary.CreateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/dict",
					Handler: dictionary.UpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/dict",
					Handler: dictionary.DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/list",
					Handler: dictionary.GetDictionaryListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/detail",
					Handler: dictionary.CreateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/dict/detail",
					Handler: dictionary.UpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/dict/detail",
					Handler: dictionary.DeleteDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/detail/list",
					Handler: dictionary.GetDetailByDictionaryNameHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: oauth.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login/callback",
				Handler: oauth.OauthCallbackHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauth/provider",
					Handler: oauth.CreateProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/oauth/provider",
					Handler: oauth.UpdateProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/oauth/provider",
					Handler: oauth.DeleteProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth/provider/list",
					Handler: oauth.GetProviderListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: token.CreateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/token",
					Handler: token.UpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/token",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/status",
					Handler: token.SetTokenStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

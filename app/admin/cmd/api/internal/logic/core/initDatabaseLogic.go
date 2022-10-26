package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"slash-admin/app/admin/cmd/api/internal/globalkey"
	"slash-admin/app/admin/cmd/api/internal/svc"
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	"slash-admin/app/admin/ent/migrate"
	"slash-admin/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.SimpleMsgResp, err error) {
	// add lock to avoid duplicate initialization
	lock := redis.NewRedisLock(l.svcCtx.RedisClient, globalkey.InitDatabaseLock)
	lock.SetExpire(60)

	if ok, err := lock.Acquire(); !ok || err != nil {
		if !ok {
			logx.Error("last initialization is running")
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.InitRunning)
		} else {
			logx.Errorw(errorx.RedisError, logx.Field("detail", err.Error()))
			return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
		}
	}

	defer func() {
		recover()
		lock.Release()
	}()

	// judge if the initialization had been done
	count, err := l.svcCtx.EntClient.SysApi.Query().Count(l.ctx)

	if count > 0 {
		l.Infof("database has been initialized")
		return &types.SimpleMsgResp{Msg: errorx.AlreadyInit}, nil
	}

	if err := l.svcCtx.RedisClient.Set(globalkey.InitDatabaseState, "1"); err != nil {
		logx.Errorw(errorx.RedisError, logx.Field("detail", err.Error()))
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.RedisError)
	}

	// set default state value
	l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, "", 300)
	l.svcCtx.RedisClient.Set(globalkey.InitDatabaseState, "0")

	if err := l.svcCtx.EntClient.Schema.Create(l.ctx, migrate.WithForeignKeys(false)); err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, err.Error(), 300)
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}

	err = l.insertUserData()

	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, err.Error(), 300)
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}

	err = l.insertRoleData()
	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, err.Error(), 300)
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}

	err = l.insertProviderData()
	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, err.Error(), 300)
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}
	// todo
	//err = l.insertRoleMenuAuthorityData()
	//if err != nil {
	//	logx.Errorw(logmessage.DatabaseError, logx.Field("detail", err.Error()))
	//	l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
	//	return nil, status.Error(codes.Internal, err.Error())
	//}
	//err = l.insertCasbinPoliciesData()
	//if err != nil {
	//	logx.Errorw(logmessage.DatabaseError, logx.Field("detail", err.Error()))
	//	l.svcCtx.Redis.Setex("database_error_msg", err.Error(), 300)
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	err = l.insertApiData()
	if err != nil {
		logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
		l.svcCtx.RedisClient.Setex(globalkey.InitDatabaseErrorMsg, err.Error(), 300)
		return nil, errorx.NewApiError(http.StatusInternalServerError, errorx.DatabaseError)
	}

	return &types.SimpleMsgResp{Msg: errorx.Success}, nil

}

func (l *InitDatabaseLogic) insertUserData() error {
	var users = []ent.CreateSysUserInput{
		{
			UUID:     utils.Wrap(uuid.NewString()),
			Username: utils.Wrap("admin"),
			Password: utils.Wrap(utils.BcryptEncrypt("admin")),
			Nickname: utils.Wrap("Admin"),
			Email:    utils.Wrap("admin@gmail.com"),
			RoleID:   utils.Wrap[uint32](1),
		},
	}

	for _, u := range users {
		if _, err := l.svcCtx.EntClient.SysUser.Create().Copy(&u).Save(l.ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *InitDatabaseLogic) insertRoleData() error {
	var roles = []ent.CreateSysRoleInput{
		{
			Name:          utils.Wrap("sys.role.admin"),
			Value:         utils.Wrap("admin"),
			DefaultRouter: utils.Wrap("dashboard"),
			Remark:        utils.Wrap("超级管理员"),
			OrderNo:       utils.Wrap[uint32](1),
		},
		{
			Name:          utils.Wrap("sys.role.stuff"),
			Value:         utils.Wrap("stuff"),
			DefaultRouter: utils.Wrap("dashboard"),
			Remark:        utils.Wrap("普通员工"),
			OrderNo:       utils.Wrap[uint32](2),
		},
		{
			Name:          utils.Wrap("sys.role.member"),
			Value:         utils.Wrap("member"),
			DefaultRouter: utils.Wrap("dashboard"),
			Remark:        utils.Wrap("注册会员"),
			OrderNo:       utils.Wrap[uint32](3),
		},
	}

	for _, r := range roles {
		if _, err := l.svcCtx.EntClient.SysRole.Create().Copy(&r).Save(l.ctx); err != nil {
			return err
		}
	}
	return nil
}

// insert init user data
func (l *InitDatabaseLogic) insertApiData() error {
	apis := []ent.CreateSysApiInput{
		// user
		{
			Path:        utils.Wrap("/user/login"),
			Description: utils.Wrap("api_desc.userLogin"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user/register"),
			Description: utils.Wrap("api_desc.userRegister"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user"),
			Description: utils.Wrap("api_desc.createOrUpdateUser"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user/change-password"),
			Description: utils.Wrap("api_desc.userChangePassword"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user/info"),
			Description: utils.Wrap("api_desc.userInfo"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("GET"),
		},
		{
			Path:        utils.Wrap("/user/list"),
			Description: utils.Wrap("api_desc.userList"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user"),
			Description: utils.Wrap("api_desc.deleteUser"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/user/perm"),
			Description: utils.Wrap("api_desc.userPermissions"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("GET"),
		},
		{
			Path:        utils.Wrap("/user/profile"),
			Description: utils.Wrap("api_desc.userProfile"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("GET"),
		},
		{
			Path:        utils.Wrap("/user/profile"),
			Description: utils.Wrap("api_desc.updateProfile"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/user/logout"),
			Description: utils.Wrap("api_desc.logout"),
			APIGroup:    utils.Wrap("user"),
			Method:      utils.Wrap("GET"),
		},
		// role
		{
			Path:        utils.Wrap("/role"),
			Description: utils.Wrap("api_desc.createOrUpdateRole"),
			APIGroup:    utils.Wrap("role"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/role"),
			Description: utils.Wrap("api_desc.deleteRole"),
			APIGroup:    utils.Wrap("role"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/role/list"),
			Description: utils.Wrap("api_desc.roleList"),
			APIGroup:    utils.Wrap("role"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/role/status"),
			Description: utils.Wrap("api_desc.setRoleStatus"),
			APIGroup:    utils.Wrap("role"),
			Method:      utils.Wrap("POST"),
		},
		// menu
		{
			Path:        utils.Wrap("/menu"),
			Description: utils.Wrap("api_desc.createOrUpdateMenu"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/menu"),
			Description: utils.Wrap("api_desc.deleteMenu"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/menu/list"),
			Description: utils.Wrap("api_desc.menuList"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("GET"),
		},
		{
			Path:        utils.Wrap("/menu/role"),
			Description: utils.Wrap("api_desc.roleMenu"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("GET"),
		},
		{
			Path:        utils.Wrap("/menu/param"),
			Description: utils.Wrap("api_desc.createOrUpdateMenuParam"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/menu/param/list"),
			Description: utils.Wrap("api_desc.menuParamListByMenuId"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/menu/param"),
			Description: utils.Wrap("api_desc.deleteMenuParam"),
			APIGroup:    utils.Wrap("menu"),
			Method:      utils.Wrap("DELETE"),
		},
		// captcha
		{
			Path:        utils.Wrap("/captcha"),
			Description: utils.Wrap("api_desc.captcha"),
			APIGroup:    utils.Wrap("captcha"),
			Method:      utils.Wrap("GET"),
		},
		// authorization
		{
			Path:        utils.Wrap("/authority/api"),
			Description: utils.Wrap("api_desc.createOrUpdateApiAuthority"),
			APIGroup:    utils.Wrap("authority"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/authority/api/role"),
			Description: utils.Wrap("api_desc.APIAuthorityOfRole"),
			APIGroup:    utils.Wrap("authority"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/authority/menu"),
			Description: utils.Wrap("api_desc.createOrUpdateMenuAuthority"),
			APIGroup:    utils.Wrap("authority"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/authority/menu/role"),
			Description: utils.Wrap("api_desc.menuAuthorityOfRole"),
			APIGroup:    utils.Wrap("authority"),
			Method:      utils.Wrap("POST"),
		},
		// api
		{
			Path:        utils.Wrap("/api"),
			Description: utils.Wrap("api_desc.createOrUpdateApi"),
			APIGroup:    utils.Wrap("api"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/api"),
			Description: utils.Wrap("api_desc.deleteAPI"),
			APIGroup:    utils.Wrap("api"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/api/list"),
			Description: utils.Wrap("api_desc.APIList"),
			APIGroup:    utils.Wrap("api"),
			Method:      utils.Wrap("POST"),
		},
		// dictionary
		{
			Path:        utils.Wrap("/dict"),
			Description: utils.Wrap("api_desc.createOrUpdateDictionary"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/dict"),
			Description: utils.Wrap("api_desc.deleteDictionary"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/dict/detail"),
			Description: utils.Wrap("api_desc.deleteDictionaryDetail"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/dict/detail"),
			Description: utils.Wrap("api_desc.createOrUpdateDictionaryDetail"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/dict/detail/list"),
			Description: utils.Wrap("api_desc.getDictionaryListDetail"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/dict/list"),
			Description: utils.Wrap("api_desc.getDictionaryList"),
			APIGroup:    utils.Wrap("dictionary"),
			Method:      utils.Wrap("POST"),
		},
		// oauth APIs
		{
			Path:        utils.Wrap("/oauth/provider"),
			Description: utils.Wrap("api_desc.createOrUpdateProvider"),
			APIGroup:    utils.Wrap("oauth"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/oauth/provider"),
			Description: utils.Wrap("api_desc.deleteProvider"),
			APIGroup:    utils.Wrap("oauth"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/oauth/provider/list"),
			Description: utils.Wrap("api_desc.geProviderList"),
			APIGroup:    utils.Wrap("oauth"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/oauth/login"),
			Description: utils.Wrap("api_desc.oauthLogin"),
			APIGroup:    utils.Wrap("oauth"),
			Method:      utils.Wrap("POST"),
		},
		// token api
		{
			Path:        utils.Wrap("/token"),
			Description: utils.Wrap("api_desc.createOrUpdateToken"),
			APIGroup:    utils.Wrap("token"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/token"),
			Description: utils.Wrap("api_desc.deleteToken"),
			APIGroup:    utils.Wrap("token"),
			Method:      utils.Wrap("DELETE"),
		},
		{
			Path:        utils.Wrap("/token/list"),
			Description: utils.Wrap("api_desc.getTokenList"),
			APIGroup:    utils.Wrap("token"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/token/status"),
			Description: utils.Wrap("api_desc.setTokenStatus"),
			APIGroup:    utils.Wrap("token"),
			Method:      utils.Wrap("POST"),
		},
		{
			Path:        utils.Wrap("/token/logout"),
			Description: utils.Wrap("sys.user.forceLoggingOut"),
			APIGroup:    utils.Wrap("token"),
			Method:      utils.Wrap("POST"),
		},
	}

	for _, api := range apis {
		if _, err := l.svcCtx.EntClient.SysApi.Create().Copy(&api).Save(l.ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *InitDatabaseLogic) insertProviderData() error {
	providers := []ent.CreateSysOauthProviderInput{
		{
			Name:         utils.Wrap("google"),
			ClientID:     utils.Wrap("your client id"),
			ClientSecret: utils.Wrap("your client secret"),
			RedirectURL:  utils.Wrap("http://localhost:3100/oauth/login/callback"),
			Scopes:       utils.Wrap("email openid"),
			AuthURL:      utils.Wrap("https://accounts.google.com/o/oauth2/auth"),
			TokenURL:     utils.Wrap("https://oauth2.googleapis.com/token"),
			AuthStyle:    utils.Wrap[uint8](1),
			InfoURL:      utils.Wrap("https://www.googleapis.com/oauth2/v2/userinfo?access_token="),
		},
		{
			Name:         utils.Wrap("github"),
			ClientID:     utils.Wrap("your client id"),
			ClientSecret: utils.Wrap("your client secret"),
			RedirectURL:  utils.Wrap("http://localhost:3100/oauth/login/callback"),
			Scopes:       utils.Wrap("email openid"),
			AuthURL:      utils.Wrap("https://github.com/login/oauth/authorize"),
			TokenURL:     utils.Wrap("https://github.com/login/oauth/access_token"),
			AuthStyle:    utils.Wrap[uint8](2),
			InfoURL:      utils.Wrap("https://api.github.com/user"),
		},
	}
	for _, p := range providers {
		if _, err := l.svcCtx.EntClient.SysOauthProvider.Create().Copy(&p).Save(l.ctx); err != nil {
			return err
		}

	}
	return nil
}

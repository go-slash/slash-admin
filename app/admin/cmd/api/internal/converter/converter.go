package converter

import (
	"slash-admin/app/admin/cmd/api/internal/types"
	"slash-admin/app/admin/ent"
	pType "slash-admin/pkg/types"
	"time"
)

// goverter:converter
// goverter:extend StatusToUint8
// goverter:extend TimeToUnixMilli
type Converter interface {
	ConvertPagination(input *ent.PageDetails) (output *types.Pagination)

	ConvertSysUser(input *ent.SysUser) (output *types.UserInfo)
	ConvertSysUserList(input []*ent.SysUser) (output []*types.UserInfo)

	ConvertSysRoleToRoleInfo(input *ent.SysRole) (output *types.RoleInfo)
	ConvertSysRoleToRoleInfoList(input []*ent.SysRole) (output []*types.RoleInfo)

	ConvertSysToken(input *ent.SysToken) (output *types.TokenInfo)
	ConvertSysTokenList(input []*ent.SysToken) (output []*types.TokenInfo)

	ConvertSysOauthProvider(input *ent.SysOauthProvider) (output *types.OauthProviderInfo)
	ConvertSysOauthProviderList(input []*ent.SysOauthProvider) (output []*types.OauthProviderInfo)

	// goverter:map Desc Description
	ConvertSysDictionary(input *ent.SysDictionary) (output *types.DictionaryInfo)
	ConvertSysDictionaryList(input []*ent.SysDictionary) (output []*types.DictionaryInfo)

	ConvertSysDictionaryDetail(input *ent.SysDictionaryDetail) (output *types.DictionaryDetailInfo)
	ConvertSysDictionaryDetailList(input []*ent.SysDictionaryDetail) (output []*types.DictionaryDetailInfo)

	// goverter:map APIGroup Group
	ConvertSysApi(input *ent.SysApi) (output *types.ApiInfo)
	ConvertSysApiList(input []*ent.SysApi) (output []*types.ApiInfo)
}

func StatusToUint8(v pType.Status) uint8 {
	return uint8(v)
}

func TimeToUnixMilli(v time.Time) int64 {
	return v.UnixMilli()
}

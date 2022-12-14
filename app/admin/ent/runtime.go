// Code generated by ent, DO NOT EDIT.

package ent

import (
	"slash-admin/app/admin/ent/casbinrule"
	"slash-admin/app/admin/ent/schema"
	"slash-admin/app/admin/ent/sysapi"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/app/admin/ent/sysmenuparam"
	"slash-admin/app/admin/ent/sysoauthprovider"
	"slash-admin/app/admin/ent/sysrole"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/app/admin/ent/sysuser"
	"slash-admin/pkg/types"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbinruleFields := schema.CasbinRule{}.Fields()
	_ = casbinruleFields
	// casbinruleDescPtype is the schema descriptor for Ptype field.
	casbinruleDescPtype := casbinruleFields[0].Descriptor()
	// casbinrule.DefaultPtype holds the default value on creation for the Ptype field.
	casbinrule.DefaultPtype = casbinruleDescPtype.Default.(string)
	// casbinruleDescV0 is the schema descriptor for V0 field.
	casbinruleDescV0 := casbinruleFields[1].Descriptor()
	// casbinrule.DefaultV0 holds the default value on creation for the V0 field.
	casbinrule.DefaultV0 = casbinruleDescV0.Default.(string)
	// casbinruleDescV1 is the schema descriptor for V1 field.
	casbinruleDescV1 := casbinruleFields[2].Descriptor()
	// casbinrule.DefaultV1 holds the default value on creation for the V1 field.
	casbinrule.DefaultV1 = casbinruleDescV1.Default.(string)
	// casbinruleDescV2 is the schema descriptor for V2 field.
	casbinruleDescV2 := casbinruleFields[3].Descriptor()
	// casbinrule.DefaultV2 holds the default value on creation for the V2 field.
	casbinrule.DefaultV2 = casbinruleDescV2.Default.(string)
	// casbinruleDescV3 is the schema descriptor for V3 field.
	casbinruleDescV3 := casbinruleFields[4].Descriptor()
	// casbinrule.DefaultV3 holds the default value on creation for the V3 field.
	casbinrule.DefaultV3 = casbinruleDescV3.Default.(string)
	// casbinruleDescV4 is the schema descriptor for V4 field.
	casbinruleDescV4 := casbinruleFields[5].Descriptor()
	// casbinrule.DefaultV4 holds the default value on creation for the V4 field.
	casbinrule.DefaultV4 = casbinruleDescV4.Default.(string)
	// casbinruleDescV5 is the schema descriptor for V5 field.
	casbinruleDescV5 := casbinruleFields[6].Descriptor()
	// casbinrule.DefaultV5 holds the default value on creation for the V5 field.
	casbinrule.DefaultV5 = casbinruleDescV5.Default.(string)
	sysapiFields := schema.SysApi{}.Fields()
	_ = sysapiFields
	// sysapiDescMethod is the schema descriptor for method field.
	sysapiDescMethod := sysapiFields[4].Descriptor()
	// sysapi.DefaultMethod holds the default value on creation for the method field.
	sysapi.DefaultMethod = sysapiDescMethod.Default.(string)
	// sysapiDescCreatedAt is the schema descriptor for created_at field.
	sysapiDescCreatedAt := sysapiFields[5].Descriptor()
	// sysapi.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysapi.DefaultCreatedAt = sysapiDescCreatedAt.Default.(func() time.Time)
	// sysapiDescUpdatedAt is the schema descriptor for updated_at field.
	sysapiDescUpdatedAt := sysapiFields[6].Descriptor()
	// sysapi.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysapi.DefaultUpdatedAt = sysapiDescUpdatedAt.Default.(func() time.Time)
	// sysapi.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysapi.UpdateDefaultUpdatedAt = sysapiDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysdictionaryFields := schema.SysDictionary{}.Fields()
	_ = sysdictionaryFields
	// sysdictionaryDescStatus is the schema descriptor for status field.
	sysdictionaryDescStatus := sysdictionaryFields[3].Descriptor()
	// sysdictionary.DefaultStatus holds the default value on creation for the status field.
	sysdictionary.DefaultStatus = types.Status(sysdictionaryDescStatus.Default.(uint8))
	// sysdictionaryDescCreatedAt is the schema descriptor for created_at field.
	sysdictionaryDescCreatedAt := sysdictionaryFields[5].Descriptor()
	// sysdictionary.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdictionary.DefaultCreatedAt = sysdictionaryDescCreatedAt.Default.(func() time.Time)
	// sysdictionaryDescUpdatedAt is the schema descriptor for updated_at field.
	sysdictionaryDescUpdatedAt := sysdictionaryFields[6].Descriptor()
	// sysdictionary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdictionary.DefaultUpdatedAt = sysdictionaryDescUpdatedAt.Default.(func() time.Time)
	// sysdictionary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdictionary.UpdateDefaultUpdatedAt = sysdictionaryDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysdictionarydetailFields := schema.SysDictionaryDetail{}.Fields()
	_ = sysdictionarydetailFields
	// sysdictionarydetailDescStatus is the schema descriptor for status field.
	sysdictionarydetailDescStatus := sysdictionarydetailFields[4].Descriptor()
	// sysdictionarydetail.DefaultStatus holds the default value on creation for the status field.
	sysdictionarydetail.DefaultStatus = types.Status(sysdictionarydetailDescStatus.Default.(uint8))
	// sysdictionarydetailDescOrderNo is the schema descriptor for order_no field.
	sysdictionarydetailDescOrderNo := sysdictionarydetailFields[7].Descriptor()
	// sysdictionarydetail.DefaultOrderNo holds the default value on creation for the order_no field.
	sysdictionarydetail.DefaultOrderNo = sysdictionarydetailDescOrderNo.Default.(uint32)
	// sysdictionarydetailDescCreatedAt is the schema descriptor for created_at field.
	sysdictionarydetailDescCreatedAt := sysdictionarydetailFields[8].Descriptor()
	// sysdictionarydetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdictionarydetail.DefaultCreatedAt = sysdictionarydetailDescCreatedAt.Default.(func() time.Time)
	// sysdictionarydetailDescUpdatedAt is the schema descriptor for updated_at field.
	sysdictionarydetailDescUpdatedAt := sysdictionarydetailFields[9].Descriptor()
	// sysdictionarydetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdictionarydetail.DefaultUpdatedAt = sysdictionarydetailDescUpdatedAt.Default.(func() time.Time)
	// sysdictionarydetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdictionarydetail.UpdateDefaultUpdatedAt = sysdictionarydetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysmenuFields := schema.SysMenu{}.Fields()
	_ = sysmenuFields
	// sysmenuDescRedirect is the schema descriptor for redirect field.
	sysmenuDescRedirect := sysmenuFields[6].Descriptor()
	// sysmenu.DefaultRedirect holds the default value on creation for the redirect field.
	sysmenu.DefaultRedirect = sysmenuDescRedirect.Default.(string)
	// sysmenuDescOrderNo is the schema descriptor for order_no field.
	sysmenuDescOrderNo := sysmenuFields[8].Descriptor()
	// sysmenu.DefaultOrderNo holds the default value on creation for the order_no field.
	sysmenu.DefaultOrderNo = sysmenuDescOrderNo.Default.(uint8)
	// sysmenuDescDisabled is the schema descriptor for disabled field.
	sysmenuDescDisabled := sysmenuFields[9].Descriptor()
	// sysmenu.DefaultDisabled holds the default value on creation for the disabled field.
	sysmenu.DefaultDisabled = sysmenuDescDisabled.Default.(bool)
	// sysmenuDescCreatedAt is the schema descriptor for created_at field.
	sysmenuDescCreatedAt := sysmenuFields[11].Descriptor()
	// sysmenu.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysmenu.DefaultCreatedAt = sysmenuDescCreatedAt.Default.(func() time.Time)
	// sysmenuDescUpdatedAt is the schema descriptor for updated_at field.
	sysmenuDescUpdatedAt := sysmenuFields[12].Descriptor()
	// sysmenu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysmenu.DefaultUpdatedAt = sysmenuDescUpdatedAt.Default.(func() time.Time)
	// sysmenu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysmenu.UpdateDefaultUpdatedAt = sysmenuDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysmenuparamFields := schema.SysMenuParam{}.Fields()
	_ = sysmenuparamFields
	// sysmenuparamDescCreatedAt is the schema descriptor for created_at field.
	sysmenuparamDescCreatedAt := sysmenuparamFields[5].Descriptor()
	// sysmenuparam.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysmenuparam.DefaultCreatedAt = sysmenuparamDescCreatedAt.Default.(func() time.Time)
	// sysmenuparamDescUpdatedAt is the schema descriptor for updated_at field.
	sysmenuparamDescUpdatedAt := sysmenuparamFields[6].Descriptor()
	// sysmenuparam.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysmenuparam.DefaultUpdatedAt = sysmenuparamDescUpdatedAt.Default.(func() time.Time)
	// sysmenuparam.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysmenuparam.UpdateDefaultUpdatedAt = sysmenuparamDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysoauthproviderFields := schema.SysOauthProvider{}.Fields()
	_ = sysoauthproviderFields
	// sysoauthproviderDescCreatedAt is the schema descriptor for created_at field.
	sysoauthproviderDescCreatedAt := sysoauthproviderFields[10].Descriptor()
	// sysoauthprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysoauthprovider.DefaultCreatedAt = sysoauthproviderDescCreatedAt.Default.(func() time.Time)
	// sysoauthproviderDescUpdatedAt is the schema descriptor for updated_at field.
	sysoauthproviderDescUpdatedAt := sysoauthproviderFields[11].Descriptor()
	// sysoauthprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysoauthprovider.DefaultUpdatedAt = sysoauthproviderDescUpdatedAt.Default.(func() time.Time)
	// sysoauthprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysoauthprovider.UpdateDefaultUpdatedAt = sysoauthproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysroleFields := schema.SysRole{}.Fields()
	_ = sysroleFields
	// sysroleDescDefaultRouter is the schema descriptor for default_router field.
	sysroleDescDefaultRouter := sysroleFields[3].Descriptor()
	// sysrole.DefaultDefaultRouter holds the default value on creation for the default_router field.
	sysrole.DefaultDefaultRouter = sysroleDescDefaultRouter.Default.(string)
	// sysroleDescStatus is the schema descriptor for status field.
	sysroleDescStatus := sysroleFields[4].Descriptor()
	// sysrole.DefaultStatus holds the default value on creation for the status field.
	sysrole.DefaultStatus = types.Status(sysroleDescStatus.Default.(uint8))
	// sysroleDescRemark is the schema descriptor for remark field.
	sysroleDescRemark := sysroleFields[5].Descriptor()
	// sysrole.DefaultRemark holds the default value on creation for the remark field.
	sysrole.DefaultRemark = sysroleDescRemark.Default.(string)
	// sysroleDescOrderNo is the schema descriptor for order_no field.
	sysroleDescOrderNo := sysroleFields[6].Descriptor()
	// sysrole.DefaultOrderNo holds the default value on creation for the order_no field.
	sysrole.DefaultOrderNo = sysroleDescOrderNo.Default.(uint32)
	// sysroleDescCreatedAt is the schema descriptor for created_at field.
	sysroleDescCreatedAt := sysroleFields[7].Descriptor()
	// sysrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysrole.DefaultCreatedAt = sysroleDescCreatedAt.Default.(func() time.Time)
	// sysroleDescUpdatedAt is the schema descriptor for updated_at field.
	sysroleDescUpdatedAt := sysroleFields[8].Descriptor()
	// sysrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysrole.DefaultUpdatedAt = sysroleDescUpdatedAt.Default.(func() time.Time)
	// sysrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysrole.UpdateDefaultUpdatedAt = sysroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	systokenFields := schema.SysToken{}.Fields()
	_ = systokenFields
	// systokenDescStatus is the schema descriptor for status field.
	systokenDescStatus := systokenFields[4].Descriptor()
	// systoken.DefaultStatus holds the default value on creation for the status field.
	systoken.DefaultStatus = types.Status(systokenDescStatus.Default.(uint8))
	// systokenDescCreatedAt is the schema descriptor for created_at field.
	systokenDescCreatedAt := systokenFields[6].Descriptor()
	// systoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	systoken.DefaultCreatedAt = systokenDescCreatedAt.Default.(func() time.Time)
	// systokenDescUpdatedAt is the schema descriptor for updated_at field.
	systokenDescUpdatedAt := systokenFields[7].Descriptor()
	// systoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	systoken.DefaultUpdatedAt = systokenDescUpdatedAt.Default.(func() time.Time)
	// systoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	systoken.UpdateDefaultUpdatedAt = systokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	sysuserFields := schema.SysUser{}.Fields()
	_ = sysuserFields
	// sysuserDescSideMode is the schema descriptor for side_mode field.
	sysuserDescSideMode := sysuserFields[5].Descriptor()
	// sysuser.DefaultSideMode holds the default value on creation for the side_mode field.
	sysuser.DefaultSideMode = sysuserDescSideMode.Default.(string)
	// sysuserDescAvatar is the schema descriptor for avatar field.
	sysuserDescAvatar := sysuserFields[6].Descriptor()
	// sysuser.DefaultAvatar holds the default value on creation for the avatar field.
	sysuser.DefaultAvatar = sysuserDescAvatar.Default.(string)
	// sysuserDescBaseColor is the schema descriptor for base_color field.
	sysuserDescBaseColor := sysuserFields[7].Descriptor()
	// sysuser.DefaultBaseColor holds the default value on creation for the base_color field.
	sysuser.DefaultBaseColor = sysuserDescBaseColor.Default.(string)
	// sysuserDescActiveColor is the schema descriptor for active_color field.
	sysuserDescActiveColor := sysuserFields[8].Descriptor()
	// sysuser.DefaultActiveColor holds the default value on creation for the active_color field.
	sysuser.DefaultActiveColor = sysuserDescActiveColor.Default.(string)
	// sysuserDescRoleID is the schema descriptor for role_id field.
	sysuserDescRoleID := sysuserFields[9].Descriptor()
	// sysuser.DefaultRoleID holds the default value on creation for the role_id field.
	sysuser.DefaultRoleID = sysuserDescRoleID.Default.(uint64)
	// sysuserDescStatus is the schema descriptor for status field.
	sysuserDescStatus := sysuserFields[12].Descriptor()
	// sysuser.DefaultStatus holds the default value on creation for the status field.
	sysuser.DefaultStatus = types.Status(sysuserDescStatus.Default.(uint8))
	// sysuserDescCreatedAt is the schema descriptor for created_at field.
	sysuserDescCreatedAt := sysuserFields[13].Descriptor()
	// sysuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysuser.DefaultCreatedAt = sysuserDescCreatedAt.Default.(func() time.Time)
	// sysuserDescUpdatedAt is the schema descriptor for updated_at field.
	sysuserDescUpdatedAt := sysuserFields[14].Descriptor()
	// sysuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysuser.DefaultUpdatedAt = sysuserDescUpdatedAt.Default.(func() time.Time)
	// sysuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysuser.UpdateDefaultUpdatedAt = sysuserDescUpdatedAt.UpdateDefault.(func() time.Time)
}

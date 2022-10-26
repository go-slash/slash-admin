// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysRoleColumns holds the columns for the "sys_role" table.
	SysRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "api_group", Type: field.TypeString},
		{Name: "method", Type: field.TypeString, Default: "POST"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysRoleTable holds the schema information for the "sys_role" table.
	SysRoleTable = &schema.Table{
		Name:       "sys_role",
		Columns:    SysRoleColumns,
		PrimaryKey: []*schema.Column{SysRoleColumns[0]},
	}
	// SysDictionaryColumns holds the columns for the "sys_dictionary" table.
	SysDictionaryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeUint8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "desc", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysDictionaryTable holds the schema information for the "sys_dictionary" table.
	SysDictionaryTable = &schema.Table{
		Name:       "sys_dictionary",
		Columns:    SysDictionaryColumns,
		PrimaryKey: []*schema.Column{SysDictionaryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysdictionary_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictionaryColumns[7]},
			},
		},
	}
	// SysDictionaryDetailColumns holds the columns for the "sys_dictionary_detail" table.
	SysDictionaryDetailColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "status", Type: field.TypeUint8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "remark", Type: field.TypeString},
		{Name: "order_no", Type: field.TypeUint32, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysDictionaryDetailTable holds the schema information for the "sys_dictionary_detail" table.
	SysDictionaryDetailTable = &schema.Table{
		Name:       "sys_dictionary_detail",
		Columns:    SysDictionaryDetailColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysdictionarydetail_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysDictionaryDetailColumns[9]},
			},
		},
	}
	// SysRoleColumns holds the columns for the "sys_role" table.
	SysRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "client_id", Type: field.TypeString},
		{Name: "client_secret", Type: field.TypeString},
		{Name: "redirect_url", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeString},
		{Name: "auth_url", Type: field.TypeString},
		{Name: "token_url", Type: field.TypeString},
		{Name: "auth_style", Type: field.TypeUint8},
		{Name: "info_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysRoleTable holds the schema information for the "sys_role" table.
	SysRoleTable = &schema.Table{
		Name:       "sys_role",
		Columns:    SysRoleColumns,
		PrimaryKey: []*schema.Column{SysRoleColumns[0]},
	}
	// SysRoleColumns holds the columns for the "sys_role" table.
	SysRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Unique: true},
		{Name: "default_router", Type: field.TypeString, Default: "dashboard"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "remark", Type: field.TypeString},
		{Name: "order_no", Type: field.TypeUint32, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysRoleTable holds the schema information for the "sys_role" table.
	SysRoleTable = &schema.Table{
		Name:       "sys_role",
		Columns:    SysRoleColumns,
		PrimaryKey: []*schema.Column{SysRoleColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysoauthprovider_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleColumns[9]},
			},
			{
				Name:    "sysrole_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysRoleColumns[9]},
			},
		},
	}
	// SysTokenColumns holds the columns for the "sys_token" table.
	SysTokenColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "status", Type: field.TypeUint8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "expired_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysTokenTable holds the schema information for the "sys_token" table.
	SysTokenTable = &schema.Table{
		Name:       "sys_token",
		Columns:    SysTokenColumns,
		PrimaryKey: []*schema.Column{SysTokenColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "systoken_uuid",
				Unique:  false,
				Columns: []*schema.Column{SysTokenColumns[1]},
			},
			{
				Name:    "systoken_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysTokenColumns[8]},
			},
			{
				Name:    "systoken_expired_at",
				Unique:  false,
				Columns: []*schema.Column{SysTokenColumns[5]},
			},
		},
	}
	// SysUserColumns holds the columns for the "sys_user" table.
	SysUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "side_mode", Type: field.TypeString, Default: "dark"},
		{Name: "avatar", Type: field.TypeString},
		{Name: "base_color", Type: field.TypeString, Default: "#fff"},
		{Name: "active_color", Type: field.TypeString, Default: "#1890ff"},
		{Name: "role_id", Type: field.TypeUint32, Default: 2},
		{Name: "mobile", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "status", Type: field.TypeUint8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SysUserTable holds the schema information for the "sys_user" table.
	SysUserTable = &schema.Table{
		Name:       "sys_user",
		Columns:    SysUserColumns,
		PrimaryKey: []*schema.Column{SysUserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sysuser_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{SysUserColumns[15]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysRoleTable,
		SysDictionaryTable,
		SysDictionaryDetailTable,
		SysRoleTable,
		SysRoleTable,
		SysTokenTable,
		SysUserTable,
	}
)

func init() {
	SysRoleTable.Annotation = &entsql.Annotation{
		Table: "sys_role",
	}
	SysDictionaryTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary",
	}
	SysDictionaryDetailTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary_detail",
	}
	SysRoleTable.Annotation = &entsql.Annotation{
		Table: "sys_role",
	}
	SysRoleTable.Annotation = &entsql.Annotation{
		Table: "sys_role",
	}
	SysTokenTable.Annotation = &entsql.Annotation{
		Table: "sys_token",
	}
	SysUserTable.Annotation = &entsql.Annotation{
		Table: "sys_user",
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"slash-admin/pkg/types"
	"time"
)

type SysUser struct {
	ent.Schema
}

func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uuid").Comment("用户 UUID"),
		field.String("username").Unique().Comment("登录名"),
		field.String("password").Comment("密码"),
		field.String("nickname").Unique().Comment("昵称"),
		field.String("side_mode").Optional().Default("dark").Comment("布局方式"),

		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("头像路径"),

		field.String("base_color").Optional().Default("#fff").Comment("后台页面色调"),
		field.String("active_color").Optional().Default("#1890ff").Comment("当前激活的颜色设定"),
		field.Uint64("role_id").Optional().Default(2).Comment("角色ID"),
		field.String("mobile").Optional().Comment("手机号"),
		field.String("email").Optional().Comment("邮箱号"),

		field.Uint8("status").
			SchemaType(map[string]string{dialect.MySQL: "tinyint unsigned"}).
			GoType(types.Status(0)).
			Default(0).
			Optional().
			Comment("0=正常 1=禁用"),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", SysRole.Type).
			Field("role_id").
			Ref("role").
			Unique().
			Comment("角色"),
	}
}

func (SysUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user"},
	}
}

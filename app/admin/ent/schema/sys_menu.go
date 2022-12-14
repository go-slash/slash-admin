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

type SysMenu struct {
	ent.Schema
}

func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("parent_id").Optional().Comment("parent menu ID"),
		field.Uint8("menu_level").Comment("menu level"),
		field.Uint8("menu_type").Comment("menu type: 0. group 1. menu"),
		field.String("path").Comment("index path"),
		field.String("name").Comment("index name"),
		field.String("redirect").Optional().Default("").Comment("redirect path"),
		field.String("component").Comment("the path of vue file"),
		field.Uint8("order_no").Default(0).Comment("numbers for sorting"),
		field.Bool("disabled").Default(false).Comment("if true, disable"),
		field.String("meta").
			Optional().
			GoType(types.MenuMeta{}).
			SchemaType(map[string]string{
				dialect.MySQL: "JSON",
			}).
			Comment("extra parameters"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", SysRole.Type),
		//https://entgo.io/docs/schema-edges/#o2m-same-type
		edge.To("children", SysMenu.Type).From("parent").Unique().Field("parent_id"),
	}
}

func (SysMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (SysMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu"},
	}
}

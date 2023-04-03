package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// CCategory holds the schema definition for the CCategory entity.
type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "c_category"},
	}
}

// Fields of the CCategory.
func (Category) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Comment("主键id").Unique(),

		field.Int64("parent_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional().Comment("类父级分类"),

		field.Int32("category_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("类别id"),

		field.String("category_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("名称"),

		field.Int32("status").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("状态"),

		field.Int64("create_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional().Comment("创建人id"),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now).Comment("创建时间"),

		field.Int64("modify_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional().Comment("修改人id"),

		field.Time("modify_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now).UpdateDefault(time.Now).Comment("修改时间"),
	}

}

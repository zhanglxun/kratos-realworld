package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"time"
)

// Website holds the schema definition for the Website entity.
type Website struct {
	ent.Schema
}

func (Website) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "c_website"},
	}
}

// Fields of the Website.
func (Website) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Comment("主键ID").Unique(),

		field.Int32("sort_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("排序id"),

		field.Int32("category").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("类别id"),

		field.Int32("type").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("分类,1:AI，2：工具"),

		field.String("website_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("名称"),

		field.String("website_icon").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("图标"),

		field.String("website_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("路径地址"),

		field.String("summary").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Optional().Comment("摘要"),

		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("描述"),

		field.String("status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
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

// Edges of the Website.
func (Website) Edges() []ent.Edge {
	return nil
}

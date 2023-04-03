package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "c_article"},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Unique(),

		field.Int32("sort_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("排序id"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Optional().Comment("标题名称"),

		field.Int32("category_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("分类类别id"),

		field.String("recommend").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("推荐语"),

		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("描述备注"),

		field.Text("content_body").SchemaType(map[string]string{
			dialect.MySQL: "longtext", // Override MySQL.
		}).Optional().Comment("内容体"),

		field.String("image_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Optional().Comment("图片URL"),

		field.Int32("status").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("状态"),

		field.Int32("click_count").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("点击数"),

		field.Int32("like_count").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("喜欢数"),

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

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}

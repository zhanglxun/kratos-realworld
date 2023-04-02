package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"time"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "s_account"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Comment("主键id，系统唯一标识").Unique(),

		field.String("account").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("用户系统账号"),

		field.String("pwd").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("系统密码"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("昵称"),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}).Optional().Comment("用户邮箱"),

		field.String("mobile").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}).Optional().Comment("手机号码"),

		field.Int64("create_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional(),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now),

		field.Int64("modify_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Optional(),

		field.Time("modify_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

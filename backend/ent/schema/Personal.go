package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("Personalname").Match(regexp.MustCompile("[a-zA-Zก-ฮ]")),
		field.String("Email").Match(regexp.MustCompile("[a-zA-Z0-9]+@([a-zA-Z0-9]+.)+[A-Za-z]")),
		field.String("Password").MinLen(4).Match(regexp.MustCompile("[0-9]")),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Customer.Type).StorageKey(edge.Column("personal_id")),
		edge.From("title", Title.Type).Ref("personal").Unique(),
		edge.From("department", Department.Type).Ref("personal").Unique(),
		edge.From("gender", Gender.Type).Ref("personal").Unique(),
		edge.To("product", Product.Type).StorageKey(edge.Column("Personal")),
		edge.To("fix", Fix.Type).StorageKey(edge.Column("personal_id")),
		edge.To("personal", Adminrepair.Type).StorageKey(edge.Column("personal_id")),
		edge.To("receipt", Receipt.Type).StorageKey(edge.Column("personal_id")),
	}
}

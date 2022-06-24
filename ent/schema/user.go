package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	foov1alpha "github.com/rltvty/ent-multiple-proto-bug/gen/ent-multiple-proto-bug/foo/v1alpha"
	barv1alpha "github.com/rltvty/ent-multiple-proto-bug/gen/ent-multiple-proto-bug/bar/v1alpha"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.Bytes("foo").GoType(&foov1alpha.Foo{}),
		field.Bytes("bar").GoType(&barv1alpha.Bar{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

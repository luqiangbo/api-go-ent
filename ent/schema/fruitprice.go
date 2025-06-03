package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FruitPrice holds the schema definition for the FruitPrice entity.
type FruitPrice struct {
	ent.Schema
}

// Fields of the FruitPrice.
func (FruitPrice) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty(),
		field.Float("price").
			Positive(),
		field.String("unit").
			NotEmpty().
			Comment("单位(如: kg, 个, 箱)"),
		field.String("remark").
			NotEmpty().
			Comment("备注信息"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the FruitPrice.
func (FruitPrice) Edges() []ent.Edge {
	return nil
}

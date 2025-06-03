package schema

import (
	"api-go-ent/ent/schema/mixin"
	"api-go-ent/ent/schema/validators"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FruitPrice holds the schema definition for the FruitPrice entity.
type FruitPrice struct {
	ent.Schema
}

// Mixin 返回实体的Mixin列表
func (FruitPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},       // 基础字段
		mixin.AuditMixin{},      // 审计字段
		mixin.SoftDeleteMixin{}, // 软删除
	}
}

// Fields of the FruitPrice.
func (FruitPrice) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(50).
			Comment("水果名称"),
		field.Float("price").
			Validate(validators.Price).
			Comment("价格"),
		field.String("unit").
			NotEmpty().
			MaxLen(10).
			Comment("单位(如: kg, 个, 箱)"),
		field.String("remark").
			NotEmpty().
			MaxLen(200).
			Comment("备注信息"),
	}
}

// Edges of the FruitPrice.
func (FruitPrice) Edges() []ent.Edge {
	return nil
}

// Indexes 返回索引定义
func (FruitPrice) Indexes() []ent.Index {
	return []ent.Index{
		// 可以添加需要的索引
	}
}

// Hooks 返回生命周期钩子
func (FruitPrice) Hooks() []ent.Hook {
	return []ent.Hook{
		// 可以添加需要的钩子
	}
}

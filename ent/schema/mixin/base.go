package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// BaseMixin 实现了基础字段Mixin
type BaseMixin struct {
	mixin.Schema
}

// Fields 返回基础字段
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("主键ID"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}

// TimeMixin 实现了时间字段Mixin
type TimeMixin struct {
	mixin.Schema
}

// Fields 返回时间字段
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}

// AuditMixin 实现了审计字段Mixin
type AuditMixin struct {
	mixin.Schema
}

// Fields 返回审计字段
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			Optional().
			Comment("创建者"),
		field.String("updated_by").
			Optional().
			Comment("更新者"),
	}
}

// SoftDeleteMixin 实现了软删除Mixin
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields 返回软删除字段
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Nillable().
			Comment("删除时间"),
	}
}

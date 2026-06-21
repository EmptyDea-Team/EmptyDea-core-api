package uqholder

import "context"

// Entities 规定实体状态索引查询能力。
type Entities interface {
	// GetByRuntimeID 按运行时 ID 查询实体。
	GetByRuntimeID(ctx context.Context, runtimeID uint64) (entity Entity, existed bool, err error)
	// GetHealth 按运行时 ID 查询实体生命值。
	GetHealth(ctx context.Context, runtimeID uint64) (health float32, ok bool, err error)
	// GetEffectTypes 按运行时 ID 查询实体效果类型。
	GetEffectTypes(ctx context.Context, runtimeID uint64) (effectTypes []int32, existed bool, err error)
}

// Entity 规定单个实体状态查询能力。
type Entity interface {
	// GetRuntimeID 返回实体运行时 ID。
	GetRuntimeID(ctx context.Context) (runtimeID uint64, ok bool, err error)
	// GetUniqueID 返回实体唯一 ID。
	GetUniqueID(ctx context.Context) (uniqueID int64, ok bool, err error)
	// GetHealth 返回实体生命值。
	GetHealth(ctx context.Context) (health float32, ok bool, err error)
	// GetEffectTypes 返回实体效果类型。
	GetEffectTypes(ctx context.Context) (effectTypes []int32, existed bool, err error)
	// GetEffect 返回指定效果类型的药水效果。
	GetEffect(ctx context.Context, effectType int32) (mobEffect MobEffect, existed bool, err error)
}

// MobEffect 规定单个药水效果状态查询能力。
type MobEffect interface {
	// GetEffectType 返回效果类型。
	GetEffectType() (effectType int32)
	// GetAmplifier 返回效果等级。
	GetAmplifier(ctx context.Context) (amplifier int32, ok bool, err error)
	// GetDuration 返回剩余时间。
	GetDuration(ctx context.Context) (duration int32, ok bool, err error)
	// GetParticles 返回是否显示粒子。
	GetParticles(ctx context.Context) (particles bool, ok bool, err error)
	// GetUpdatedTick 返回最后更新时间刻。
	GetUpdatedTick(ctx context.Context) (updatedTick uint64, ok bool, err error)
}

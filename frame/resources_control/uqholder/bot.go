package uqholder

import (
	"context"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// Bot 规定机器人自身状态查询能力。
type Bot interface {
	// GetName 返回机器人玩家名。
	GetName(ctx context.Context) (name string, ok bool, err error)
	// GetXUID 返回机器人 XUID。
	GetXUID(ctx context.Context) (xuid string, ok bool, err error)
	// GetUUID 返回机器人 UUID 字符串。
	GetUUID(ctx context.Context) (uuid string, ok bool, err error)
	// GetEntityUniqueID 返回机器人实体唯一 ID。
	GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error)
	// GetEntityRuntimeID 返回机器人实体运行时 ID。
	GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error)
	// GetPosition 返回机器人当前位置。
	GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error)
	// GetDimension 返回机器人所在维度。
	GetDimension(ctx context.Context) (dimension int32, ok bool, err error)
	// GetGameMode 返回机器人游戏模式。
	GetGameMode(ctx context.Context) (gameMode int32, ok bool, err error)
	// GetHealth 返回机器人生命值。
	GetHealth(ctx context.Context) (health float32, ok bool, err error)
	// GetHunger 返回机器人饥饿值。
	GetHunger(ctx context.Context) (hunger float32, ok bool, err error)
	// GetSaturation 返回机器人饱和度。
	GetSaturation(ctx context.Context) (saturation float32, ok bool, err error)
	// GetHotBarSlot 返回机器人当前快捷栏槽位。
	GetHotBarSlot(ctx context.Context) (hotBarSlot byte, ok bool, err error)
}

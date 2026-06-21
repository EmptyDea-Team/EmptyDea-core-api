package define

import (
	"context"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// Resources 规定资源层客户端集合的标准能力。
type Resources[
	I Inventories[IV],
	IV Inventory,
	CM ContainerManager,
	CP ConstantPacket,
	PL PacketListener,
	U UQHolder[B, PS, P, PA, W, GR, ES, E, ME],
	B Bot,
	PS Players[P, PA],
	P UQPlayer[PA],
	PA PlayerAbilities,
	W World[GR],
	GR GameRule,
	ES Entities[E, ME],
	E Entity[ME],
	ME MobEffect,
] interface {
	// BotInfo 返回机器人基础信息。
	BotInfo(ctx context.Context) (BotInfo, error)
	// WritePacket 向服务端发送数据包。
	WritePacket(ctx context.Context, p *packet_pb.Packet) error
	// Inventories 返回库存相关资源。
	Inventories() I
	// Container 返回容器相关资源。
	Container() CM
	// PacketListener 返回数据包监听相关资源。
	PacketListener() PL
	// ConstantPacket 返回常量数据包相关资源。
	ConstantPacket() CP
	// UQHolder 返回轻量状态集合。
	UQHolder() U
}

// BotInfo 是当前机器人基础登录信息。
type BotInfo struct {
	// BotName 是机器人玩家名。
	BotName string
	// XUID 是机器人 Xbox 用户 ID。
	XUID string
	// EntityUniqueID 是机器人实体唯一 ID。
	EntityUniqueID int64
	// EntityRuntimeID 是机器人实体运行时 ID。
	EntityRuntimeID uint64
}

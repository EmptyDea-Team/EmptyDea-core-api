package game_interface

import (
	"context"

	resources_control "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
)

// BotClick 规定点击操作客户端的标准能力。
type BotClick interface {
	// ChangeSelectedHotbarSlot 切换机器人选中的快捷栏槽位。
	ChangeSelectedHotbarSlot(ctx context.Context, hotbarSlotID resources_control.SlotID) error
	// ClickBlock 对方块执行点击。
	ClickBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks) error
	// ClickBlockWithPosition 以真实点击位置点击方块。
	ClickBlockWithPosition(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, position *mgl32_pb.Vec3) (bool, error)
	// ClickAir 使用指定快捷栏物品点击空气。
	ClickAir(ctx context.Context, hotbarSlot resources_control.SlotID, realPosition *mgl32_pb.Vec3) error
	// PlaceBlock 使用指定请求放置方块。
	PlaceBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, blockFace int32) error
	// PlaceBlockHighLevel 使用高级封装放置方块。
	PlaceBlockHighLevel(ctx context.Context, blockPos *protocol_pb.BlockPos, hotBarSlot resources_control.SlotID, facing uint8) (*protocol_pb.BlockPos, *protocol_pb.BlockPos, *protocol_pb.BlockPos, error)
	// PickBlock 对指定位置执行选取方块。
	PickBlock(ctx context.Context, pos *protocol_pb.BlockPos, assignNBTData bool) (bool, resources_control.SlotID, error)
}

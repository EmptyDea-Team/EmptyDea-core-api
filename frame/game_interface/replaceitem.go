package game_interface

import (
	"context"
	resources_control "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// Replaceitem 规定 replaceitem 客户端的标准能力。
type Replaceitem interface {
	// ReplaceitemInInventory 对实体物品栏执行 replaceitem。
	ReplaceitemInInventory(ctx context.Context, target string, path ReplaceitemPath, itemInfo ReplaceitemInfo, method string, blocked bool) error
	// ReplaceitemInContainerAsync 对容器方块异步执行 replaceitem。
	ReplaceitemInContainerAsync(ctx context.Context, blockPos *protocol_pb.BlockPos, itemInfo ReplaceitemInfo, method string) error
}

// ReplaceitemPath 指示 replaceitem 时目标物品栏的槽位类型。
type ReplaceitemPath string

const (
	// ReplacePathInventoryOnly 表示只操作背包槽位。
	ReplacePathInventoryOnly ReplaceitemPath = "slot.inventory"
	// ReplacePathHotbarOnly 表示只操作快捷栏槽位。
	ReplacePathHotbarOnly ReplaceitemPath = "slot.hotbar"
	// ReplacePathInventory 表示同时允许背包和快捷栏槽位。
	ReplacePathInventory ReplaceitemPath = "slot.inventory | slot.hotbar"
)

// ReplaceitemInfo 指示要通过 replaceitem 生成的物品的基本信息。
type ReplaceitemInfo struct {
	// Name 是物品名称。
	Name string
	// Count 是物品数量。
	Count uint8
	// MetaData 是物品元数据。
	MetaData int16
	// Slot 是目标槽位。
	Slot resources_control.SlotID
}

package game_interface

import (
	"context"
	resources_control "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// ItemCopy 规定物品复制客户端的标准能力。
type ItemCopy interface {
	// CopyItem 将基础物品复制为目标物品。
	CopyItem(ctx context.Context, containerInfo *game_interface_pb.UseItemOnBlocks, baseItems []ItemInfoWithSlot, targetItems []*ItemInfo) error
}

// ItemType 指示该物品在单次操作中的物品类型。
type ItemType uint8

// ItemInfo 是物品的信息。
type ItemInfo struct {
	// Count 是物品数量。
	Count uint8
	// ItemType 指示该物品在单次操作中的物品类型。
	ItemType ItemType
}

// ItemInfoWithSlot 是物品的信息，同时指示该物品位于哪个槽位。
type ItemInfoWithSlot struct {
	// Slot 是物品所在槽位。
	Slot resources_control.SlotID
	// ItemInfo 是该槽位上的物品信息。
	ItemInfo ItemInfo
}

const (
	// IDItemStackOperationMove 是物品移动操作编号。
	IDItemStackOperationMove uint8 = iota
	// IDItemStackOperationSwap 是物品交换操作编号。
	IDItemStackOperationSwap
	// IDItemStackOperationDrop 是物品丢弃操作编号。
	IDItemStackOperationDrop
	// IDItemStackOperationCreativeItem 是创造物品获取操作编号。
	IDItemStackOperationCreativeItem
	// IDItemStackOperationHighLevelRenaming 是高层重命名操作编号。
	IDItemStackOperationHighLevelRenaming
	// IDItemStackOperationHighLevelLooming 是高层织布机操作编号。
	IDItemStackOperationHighLevelLooming
	// IDItemStackOperationHighLevelCrafting 是高层合成操作编号。
	IDItemStackOperationHighLevelCrafting
	// IDItemStackOperationHighLevelTrimming 是高层锻造台纹饰操作编号。
	IDItemStackOperationHighLevelTrimming
	// IDItemStackOperationHighLevelBeaconPayment 是高层信标支付操作编号。
	IDItemStackOperationHighLevelBeaconPayment
)

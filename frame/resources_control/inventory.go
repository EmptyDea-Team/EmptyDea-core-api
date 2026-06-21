package resources_control

import (
	"context"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// Inventories 规定库存集合客户端的标准能力。
type Inventories[IV Inventory] interface {
	// GetInventory 返回指定窗口名对应的库存。
	GetInventory(ctx context.Context, windowName WindowName) (inventory IV, existed bool, err error)
	// GetItemStack 返回指定窗口和槽位上的物品堆栈。
	GetItemStack(ctx context.Context, windowName WindowName, slotID SlotID) (item *protocol_pb.ItemInstance, inventoryExisted bool, err error)
	// GetAllItemStack 返回指定窗口内的所有物品堆栈。
	GetAllItemStack(ctx context.Context, windowName WindowName) (mapping map[SlotID]*protocol_pb.ItemInstance, inventoryExisted bool, err error)
	// GetAllWindowName 返回当前所有窗口名。
	GetAllWindowName(ctx context.Context) (result []WindowName, err error)
}

// Inventory 规定单个库存客户端的标准能力。
type Inventory interface {
	// GetItemStack 返回当前库存指定槽位上的物品堆栈。
	GetItemStack(ctx context.Context, slotID SlotID) (item *protocol_pb.ItemInstance, err error)
	// GetAllItemStack 返回当前库存的所有物品堆栈。
	GetAllItemStack(ctx context.Context) (mapping map[SlotID]*protocol_pb.ItemInstance, err error)
}

// SlotID 是单个物品栏槽位的索引，它是从 0 开始索引的。
type SlotID uint8

// WindowID 是机器人已打开或持有的库存的窗口 ID。
type WindowID uint32

// DynamicContainerID 是机器人已打开或持有的动态库存的容器 ID。
type DynamicContainerID uint32

// WindowName 唯一标识一个普通窗口或动态容器窗口。
type WindowName struct {
	// WindowID 是窗口 ID。
	WindowID WindowID
	// DynamicContainerID 是动态容器 ID；非动态容器窗口时为 0。
	DynamicContainerID DynamicContainerID
}

// SlotLocation 描述一个物品所在的位置。
type SlotLocation struct {
	WindowName
	// SlotID 是窗口内槽位索引。
	SlotID SlotID
}

const (
	// WindowIDInventory 是玩家背包窗口 ID。
	WindowIDInventory WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDInventory)
	// WindowIDOffHand 是副手窗口 ID。
	WindowIDOffHand WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDOffHand)
	// WindowIDArmour 是盔甲栏窗口 ID。
	WindowIDArmour WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDArmour)
	// WindowIDCrafting 是合成栏窗口 ID。
	WindowIDCrafting WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDCrafting)
	// WindowIDUI 是 UI 窗口 ID。
	WindowIDUI WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDUI)
	// WindowIDDynamic 是动态容器窗口 ID。
	WindowIDDynamic WindowID = WindowID(protocol_pb.WindowIDEnum_WindowIDDynamic)
)

var (
	// WindowNameInventory 是玩家背包窗口名。
	WindowNameInventory = WindowName{WindowID: WindowIDInventory}
	// WindowNameOffHand 是副手窗口名。
	WindowNameOffHand = WindowName{WindowID: WindowIDOffHand}
	// WindowNameArmour 是盔甲栏窗口名。
	WindowNameArmour = WindowName{WindowID: WindowIDArmour}
	// WindowNameCrafting 是合成栏窗口名。
	WindowNameCrafting = WindowName{WindowID: WindowIDCrafting}
	// WindowNameUI 是 UI 窗口名。
	WindowNameUI = WindowName{WindowID: WindowIDUI}
)

// NewDynamicContainerWindowName 基于动态容器 ID 构造动态容器窗口名。
func NewDynamicContainerWindowName(dynamicContainerID DynamicContainerID) WindowName {
	return WindowName{
		WindowID:           WindowIDDynamic,
		DynamicContainerID: dynamicContainerID,
	}
}

// NewWindowName 基于窗口 ID 和动态容器 ID 构造窗口名。
func NewWindowName(windowID WindowID, dynamicContainerID DynamicContainerID) WindowName {
	if windowID != WindowIDDynamic {
		dynamicContainerID = 0
	}
	return WindowName{WindowID: windowID, DynamicContainerID: dynamicContainerID}
}

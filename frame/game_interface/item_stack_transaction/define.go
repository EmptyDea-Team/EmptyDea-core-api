package item_stack_transaction

import (
	"context"

	item_stack_operation "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_operation"
	resources_control "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// ItemStackTransaction 规定单个物品堆栈事务的标准能力。
type ItemStackTransaction[TX any, OP item_stack_operation.ItemStackOperation] interface {
	// Discord 丢弃当前事务内的全部操作。
	Discord() TX
	// Discard 是 Discord 的可读别名。
	Discard() TX
	// Commit 提交当前事务内的全部操作。
	Commit(ctx context.Context) (success bool, pk *packet_pb.ItemStackRequest, serverResponse []*protocol_pb.ItemStackResponse, err error)
	// CommitItemStackOperations 一次性提交指定操作列表。
	CommitItemStackOperations(ctx context.Context, operations []OP) (*game_interface_pb.CommitTransactionResponse, error)
	// MoveItem 将物品从任意位置移动到任意位置。
	MoveItem(source resources_control.SlotLocation, destination resources_control.SlotLocation, count uint8) TX
	// MoveBetweenInventory 在背包槽位之间移动物品。
	MoveBetweenInventory(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveBetweenDynamicContainer 在动态容器槽位之间移动物品。
	MoveBetweenDynamicContainer(dynamicContainerID resources_control.DynamicContainerID, source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveToDynamicContainer 从背包移动物品到动态容器。
	MoveToDynamicContainer(source resources_control.SlotID, dynamicContainerID resources_control.DynamicContainerID, destination resources_control.SlotID, count uint8) TX
	// MoveFromDynamicContainer 从动态容器移动物品到背包。
	MoveFromDynamicContainer(dynamicContainerID resources_control.DynamicContainerID, source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveBetweenContainer 在已打开容器槽位之间移动物品。
	MoveBetweenContainer(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveToContainer 从背包移动物品到已打开容器。
	MoveToContainer(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveToInventory 从已打开容器移动物品到背包。
	MoveToInventory(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveToCraftingTable 从背包移动物品到合成栏。
	MoveToCraftingTable(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// MoveFromCraftingTable 从合成栏移动物品到背包。
	MoveFromCraftingTable(source resources_control.SlotID, destination resources_control.SlotID, count uint8) TX
	// SwapItem 交换任意两个位置的物品。
	SwapItem(source resources_control.SlotLocation, destination resources_control.SlotLocation) TX
	// SwapBetweenInventory 交换背包槽位物品。
	SwapBetweenInventory(source resources_control.SlotID, destination resources_control.SlotID) TX
	// SwapBetweenDynamicContainer 交换动态容器槽位物品。
	SwapBetweenDynamicContainer(dynamicContainerID resources_control.DynamicContainerID, source resources_control.SlotID, destination resources_control.SlotID) TX
	// SwapInventoryBetweenDynamicContainer 交换背包和动态容器槽位物品。
	SwapInventoryBetweenDynamicContainer(source resources_control.SlotID, dynamicContainerID resources_control.DynamicContainerID, destination resources_control.SlotID) TX
	// SwapInventoryBetweenContainer 交换背包和已打开容器槽位物品。
	SwapInventoryBetweenContainer(source resources_control.SlotID, destination resources_control.SlotID) TX
	// DropItem 丢弃任意位置的物品。
	DropItem(slot resources_control.SlotLocation, count uint8) TX
	// DropInventoryItem 丢弃背包槽位物品。
	DropInventoryItem(slot resources_control.SlotID, count uint8) TX
	// DropDynamicContainerItem 丢弃动态容器槽位物品。
	DropDynamicContainerItem(dynamicContainerID resources_control.DynamicContainerID, slot resources_control.SlotID, count uint8) TX
	// DropContainerItem 丢弃已打开容器槽位物品。
	DropContainerItem(slot resources_control.SlotID, count uint8) TX
	// GetCreativeItem 获取创造物品到任意位置。
	GetCreativeItem(creativeItemNetworkID uint32, slot resources_control.SlotLocation, count uint8) TX
	// GetCreativeItemToInventory 获取创造物品到背包槽位。
	GetCreativeItemToInventory(creativeItemNetworkID uint32, slot resources_control.SlotID, count uint8) TX
	// GetCreativeItemToDynamicContainer 获取创造物品到动态容器槽位。
	GetCreativeItemToDynamicContainer(creativeItemNetworkID uint32, dynamicContainerID resources_control.DynamicContainerID, slot resources_control.SlotID, count uint8) TX
	// RenameItem 重命名任意位置的物品。
	RenameItem(slot resources_control.SlotLocation, newName string) TX
	// RenameInventoryItem 重命名背包槽位物品。
	RenameInventoryItem(slot resources_control.SlotID, newName string) TX
	// Looming 执行织布机物品操作。
	Looming(patternName string, patternSlot resources_control.SlotLocation, bannerSlot resources_control.SlotLocation, dyeSlot resources_control.SlotLocation, resultItem item_stack_operation.ExpectedNewItem) TX
	// LoomingFromInventory 使用背包槽位执行织布机物品操作。
	LoomingFromInventory(patternName string, patternSlot resources_control.SlotID, bannerSlot resources_control.SlotID, dyeSlot resources_control.SlotID, resultItem item_stack_operation.ExpectedNewItem) TX
	// Crafting 执行合成物品操作。
	Crafting(recipeNetworkID uint32, resultSlotID resources_control.SlotID, resultCount uint8, resultItem item_stack_operation.ExpectedNewItem) TX
	// Trimming 执行锻造台纹饰操作。
	Trimming(trimItemPath resources_control.SlotLocation, materialPath resources_control.SlotLocation, templatePath resources_control.SlotLocation, resultItem item_stack_operation.ExpectedNewItem) TX
	// TrimmingFromInventory 使用背包槽位执行锻造台纹饰操作。
	TrimmingFromInventory(trimItemSlot resources_control.SlotID, materialSlot resources_control.SlotID, templateSlot resources_control.SlotID, resultItem item_stack_operation.ExpectedNewItem) TX
	// BeaconPayment 执行信标支付操作。
	BeaconPayment(paymentPath resources_control.SlotLocation, primaryEffect int32, secondaryEffect int32) TX
	// BeaconPaymentFromInventory 使用背包槽位执行信标支付操作。
	BeaconPaymentFromInventory(paymentSlot resources_control.SlotID, primaryEffect int32, secondaryEffect int32) TX
}

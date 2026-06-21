package define

import (
	"context"
	"time"

	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// Client 规定 EmptyDea Core 客户端入口的标准形状。
type Client[F Frame, R any, G any] interface {
	// Conn 返回底层 gRPC 连接；非 gRPC 连接实现可以返回 nil。
	Conn() *grpc.ClientConn
	// Frame 返回框架层客户端。
	Frame() F
	// Resources 返回资源层客户端集合。
	Resources() R
	// GameInterface 返回游戏交互层客户端集合。
	GameInterface() G
	// Close 关闭客户端持有的底层连接。
	Close() error
}

// Frame 规定框架层客户端的标准能力。
type Frame interface {
	// StartConnection 使用配置启动 Minecraft 连接。
	StartConnection(ctx context.Context, config FrameConfig) (string, error)
	// StopConnection 主动关闭当前连接。
	StopConnection(ctx context.Context) error
	// GetConnectionState 查询当前连接状态。
	GetConnectionState(ctx context.Context) (ConnectionState, error)
	// Ping 检查 API 服务是否可响应。
	Ping(ctx context.Context) (bool, error)
	// WatchClosed 监听连接关闭事件。
	WatchClosed(ctx context.Context, callback func(ClosedEvent, error)) error
}

// GameInterface 规定游戏交互层客户端集合的标准形状。
type GameInterface[
	C Commands,
	SB StructureBackup,
	Q Querytarget,
	M Movement,
	S SetBlock,
	R Replaceitem,
	B BotClick,
	ISO ItemStackOperation[TX, OP],
	TX ItemStackTransaction[TX, OP],
	OP ItemStackOperationData,
	CO ContainerOpenAndClose,
	IC ItemCopy,
	IT ItemTransition,
	PK PlayerKit[P, AB],
	P Player[AB],
	AB AbilityBuilder[AB],
] interface {
	// Commands 返回命令相关实现。
	Commands() C
	// StructureBackup 返回结构备份相关实现。
	StructureBackup() SB
	// Querytarget 返回 querytarget 相关实现。
	Querytarget() Q
	// Movement 返回移动控制相关实现。
	Movement() M
	// SetBlock 返回方块放置相关实现。
	SetBlock() S
	// Replaceitem 返回 replaceitem 相关实现。
	Replaceitem() R
	// BotClick 返回点击操作相关实现。
	BotClick() B
	// ItemStackOperation 返回物品堆栈操作入口。
	ItemStackOperation() ISO
	// ContainerOpenAndClose 返回容器打开关闭相关实现。
	ContainerOpenAndClose() CO
	// ItemCopy 返回物品复制相关实现。
	ItemCopy() IC
	// ItemTransition 返回物品转移相关实现。
	ItemTransition() IT
	// PlayerKit 返回玩家交互相关实现。
	PlayerKit() PK
}

// Commands 规定命令客户端的标准能力。
type Commands interface {
	// SendSettingsCommand 发送设置类命令。
	SendSettingsCommand(ctx context.Context, command string, dimensional bool) error
	// SendPlayerCommand 发送玩家命令。
	SendPlayerCommand(ctx context.Context, command string) error
	// SendWSCommand 发送 WebSocket 命令。
	SendWSCommand(ctx context.Context, command string) error
	// SendPlayerCommandWithResp 发送玩家命令并返回输出。
	SendPlayerCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error)
	// SendWSCommandWithResp 发送 WebSocket 命令并返回输出。
	SendWSCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error)
	// AwaitChangesGeneral 等待服务端同步完成。
	AwaitChangesGeneral(ctx context.Context) error
	// SendChat 发送聊天消息。
	SendChat(ctx context.Context, content string) error
	// Title 发送标题文本。
	Title(ctx context.Context, message string) error
}

// StructureBackup 规定结构备份客户端的标准能力。
type StructureBackup interface {
	// BackupStructure 备份指定位置的结构。
	BackupStructure(ctx context.Context, pos *protocol_pb.BlockPos) (uuid.UUID, error)
	// BackupOffset 以偏移位置备份结构。
	BackupOffset(ctx context.Context, pos *protocol_pb.BlockPos, offset *protocol_pb.BlockPos) (uuid.UUID, error)
	// RevertStructure 恢复指定结构备份。
	RevertStructure(ctx context.Context, uniqueID uuid.UUID, pos *protocol_pb.BlockPos) error
	// DeleteStructure 删除指定结构备份。
	DeleteStructure(ctx context.Context, uniqueID uuid.UUID) error
}

// Querytarget 规定 querytarget 客户端的标准能力。
type Querytarget interface {
	// DoQuerytarget 执行 querytarget 并返回解析后的目标信息。
	DoQuerytarget(ctx context.Context, target string) ([]*game_interface_pb.TargetQueryingInfo, error)
}

// Movement 规定移动客户端的标准能力。
type Movement interface {
	// StartFlying 让机器人开始飞行。
	StartFlying(ctx context.Context) error
	// StopFlying 让机器人停止飞行。
	StopFlying(ctx context.Context) error
}

// SetBlock 规定方块放置客户端的标准能力。
type SetBlock interface {
	// SetBlock 同步放置方块。
	SetBlock(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error
	// SetBlockAsync 异步放置方块。
	SetBlockAsync(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error
	// SetAnvil 在指定位置放置铁砧并返回方块实体数据。
	SetAnvil(ctx context.Context, pos *protocol_pb.BlockPos, placeBaseBlock bool) (map[string]any, error)
}

// Replaceitem 规定 replaceitem 客户端的标准能力。
type Replaceitem interface {
	// ReplaceitemInInventory 对实体物品栏执行 replaceitem。
	ReplaceitemInInventory(ctx context.Context, target string, path ReplaceitemPath, itemInfo ReplaceitemInfo, method string, blocked bool) error
	// ReplaceitemInContainerAsync 对容器方块异步执行 replaceitem。
	ReplaceitemInContainerAsync(ctx context.Context, blockPos *protocol_pb.BlockPos, itemInfo ReplaceitemInfo, method string) error
}

// BotClick 规定点击操作客户端的标准能力。
type BotClick interface {
	// ChangeSelectedHotbarSlot 切换机器人选中的快捷栏槽位。
	ChangeSelectedHotbarSlot(ctx context.Context, hotbarSlotID SlotID) error
	// ClickBlock 对方块执行点击。
	ClickBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks) error
	// ClickBlockWithPosition 以真实点击位置点击方块。
	ClickBlockWithPosition(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, position *mgl32_pb.Vec3) (bool, error)
	// ClickAir 使用指定快捷栏物品点击空气。
	ClickAir(ctx context.Context, hotbarSlot SlotID, realPosition *mgl32_pb.Vec3) error
	// PlaceBlock 使用指定请求放置方块。
	PlaceBlock(ctx context.Context, request *game_interface_pb.UseItemOnBlocks, blockFace int32) error
	// PlaceBlockHighLevel 使用高级封装放置方块。
	PlaceBlockHighLevel(ctx context.Context, blockPos *protocol_pb.BlockPos, hotBarSlot SlotID, facing uint8) (*protocol_pb.BlockPos, *protocol_pb.BlockPos, *protocol_pb.BlockPos, error)
	// PickBlock 对指定位置执行选取方块。
	PickBlock(ctx context.Context, pos *protocol_pb.BlockPos, assignNBTData bool) (bool, SlotID, error)
}

// ItemStackOperation 规定物品堆栈操作入口的标准能力。
type ItemStackOperation[TX ItemStackTransaction[TX, OP], OP ItemStackOperationData] interface {
	// OpenTransaction 创建一个新的物品堆栈事务。
	OpenTransaction() TX
}

// ItemStackTransaction 规定单个物品堆栈事务的标准能力。
type ItemStackTransaction[TX any, OP ItemStackOperationData] interface {
	// Discord 丢弃当前事务内的全部操作。
	Discord() TX
	// Discard 是 Discord 的可读别名。
	Discard() TX
	// Commit 提交当前事务内的全部操作。
	Commit(ctx context.Context) (success bool, pk *packet_pb.ItemStackRequest, serverResponse []*protocol_pb.ItemStackResponse, err error)
	// CommitItemStackOperations 一次性提交指定操作列表。
	CommitItemStackOperations(ctx context.Context, operations []OP) (*game_interface_pb.CommitTransactionResponse, error)
	// MoveItem 将物品从任意位置移动到任意位置。
	MoveItem(source SlotLocation, destination SlotLocation, count uint8) TX
	// MoveBetweenInventory 在背包槽位之间移动物品。
	MoveBetweenInventory(source SlotID, destination SlotID, count uint8) TX
	// MoveBetweenDynamicContainer 在动态容器槽位之间移动物品。
	MoveBetweenDynamicContainer(dynamicContainerID DynamicContainerID, source SlotID, destination SlotID, count uint8) TX
	// MoveToDynamicContainer 从背包移动物品到动态容器。
	MoveToDynamicContainer(source SlotID, dynamicContainerID DynamicContainerID, destination SlotID, count uint8) TX
	// MoveFromDynamicContainer 从动态容器移动物品到背包。
	MoveFromDynamicContainer(dynamicContainerID DynamicContainerID, source SlotID, destination SlotID, count uint8) TX
	// MoveBetweenContainer 在已打开容器槽位之间移动物品。
	MoveBetweenContainer(source SlotID, destination SlotID, count uint8) TX
	// MoveToContainer 从背包移动物品到已打开容器。
	MoveToContainer(source SlotID, destination SlotID, count uint8) TX
	// MoveToInventory 从已打开容器移动物品到背包。
	MoveToInventory(source SlotID, destination SlotID, count uint8) TX
	// MoveToCraftingTable 从背包移动物品到合成栏。
	MoveToCraftingTable(source SlotID, destination SlotID, count uint8) TX
	// MoveFromCraftingTable 从合成栏移动物品到背包。
	MoveFromCraftingTable(source SlotID, destination SlotID, count uint8) TX
	// SwapItem 交换任意两个位置的物品。
	SwapItem(source SlotLocation, destination SlotLocation) TX
	// SwapBetweenInventory 交换背包槽位物品。
	SwapBetweenInventory(source SlotID, destination SlotID) TX
	// SwapBetweenDynamicContainer 交换动态容器槽位物品。
	SwapBetweenDynamicContainer(dynamicContainerID DynamicContainerID, source SlotID, destination SlotID) TX
	// SwapInventoryBetweenDynamicContainer 交换背包和动态容器槽位物品。
	SwapInventoryBetweenDynamicContainer(source SlotID, dynamicContainerID DynamicContainerID, destination SlotID) TX
	// SwapInventoryBetweenContainer 交换背包和已打开容器槽位物品。
	SwapInventoryBetweenContainer(source SlotID, destination SlotID) TX
	// DropItem 丢弃任意位置的物品。
	DropItem(slot SlotLocation, count uint8) TX
	// DropInventoryItem 丢弃背包槽位物品。
	DropInventoryItem(slot SlotID, count uint8) TX
	// DropDynamicContainerItem 丢弃动态容器槽位物品。
	DropDynamicContainerItem(dynamicContainerID DynamicContainerID, slot SlotID, count uint8) TX
	// DropContainerItem 丢弃已打开容器槽位物品。
	DropContainerItem(slot SlotID, count uint8) TX
	// GetCreativeItem 获取创造物品到任意位置。
	GetCreativeItem(creativeItemNetworkID uint32, slot SlotLocation, count uint8) TX
	// GetCreativeItemToInventory 获取创造物品到背包槽位。
	GetCreativeItemToInventory(creativeItemNetworkID uint32, slot SlotID, count uint8) TX
	// GetCreativeItemToDynamicContainer 获取创造物品到动态容器槽位。
	GetCreativeItemToDynamicContainer(creativeItemNetworkID uint32, dynamicContainerID DynamicContainerID, slot SlotID, count uint8) TX
	// RenameItem 重命名任意位置的物品。
	RenameItem(slot SlotLocation, newName string) TX
	// RenameInventoryItem 重命名背包槽位物品。
	RenameInventoryItem(slot SlotID, newName string) TX
	// Looming 执行织布机物品操作。
	Looming(patternName string, patternSlot SlotLocation, bannerSlot SlotLocation, dyeSlot SlotLocation, resultItem ExpectedNewItem) TX
	// LoomingFromInventory 使用背包槽位执行织布机物品操作。
	LoomingFromInventory(patternName string, patternSlot SlotID, bannerSlot SlotID, dyeSlot SlotID, resultItem ExpectedNewItem) TX
	// Crafting 执行合成物品操作。
	Crafting(recipeNetworkID uint32, resultSlotID SlotID, resultCount uint8, resultItem ExpectedNewItem) TX
	// Trimming 执行锻造台纹饰操作。
	Trimming(trimItemPath SlotLocation, materialPath SlotLocation, templatePath SlotLocation, resultItem ExpectedNewItem) TX
	// TrimmingFromInventory 使用背包槽位执行锻造台纹饰操作。
	TrimmingFromInventory(trimItemSlot SlotID, materialSlot SlotID, templateSlot SlotID, resultItem ExpectedNewItem) TX
	// BeaconPayment 执行信标支付操作。
	BeaconPayment(paymentPath SlotLocation, primaryEffect int32, secondaryEffect int32) TX
	// BeaconPaymentFromInventory 使用背包槽位执行信标支付操作。
	BeaconPaymentFromInventory(paymentSlot SlotID, primaryEffect int32, secondaryEffect int32) TX
}

// ItemStackOperationData 指示所有可提交的物品堆栈操作数据。
type ItemStackOperationData interface {
	// CanInline 指示该物品操作是否可以内联到单个请求中。
	CanInline() bool
	// ID 返回该物品操作的自定义编号。
	ID() uint8
}

// ContainerOpenAndClose 规定容器打开关闭客户端的标准能力。
type ContainerOpenAndClose interface {
	// OpenContainer 打开指定容器。
	OpenContainer(ctx context.Context, container *game_interface_pb.UseItemOnBlocks, changeToTargetSlot bool) (bool, error)
	// OpenInventory 打开玩家背包。
	OpenInventory(ctx context.Context) (bool, error)
	// CloseContainer 关闭当前容器。
	CloseContainer(ctx context.Context) error
}

// ItemCopy 规定物品复制客户端的标准能力。
type ItemCopy interface {
	// CopyItem 将基础物品复制为目标物品。
	CopyItem(ctx context.Context, containerInfo *game_interface_pb.UseItemOnBlocks, baseItems []ItemInfoWithSlot, targetItems []*ItemInfo) error
}

// ItemTransition 规定物品转移客户端的标准能力。
type ItemTransition interface {
	// Transition 在指定窗口间转移物品。
	Transition(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot, srcWindowName WindowName, dstWindowName WindowName) (bool, error)
	// TransitionBetweenInventory 在背包内转移物品。
	TransitionBetweenInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionBetweenContainer 在已打开容器内转移物品。
	TransitionBetweenContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionToContainer 从背包转移物品到已打开容器。
	TransitionToContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionToInventory 从已打开容器转移物品到背包。
	TransitionToInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
}

// PlayerKit 规定玩家交互入口的标准能力。
type PlayerKit[P Player[AB], AB AbilityBuilder[AB]] interface {
	// ListOnlinePlayers 列出在线玩家。
	ListOnlinePlayers(ctx context.Context) ([]P, error)
	// GetPlayerByName 按玩家名查找玩家。
	GetPlayerByName(ctx context.Context, name string) (P, bool, error)
	// GetPlayerByUUIDString 按 UUID 字符串查找玩家。
	GetPlayerByUUIDString(ctx context.Context, uuidString string) (P, bool, error)
	// GetPlayerByUniqueID 按实体唯一 ID 查找玩家。
	GetPlayerByUniqueID(ctx context.Context, id int64) (P, bool, error)
	// GetPlayerByRuntimeID 按实体运行时 ID 查找玩家。
	GetPlayerByRuntimeID(ctx context.Context, id uint64) (P, bool, error)
}

// Player 规定可交互玩家客户端的标准能力。
type Player[AB AbilityBuilder[AB]] interface {
	// SendChat 发送普通聊天文本。
	SendChat(ctx context.Context, msg string) error
	// SendRawChat 发送原始聊天文本。
	SendRawChat(ctx context.Context, rawText string) error
	// SendTitle 发送标题文本。
	SendTitle(ctx context.Context, title string) error
	// SendRawTitle 发送原始标题文本。
	SendRawTitle(ctx context.Context, rawTitle string) error
	// SendSubTitle 发送副标题文本。
	SendSubTitle(ctx context.Context, subtitle string) error
	// SendRawSubTitle 发送原始副标题文本。
	SendRawSubTitle(ctx context.Context, rawSubtitle string) error
	// SendActionBar 发送 actionbar 文本。
	SendActionBar(ctx context.Context, actionBar string) error
	// SendRawActionBar 发送原始 actionbar 文本。
	SendRawActionBar(ctx context.Context, rawText string) error
	// OpenAbility 打开玩家能力修改构造器。
	OpenAbility(ctx context.Context) AB
}

// AbilityBuilder 规定玩家能力修改构造器的标准能力。
type AbilityBuilder[AB any] interface {
	// SetBuildAbility 设置建造能力。
	SetBuildAbility(allow bool) AB
	// SetMineAbility 设置挖掘能力。
	SetMineAbility(allow bool) AB
	// SetDoorsAndSwitchesAbility 设置使用门和开关能力。
	SetDoorsAndSwitchesAbility(allow bool) AB
	// SetOpenContainersAbility 设置打开容器能力。
	SetOpenContainersAbility(allow bool) AB
	// SetAttackPlayersAbility 设置攻击玩家能力。
	SetAttackPlayersAbility(allow bool) AB
	// SetAttackMobsAbility 设置攻击生物能力。
	SetAttackMobsAbility(allow bool) AB
	// SetOperatorCommandsAbility 设置操作员命令能力。
	SetOperatorCommandsAbility(allow bool) AB
	// SetTeleportAbility 设置传送能力。
	SetTeleportAbility(allow bool) AB
	// Commit 提交能力修改。
	Commit() error
}

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

// ContainerManager 规定容器资源客户端的标准能力。
type ContainerManager interface {
	// States 返回已打开容器的状态。
	States(ctx context.Context) (state uint8, err error)
	// ContainerData 返回当前已打开容器数据。
	ContainerData(ctx context.Context) (data *packet_pb.ContainerOpen, containerID ContainerID, existed bool, err error)
}

// PacketListener 规定数据包监听客户端的标准能力。
type PacketListener interface {
	// ListenPacket 注册数据包监听器。
	ListenPacket(ctx context.Context, packetID []uint32, callback func(*packet_pb.Packet, error)) (uniqueID string, err error)
	// DestroyListener 销毁指定数据包监听器。
	DestroyListener(ctx context.Context, uniqueID string) (err error)
}

// ConstantPacket 规定常量数据包资源客户端的标准能力。
type ConstantPacket interface {
	// AllCreativeContent 返回全部创造物品。
	AllCreativeContent(ctx context.Context) (items []*protocol_pb.CreativeItem, err error)
	// CreativeItemByCNI 按创造物品网络 ID 查询创造物品。
	CreativeItemByCNI(ctx context.Context, creativeNetworkID uint32) (item *protocol_pb.CreativeItem, existed bool, err error)
	// CreativeItemByNI 按物品网络 ID 查询创造物品。
	CreativeItemByNI(ctx context.Context, networkID int32) (items []*protocol_pb.CreativeItem, err error)
	// CreativeItemByName 按名称查询创造物品。
	CreativeItemByName(ctx context.Context, name string) (items []*protocol_pb.CreativeItem, err error)
	// AllAvailableItems 返回全部可用物品。
	AllAvailableItems(ctx context.Context) (items []*protocol_pb.ItemEntry, err error)
	// ItemByNetworkID 按网络 ID 查询物品。
	ItemByNetworkID(ctx context.Context, networkID int32) (item *protocol_pb.ItemEntry, existed bool, err error)
	// ItemByName 按名称查询物品。
	ItemByName(ctx context.Context, name string) (item *protocol_pb.ItemEntry, existed bool, err error)
	// ItemNameByNetworkID 按网络 ID 查询物品名。
	ItemNameByNetworkID(ctx context.Context, networkID int32) (name string, existed bool, err error)
	// AllCommandItems 返回全部可通过命令获取的物品名。
	AllCommandItems(ctx context.Context) (items []string, err error)
	// ItemCanGetByCommand 判断物品是否可通过命令获取。
	ItemCanGetByCommand(ctx context.Context, name string) (canGet bool, err error)
	// TrimRecipeNetworkID 返回纹饰配方网络 ID。
	TrimRecipeNetworkID(ctx context.Context) (networkID uint32, err error)
}

// UQHolder 规定由数据包维护的轻量状态集合标准形状。
type UQHolder[
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
	// Bot 返回机器人自身状态。
	Bot() B
	// Players 返回在线玩家状态索引。
	Players() PS
	// World 返回世界状态。
	World() W
	// Entities 返回实体状态索引。
	Entities() ES
}

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

// Players 规定在线玩家索引查询能力。
type Players[P UQPlayer[PA], PA PlayerAbilities] interface {
	// GetOnline 返回当前在线玩家列表。
	GetOnline(ctx context.Context) (players []P, err error)
	// GetByUUIDString 按 UUID 字符串查找玩家。
	GetByUUIDString(ctx context.Context, id string) (player P, existed bool, err error)
	// GetByName 按玩家名查找玩家。
	GetByName(ctx context.Context, name string) (player P, existed bool, err error)
	// GetByUniqueID 按实体唯一 ID 查找玩家。
	GetByUniqueID(ctx context.Context, id int64) (player P, existed bool, err error)
	// GetByRuntimeID 按实体运行时 ID 查找玩家。
	GetByRuntimeID(ctx context.Context, id uint64) (player P, existed bool, err error)
}

// UQPlayer 规定玩家状态只读视图能力。
type UQPlayer[PA PlayerAbilities] interface {
	// GetUUIDString 返回玩家 UUID 字符串。
	GetUUIDString(ctx context.Context) (uuid string, ok bool, err error)
	// GetName 返回玩家名称。
	GetName(ctx context.Context) (name string, ok bool, err error)
	// GetXUID 返回玩家 XUID。
	GetXUID(ctx context.Context) (xuid string, ok bool, err error)
	// GetEntityUniqueID 返回玩家实体唯一 ID。
	GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error)
	// GetEntityRuntimeID 返回玩家实体运行时 ID。
	GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error)
	// GetNeteaseUID 返回玩家网易 UID。
	GetNeteaseUID(ctx context.Context) (neteaseUID int64, ok bool, err error)
	// GetLoginTime 返回玩家登录时间。
	GetLoginTime(ctx context.Context) (loginTime time.Time, ok bool, err error)
	// GetPlatformChatID 返回玩家平台聊天 ID。
	GetPlatformChatID(ctx context.Context) (platformChatID string, ok bool, err error)
	// GetBuildPlatform 返回玩家构建平台。
	GetBuildPlatform(ctx context.Context) (buildPlatform int32, ok bool, err error)
	// GetSkinID 返回玩家皮肤 ID。
	GetSkinID(ctx context.Context) (skinID string, ok bool, err error)
	// GetDeviceID 返回玩家设备 ID。
	GetDeviceID(ctx context.Context) (deviceID string, ok bool, err error)
	// GetEntityMetadata 返回玩家实体元数据。
	GetEntityMetadata(ctx context.Context) (metadata *protocol_pb.EntityMetadata, ok bool, err error)
	// GetPosition 返回玩家位置。
	GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error)
	// GetPitch 返回玩家俯仰角。
	GetPitch(ctx context.Context) (pitch float32, ok bool, err error)
	// GetYaw 返回玩家偏航角。
	GetYaw(ctx context.Context) (yaw float32, ok bool, err error)
	// GetHeadYaw 返回玩家头部偏航角。
	GetHeadYaw(ctx context.Context) (headYaw float32, ok bool, err error)
	// GetMoveMode 返回玩家移动模式。
	GetMoveMode(ctx context.Context) (moveMode byte, ok bool, err error)
	// GetOnGround 返回玩家是否在地面上。
	GetOnGround(ctx context.Context) (onGround bool, ok bool, err error)
	// GetRiddenEntityRuntimeID 返回骑乘实体运行时 ID。
	GetRiddenEntityRuntimeID(ctx context.Context) (riddenEntityRuntimeID uint64, ok bool, err error)
	// GetTick 返回玩家状态最后更新时间刻。
	GetTick(ctx context.Context) (tick uint64, ok bool, err error)
	// GetAbilities 返回玩家权限和能力查询接口。
	GetAbilities() PA
	// GetOnline 返回玩家是否在线。
	GetOnline(ctx context.Context) (online bool, ok bool, err error)
}

// PlayerAbilities 规定玩家权限和能力查询能力。
type PlayerAbilities interface {
	// GetCommandPermissions 返回命令权限等级。
	GetCommandPermissions(ctx context.Context) (commandPermissions byte, ok bool, err error)
	// GetPlayerPermissions 返回玩家权限等级。
	GetPlayerPermissions(ctx context.Context) (playerPermissions byte, ok bool, err error)
	// GetFlySpeed 返回飞行速度。
	GetFlySpeed(ctx context.Context) (flySpeed float32, ok bool, err error)
	// GetWalkSpeed 返回行走速度。
	GetWalkSpeed(ctx context.Context) (walkSpeed float32, ok bool, err error)
	// GetCanBuild 返回是否可以建造。
	GetCanBuild(ctx context.Context) (canBuild bool, ok bool, err error)
	// GetCanMine 返回是否可以挖掘。
	GetCanMine(ctx context.Context) (canMine bool, ok bool, err error)
	// GetCanUseDoorsAndSwitches 返回是否可以使用门和开关。
	GetCanUseDoorsAndSwitches(ctx context.Context) (canUseDoorsAndSwitches bool, ok bool, err error)
	// GetCanOpenContainers 返回是否可以打开容器。
	GetCanOpenContainers(ctx context.Context) (canOpenContainers bool, ok bool, err error)
	// GetCanAttackPlayers 返回是否可以攻击玩家。
	GetCanAttackPlayers(ctx context.Context) (canAttackPlayers bool, ok bool, err error)
	// GetCanAttackMobs 返回是否可以攻击生物。
	GetCanAttackMobs(ctx context.Context) (canAttackMobs bool, ok bool, err error)
	// GetCanUseOperatorCommands 返回是否可以使用操作员命令。
	GetCanUseOperatorCommands(ctx context.Context) (canUseOperatorCommands bool, ok bool, err error)
	// GetCanTeleport 返回是否可以传送。
	GetCanTeleport(ctx context.Context) (canTeleport bool, ok bool, err error)
	// GetIsInvulnerable 返回是否无敌。
	GetIsInvulnerable(ctx context.Context) (isInvulnerable bool, ok bool, err error)
	// GetIsFlying 返回是否正在飞行。
	GetIsFlying(ctx context.Context) (isFlying bool, ok bool, err error)
	// GetCanFly 返回是否可以飞行。
	GetCanFly(ctx context.Context) (canFly bool, ok bool, err error)
	// GetCanInstantBuild 返回是否可以瞬间建造。
	GetCanInstantBuild(ctx context.Context) (canInstantBuild bool, ok bool, err error)
	// GetCanUseLightning 返回是否可以使用闪电。
	GetCanUseLightning(ctx context.Context) (canUseLightning bool, ok bool, err error)
	// GetIsMuted 返回是否被禁言。
	GetIsMuted(ctx context.Context) (isMuted bool, ok bool, err error)
	// GetIsWorldBuilder 返回是否为世界建造者。
	GetIsWorldBuilder(ctx context.Context) (isWorldBuilder bool, ok bool, err error)
	// GetHasNoClip 返回是否启用无碰撞。
	GetHasNoClip(ctx context.Context) (hasNoClip bool, ok bool, err error)
	// GetIsPrivilegedBuilder 返回是否为特权建造者。
	GetIsPrivilegedBuilder(ctx context.Context) (isPrivilegedBuilder bool, ok bool, err error)
}

// World 规定世界状态查询能力。
type World[GR GameRule] interface {
	// GetCurrentTick 返回当前 tick。
	GetCurrentTick(ctx context.Context) (currentTick int64, ok bool, err error)
	// GetSyncRatio 返回同步比例。
	GetSyncRatio(ctx context.Context) (syncRatio float32, ok bool, err error)
	// GetTime 返回世界时间。
	GetTime(ctx context.Context) (time int32, ok bool, err error)
	// GetDayTime 返回日间时间。
	GetDayTime(ctx context.Context) (dayTime int32, ok bool, err error)
	// GetDayTimePercent 返回日间时间百分比。
	GetDayTimePercent(ctx context.Context) (dayTimePercent float32, ok bool, err error)
	// GetDifficulty 返回世界难度。
	GetDifficulty(ctx context.Context) (difficulty uint32, ok bool, err error)
	// GetWorldGameMode 返回世界默认游戏模式。
	GetWorldGameMode(ctx context.Context) (worldGameMode int32, ok bool, err error)
	// GetGameRuleNames 返回全部游戏规则名。
	GetGameRuleNames(ctx context.Context) (names []string, err error)
	// GetGameRule 按名称返回游戏规则。
	GetGameRule(ctx context.Context, name string) (gameRule GR, existed bool, err error)
}

// GameRule 规定游戏规则只读能力。
type GameRule interface {
	// GetCanBeModifiedByPlayer 返回该规则是否可由玩家修改。
	GetCanBeModifiedByPlayer() (canBeModifiedByPlayer bool)
	// GetValue 返回规则值。
	GetValue() (value string)
}

// Entities 规定实体状态索引查询能力。
type Entities[E Entity[ME], ME MobEffect] interface {
	// GetByRuntimeID 按运行时 ID 查询实体。
	GetByRuntimeID(ctx context.Context, runtimeID uint64) (entity E, existed bool, err error)
	// GetHealth 按运行时 ID 查询实体生命值。
	GetHealth(ctx context.Context, runtimeID uint64) (health float32, ok bool, err error)
	// GetEffectTypes 按运行时 ID 查询实体效果类型。
	GetEffectTypes(ctx context.Context, runtimeID uint64) (effectTypes []int32, existed bool, err error)
}

// Entity 规定单个实体状态查询能力。
type Entity[ME MobEffect] interface {
	// GetRuntimeID 返回实体运行时 ID。
	GetRuntimeID(ctx context.Context) (runtimeID uint64, ok bool, err error)
	// GetUniqueID 返回实体唯一 ID。
	GetUniqueID(ctx context.Context) (uniqueID int64, ok bool, err error)
	// GetHealth 返回实体生命值。
	GetHealth(ctx context.Context) (health float32, ok bool, err error)
	// GetEffectTypes 返回实体效果类型。
	GetEffectTypes(ctx context.Context) (effectTypes []int32, existed bool, err error)
	// GetEffect 返回指定效果类型的药水效果。
	GetEffect(ctx context.Context, effectType int32) (mobEffect ME, existed bool, err error)
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

// ConnectionState 描述当前受管连接是否可用。
type ConnectionState struct {
	// Connected 表示当前连接是否可用。
	Connected bool
	// CloseReason 是连接关闭原因。
	CloseReason string
}

// ClosedEvent 是连接关闭事件。
type ClosedEvent struct {
	// Reason 是连接关闭原因。
	Reason string
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

// ContainerID 是容器 ID。
type ContainerID uint8

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
	Slot SlotID
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
	Slot SlotID
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

// ExpectedNewItem 描述单个物品堆栈在经历一次物品堆栈操作后，
// 其最终应当拥有的一些数据信息。
type ExpectedNewItem struct {
	// ItemType 描述物品基本类型字段的期望更新。
	ItemType ItemNewType
	// BlockRuntimeID 描述物品方块运行时 ID 的期望更新。
	BlockRuntimeID ItemNewBlockRuntimeID
	// NBT 描述物品 NBT 数据的期望更新。
	NBT ItemNewNBTData
	// Component 描述物品 Legacy 组件的期望更新。
	Component ItemNewComponent
}

// ItemNewType 描述物品的一些基本字段应如何更新。
type ItemNewType struct {
	// UseNetworkID 指示是否更新物品网络 ID。
	UseNetworkID bool
	// NetworkID 是新的物品网络 ID。
	NetworkID int32
	// UseMetadata 指示是否更新物品元数据。
	UseMetadata bool
	// Metadata 是新的物品元数据。
	Metadata uint32
}

// ItemNewBlockRuntimeID 描述物品对应的方块运行时数据应该如何更新。
type ItemNewBlockRuntimeID struct {
	// UseBlockRuntimeID 指示是否更新方块运行时 ID。
	UseBlockRuntimeID bool
	// BlockRuntimeID 是新的方块运行时 ID。
	BlockRuntimeID int32
}

// ItemNewNBTData 描述物品的新 NBT 字段如何更新。
type ItemNewNBTData struct {
	// UseNBTData 指示是否更新 NBT 数据。
	UseNBTData bool
	// UseOriginDamage 指示是否沿用原始 damage。
	UseOriginDamage bool
	// NBTData 是新的 NBT 数据。
	NBTData map[string]any
	// ChangeRepairCost 指示是否调整 RepairCost。
	ChangeRepairCost bool
	// RepairCostDelta 是 RepairCost 的变化量。
	RepairCostDelta int32
	// ChangeDamage 指示是否调整 Damage。
	ChangeDamage bool
	// DamageDelta 是 Damage 的变化量。
	DamageDelta int32
}

// ItemNewComponent 描述物品的 Legacy 物品组件应当如何更新。
type ItemNewComponent struct {
	// UseCanPlaceOn 指示是否更新可放置方块列表。
	UseCanPlaceOn bool
	// CanPlaceOn 是新的可放置方块列表。
	CanPlaceOn []string
	// UseCanDestroy 指示是否更新可破坏方块列表。
	UseCanDestroy bool
	// CanDestroy 是新的可破坏方块列表。
	CanDestroy []string
}

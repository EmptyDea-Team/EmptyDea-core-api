package game_interface

import (
	item_stack_operation "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_operation"
	item_stack_transaction "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_transaction"
)

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
	TX item_stack_transaction.ItemStackTransaction[TX, OP],
	OP item_stack_operation.ItemStackOperation,
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

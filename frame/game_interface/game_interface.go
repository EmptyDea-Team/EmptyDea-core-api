package game_interface

// GameInterface 规定游戏交互层客户端集合的标准形状。
type GameInterface interface {
	// Commands 返回命令相关实现。
	Commands() Commands
	// StructureBackup 返回结构备份相关实现。
	StructureBackup() StructureBackup
	// Querytarget 返回 querytarget 相关实现。
	Querytarget() Querytarget
	// Movement 返回移动控制相关实现。
	Movement() Movement
	// SetBlock 返回方块放置相关实现。
	SetBlock() SetBlock
	// Replaceitem 返回 replaceitem 相关实现。
	Replaceitem() Replaceitem
	// BotClick 返回点击操作相关实现。
	BotClick() BotClick
	// ItemStackOperation 返回物品堆栈操作入口。
	ItemStackOperation() ItemStackOperation
	// ContainerOpenAndClose 返回容器打开关闭相关实现。
	ContainerOpenAndClose() ContainerOpenAndClose
	// ItemCopy 返回物品复制相关实现。
	ItemCopy() ItemCopy
	// ItemTransition 返回物品转移相关实现。
	ItemTransition() ItemTransition
	// PlayerKit 返回玩家交互相关实现。
	PlayerKit() PlayerKit
}

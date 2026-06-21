package game_interface

import (
	item_stack_operation "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_operation"
	item_stack_transaction "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface/item_stack_transaction"
)

// ItemStackOperation 规定物品堆栈操作入口的标准能力。
type ItemStackOperation[TX item_stack_transaction.ItemStackTransaction[TX, OP], OP item_stack_operation.ItemStackOperation] interface {
	// OpenTransaction 创建一个新的物品堆栈事务。
	OpenTransaction() TX
}

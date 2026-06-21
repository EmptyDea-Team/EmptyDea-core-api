package item_stack_operation

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

// ItemStackOperation 指示所有可提交的物品堆栈操作数据。
type ItemStackOperation interface {
	// CanInline 指示该物品操作是否可以内联到单个请求中。
	CanInline() bool
	// ID 返回该物品操作的自定义编号。
	ID() uint8
}

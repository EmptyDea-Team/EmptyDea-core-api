package define

// ItemStackOperation 规定物品堆栈操作入口的标准能力。
type ItemStackOperation[TX ItemStackTransaction[TX, OP], OP ItemStackOperationData] interface {
	// OpenTransaction 创建一个新的物品堆栈事务。
	OpenTransaction() TX
}

// ItemStackOperationData 指示所有可提交的物品堆栈操作数据。
type ItemStackOperationData interface {
	// CanInline 指示该物品操作是否可以内联到单个请求中。
	CanInline() bool
	// ID 返回该物品操作的自定义编号。
	ID() uint8
}

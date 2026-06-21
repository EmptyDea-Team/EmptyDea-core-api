package item_stack_operation

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

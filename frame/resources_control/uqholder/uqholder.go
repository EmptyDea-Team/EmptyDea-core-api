package uqholder

// UQHolder 规定由数据包维护的轻量状态集合标准形状。
type UQHolder interface {
	// Bot 返回机器人自身状态。
	Bot() Bot
	// Players 返回在线玩家状态索引。
	Players() Players
	// World 返回世界状态。
	World() World
	// Entities 返回实体状态索引。
	Entities() Entities
}

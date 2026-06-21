package uqholder

// UQHolder 规定由数据包维护的轻量状态集合标准形状。
type UQHolder[
	B Bot,
	PS Players[P, PA],
	P Player[PA],
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

package game_interface

import "context"

// PlayerKit 规定玩家交互入口的标准能力。
type PlayerKit interface {
	// ListOnlinePlayers 列出在线玩家。
	ListOnlinePlayers(ctx context.Context) ([]Player, error)
	// GetPlayerByName 按玩家名查找玩家。
	GetPlayerByName(ctx context.Context, name string) (Player, bool, error)
	// GetPlayerByUUIDString 按 UUID 字符串查找玩家。
	GetPlayerByUUIDString(ctx context.Context, uuidString string) (Player, bool, error)
	// GetPlayerByUniqueID 按实体唯一 ID 查找玩家。
	GetPlayerByUniqueID(ctx context.Context, id int64) (Player, bool, error)
	// GetPlayerByRuntimeID 按实体运行时 ID 查找玩家。
	GetPlayerByRuntimeID(ctx context.Context, id uint64) (Player, bool, error)
}

// Player 规定可交互玩家客户端的标准能力。
type Player interface {
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
	OpenAbility(ctx context.Context) AbilityBuilder
}

// AbilityBuilder 规定玩家能力修改构造器的标准能力。
type AbilityBuilder interface {
	// SetBuildAbility 设置建造能力。
	SetBuildAbility(allow bool) AbilityBuilder
	// SetMineAbility 设置挖掘能力。
	SetMineAbility(allow bool) AbilityBuilder
	// SetDoorsAndSwitchesAbility 设置使用门和开关能力。
	SetDoorsAndSwitchesAbility(allow bool) AbilityBuilder
	// SetOpenContainersAbility 设置打开容器能力。
	SetOpenContainersAbility(allow bool) AbilityBuilder
	// SetAttackPlayersAbility 设置攻击玩家能力。
	SetAttackPlayersAbility(allow bool) AbilityBuilder
	// SetAttackMobsAbility 设置攻击生物能力。
	SetAttackMobsAbility(allow bool) AbilityBuilder
	// SetOperatorCommandsAbility 设置操作员命令能力。
	SetOperatorCommandsAbility(allow bool) AbilityBuilder
	// SetTeleportAbility 设置传送能力。
	SetTeleportAbility(allow bool) AbilityBuilder
	// Commit 提交能力修改。
	Commit() error
}

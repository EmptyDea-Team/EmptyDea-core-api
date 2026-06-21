package uqholder

import "context"

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

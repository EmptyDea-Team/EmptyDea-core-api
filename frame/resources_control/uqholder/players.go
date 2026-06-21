package uqholder

import "context"

// Players 规定在线玩家索引查询能力。
type Players interface {
	// GetOnline 返回当前在线玩家列表。
	GetOnline(ctx context.Context) (players []Player, err error)
	// GetByUUIDString 按 UUID 字符串查找玩家。
	GetByUUIDString(ctx context.Context, id string) (player Player, existed bool, err error)
	// GetByName 按玩家名查找玩家。
	GetByName(ctx context.Context, name string) (player Player, existed bool, err error)
	// GetByUniqueID 按实体唯一 ID 查找玩家。
	GetByUniqueID(ctx context.Context, id int64) (player Player, existed bool, err error)
	// GetByRuntimeID 按实体运行时 ID 查找玩家。
	GetByRuntimeID(ctx context.Context, id uint64) (player Player, existed bool, err error)
}

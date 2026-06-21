package define

import (
	"context"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// ContainerOpenAndClose 规定容器打开关闭客户端的标准能力。
type ContainerOpenAndClose interface {
	// OpenContainer 打开指定容器。
	OpenContainer(ctx context.Context, container *game_interface_pb.UseItemOnBlocks, changeToTargetSlot bool) (bool, error)
	// OpenInventory 打开玩家背包。
	OpenInventory(ctx context.Context) (bool, error)
	// CloseContainer 关闭当前容器。
	CloseContainer(ctx context.Context) error
}

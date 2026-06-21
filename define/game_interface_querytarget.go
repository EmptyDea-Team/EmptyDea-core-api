package define

import (
	"context"
	game_interface_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/game_control/game_interface"
)

// Querytarget 规定 querytarget 客户端的标准能力。
type Querytarget interface {
	// DoQuerytarget 执行 querytarget 并返回解析后的目标信息。
	DoQuerytarget(ctx context.Context, target string) ([]*game_interface_pb.TargetQueryingInfo, error)
}

package game_interface

import (
	"context"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// SetBlock 规定方块放置客户端的标准能力。
type SetBlock interface {
	// SetBlock 同步放置方块。
	SetBlock(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error
	// SetBlockAsync 异步放置方块。
	SetBlockAsync(ctx context.Context, pos *protocol_pb.BlockPos, name string, states string) error
	// SetAnvil 在指定位置放置铁砧并返回方块实体数据。
	SetAnvil(ctx context.Context, pos *protocol_pb.BlockPos, placeBaseBlock bool) (map[string]any, error)
}

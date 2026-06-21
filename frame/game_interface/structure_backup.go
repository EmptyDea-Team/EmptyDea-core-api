package game_interface

import (
	"context"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	"github.com/google/uuid"
)

// StructureBackup 规定结构备份客户端的标准能力。
type StructureBackup interface {
	// BackupStructure 备份指定位置的结构。
	BackupStructure(ctx context.Context, pos *protocol_pb.BlockPos) (uuid.UUID, error)
	// BackupOffset 以偏移位置备份结构。
	BackupOffset(ctx context.Context, pos *protocol_pb.BlockPos, offset *protocol_pb.BlockPos) (uuid.UUID, error)
	// RevertStructure 恢复指定结构备份。
	RevertStructure(ctx context.Context, uniqueID uuid.UUID, pos *protocol_pb.BlockPos) error
	// DeleteStructure 删除指定结构备份。
	DeleteStructure(ctx context.Context, uniqueID uuid.UUID) error
}

package define

import (
	"context"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// PacketListener 规定数据包监听客户端的标准能力。
type PacketListener interface {
	// ListenPacket 注册数据包监听器。
	ListenPacket(ctx context.Context, packetID []uint32, callback func(*packet_pb.Packet, error)) (uniqueID string, err error)
	// DestroyListener 销毁指定数据包监听器。
	DestroyListener(ctx context.Context, uniqueID string) (err error)
}

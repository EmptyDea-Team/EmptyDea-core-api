package define

import (
	"context"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// ContainerManager 规定容器资源客户端的标准能力。
type ContainerManager interface {
	// States 返回已打开容器的状态。
	States(ctx context.Context) (state uint8, err error)
	// ContainerData 返回当前已打开容器数据。
	ContainerData(ctx context.Context) (data *packet_pb.ContainerOpen, containerID ContainerID, existed bool, err error)
}

// ContainerID 是容器 ID。
type ContainerID uint8

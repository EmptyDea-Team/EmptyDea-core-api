package define

import (
	"context"
	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// Commands 规定命令客户端的标准能力。
type Commands interface {
	// SendSettingsCommand 发送设置类命令。
	SendSettingsCommand(ctx context.Context, command string, dimensional bool) error
	// SendPlayerCommand 发送玩家命令。
	SendPlayerCommand(ctx context.Context, command string) error
	// SendWSCommand 发送 WebSocket 命令。
	SendWSCommand(ctx context.Context, command string) error
	// SendPlayerCommandWithResp 发送玩家命令并返回输出。
	SendPlayerCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error)
	// SendWSCommandWithResp 发送 WebSocket 命令并返回输出。
	SendWSCommandWithResp(ctx context.Context, command string) (*packet_pb.CommandOutput, error)
	// AwaitChangesGeneral 等待服务端同步完成。
	AwaitChangesGeneral(ctx context.Context) error
	// SendChat 发送聊天消息。
	SendChat(ctx context.Context, content string) error
	// Title 发送标题文本。
	Title(ctx context.Context, message string) error
}

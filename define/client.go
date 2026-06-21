package define

import (
	"context"

	packet_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/packet"
)

// Client 规定 EmptyDea Core 客户端入口的标准形状。
type Client[F Frame, R Resources, C Commands, G GameInterface[C]] interface {
	// Frame 返回框架层客户端。
	Frame() F
	// Resources 返回资源层客户端集合。
	Resources() R
	// GameInterface 返回游戏交互层客户端集合。
	GameInterface() G
	// Close 关闭客户端持有的底层连接。
	Close() error
}

// Frame 规定框架层客户端的标准能力。
type Frame interface {
	// StartConnection 使用配置启动 Minecraft 连接。
	StartConnection(ctx context.Context, config FrameConfig) (string, error)
	// StopConnection 主动关闭当前连接。
	StopConnection(ctx context.Context) error
	// GetConnectionState 查询当前连接状态。
	GetConnectionState(ctx context.Context) (ConnectionState, error)
	// Ping 检查 API 服务是否可响应。
	Ping(ctx context.Context) (bool, error)
	// WatchClosed 监听连接关闭事件。
	WatchClosed(ctx context.Context, callback func(ClosedEvent, error)) error
}

// GameInterface 规定游戏交互层客户端集合的标准形状。
type GameInterface[C Commands] interface {
	// Commands 返回命令相关实现。
	Commands() C
}

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

// Resources 规定资源层客户端集合的标准能力。
type Resources interface {
	// BotInfo 返回机器人基础信息。
	BotInfo(ctx context.Context) (BotInfo, error)
	// WritePacket 向服务端发送数据包。
	WritePacket(ctx context.Context, p *packet_pb.Packet) error
}

// ConnectionState 描述当前受管连接是否可用。
type ConnectionState struct {
	// Connected 表示当前连接是否可用。
	Connected bool
	// CloseReason 是连接关闭原因。
	CloseReason string
}

// ClosedEvent 是连接关闭事件。
type ClosedEvent struct {
	// Reason 是连接关闭原因。
	Reason string
}

// BotInfo 是当前机器人基础登录信息。
type BotInfo struct {
	// BotName 是机器人玩家名。
	BotName string
	// XUID 是机器人 Xbox 用户 ID。
	XUID string
	// EntityUniqueID 是机器人实体唯一 ID。
	EntityUniqueID int64
	// EntityRuntimeID 是机器人实体运行时 ID。
	EntityRuntimeID uint64
}

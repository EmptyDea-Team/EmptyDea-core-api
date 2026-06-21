package frame

import "context"

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

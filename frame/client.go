package frame

import (
	game_interface "github.com/EmptyDea-Team/EmptyDea-core-api/frame/game_interface"
	resources_control "github.com/EmptyDea-Team/EmptyDea-core-api/frame/resources_control"
	"google.golang.org/grpc"
)

// Client 规定 EmptyDea Core 客户端入口的标准形状。
type Client interface {
	// Conn 返回底层 gRPC 连接；非 gRPC 连接实现可以返回 nil。
	Conn() *grpc.ClientConn
	// Frame 返回框架层客户端。
	Frame() Frame
	// Resources 返回资源层客户端集合。
	Resources() resources_control.Resources
	// GameInterface 返回游戏交互层客户端集合。
	GameInterface() game_interface.GameInterface
	// Close 关闭客户端持有的底层连接。
	Close() error
}

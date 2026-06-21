package define

import "context"

// Movement 规定移动客户端的标准能力。
type Movement interface {
	// StartFlying 让机器人开始飞行。
	StartFlying(ctx context.Context) error
	// StopFlying 让机器人停止飞行。
	StopFlying(ctx context.Context) error
}

package define

// FrameConfig 是启动 Minecraft 连接所需的配置。
type FrameConfig struct {
	// AuthServer 是登录认证服务地址。
	AuthServer string
	// UserToken 是用于认证当前用户的令牌。
	UserToken string
	// ServerCode 是目标租赁服或服务器的连接代码。
	ServerCode string
	// ServerPassword 是目标服务器的连接密码。
	ServerPassword string
}

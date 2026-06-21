package uqholder

import (
	"context"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
	mgl32_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol/mgl32"
	"time"
)

// Player 规定玩家状态只读视图能力。
type Player[PA PlayerAbilities] interface {
	// GetUUIDString 返回玩家 UUID 字符串。
	GetUUIDString(ctx context.Context) (uuid string, ok bool, err error)
	// GetName 返回玩家名称。
	GetName(ctx context.Context) (name string, ok bool, err error)
	// GetXUID 返回玩家 XUID。
	GetXUID(ctx context.Context) (xuid string, ok bool, err error)
	// GetEntityUniqueID 返回玩家实体唯一 ID。
	GetEntityUniqueID(ctx context.Context) (entityUniqueID int64, ok bool, err error)
	// GetEntityRuntimeID 返回玩家实体运行时 ID。
	GetEntityRuntimeID(ctx context.Context) (entityRuntimeID uint64, ok bool, err error)
	// GetNeteaseUID 返回玩家网易 UID。
	GetNeteaseUID(ctx context.Context) (neteaseUID int64, ok bool, err error)
	// GetLoginTime 返回玩家登录时间。
	GetLoginTime(ctx context.Context) (loginTime time.Time, ok bool, err error)
	// GetPlatformChatID 返回玩家平台聊天 ID。
	GetPlatformChatID(ctx context.Context) (platformChatID string, ok bool, err error)
	// GetBuildPlatform 返回玩家构建平台。
	GetBuildPlatform(ctx context.Context) (buildPlatform int32, ok bool, err error)
	// GetSkinID 返回玩家皮肤 ID。
	GetSkinID(ctx context.Context) (skinID string, ok bool, err error)
	// GetDeviceID 返回玩家设备 ID。
	GetDeviceID(ctx context.Context) (deviceID string, ok bool, err error)
	// GetEntityMetadata 返回玩家实体元数据。
	GetEntityMetadata(ctx context.Context) (metadata *protocol_pb.EntityMetadata, ok bool, err error)
	// GetPosition 返回玩家位置。
	GetPosition(ctx context.Context) (pos *mgl32_pb.Vec3, ok bool, err error)
	// GetPitch 返回玩家俯仰角。
	GetPitch(ctx context.Context) (pitch float32, ok bool, err error)
	// GetYaw 返回玩家偏航角。
	GetYaw(ctx context.Context) (yaw float32, ok bool, err error)
	// GetHeadYaw 返回玩家头部偏航角。
	GetHeadYaw(ctx context.Context) (headYaw float32, ok bool, err error)
	// GetMoveMode 返回玩家移动模式。
	GetMoveMode(ctx context.Context) (moveMode byte, ok bool, err error)
	// GetOnGround 返回玩家是否在地面上。
	GetOnGround(ctx context.Context) (onGround bool, ok bool, err error)
	// GetRiddenEntityRuntimeID 返回骑乘实体运行时 ID。
	GetRiddenEntityRuntimeID(ctx context.Context) (riddenEntityRuntimeID uint64, ok bool, err error)
	// GetTick 返回玩家状态最后更新时间刻。
	GetTick(ctx context.Context) (tick uint64, ok bool, err error)
	// GetAbilities 返回玩家权限和能力查询接口。
	GetAbilities() PA
	// GetOnline 返回玩家是否在线。
	GetOnline(ctx context.Context) (online bool, ok bool, err error)
}

// PlayerAbilities 规定玩家权限和能力查询能力。
type PlayerAbilities interface {
	// GetCommandPermissions 返回命令权限等级。
	GetCommandPermissions(ctx context.Context) (commandPermissions byte, ok bool, err error)
	// GetPlayerPermissions 返回玩家权限等级。
	GetPlayerPermissions(ctx context.Context) (playerPermissions byte, ok bool, err error)
	// GetFlySpeed 返回飞行速度。
	GetFlySpeed(ctx context.Context) (flySpeed float32, ok bool, err error)
	// GetWalkSpeed 返回行走速度。
	GetWalkSpeed(ctx context.Context) (walkSpeed float32, ok bool, err error)
	// GetCanBuild 返回是否可以建造。
	GetCanBuild(ctx context.Context) (canBuild bool, ok bool, err error)
	// GetCanMine 返回是否可以挖掘。
	GetCanMine(ctx context.Context) (canMine bool, ok bool, err error)
	// GetCanUseDoorsAndSwitches 返回是否可以使用门和开关。
	GetCanUseDoorsAndSwitches(ctx context.Context) (canUseDoorsAndSwitches bool, ok bool, err error)
	// GetCanOpenContainers 返回是否可以打开容器。
	GetCanOpenContainers(ctx context.Context) (canOpenContainers bool, ok bool, err error)
	// GetCanAttackPlayers 返回是否可以攻击玩家。
	GetCanAttackPlayers(ctx context.Context) (canAttackPlayers bool, ok bool, err error)
	// GetCanAttackMobs 返回是否可以攻击生物。
	GetCanAttackMobs(ctx context.Context) (canAttackMobs bool, ok bool, err error)
	// GetCanUseOperatorCommands 返回是否可以使用操作员命令。
	GetCanUseOperatorCommands(ctx context.Context) (canUseOperatorCommands bool, ok bool, err error)
	// GetCanTeleport 返回是否可以传送。
	GetCanTeleport(ctx context.Context) (canTeleport bool, ok bool, err error)
	// GetIsInvulnerable 返回是否无敌。
	GetIsInvulnerable(ctx context.Context) (isInvulnerable bool, ok bool, err error)
	// GetIsFlying 返回是否正在飞行。
	GetIsFlying(ctx context.Context) (isFlying bool, ok bool, err error)
	// GetCanFly 返回是否可以飞行。
	GetCanFly(ctx context.Context) (canFly bool, ok bool, err error)
	// GetCanInstantBuild 返回是否可以瞬间建造。
	GetCanInstantBuild(ctx context.Context) (canInstantBuild bool, ok bool, err error)
	// GetCanUseLightning 返回是否可以使用闪电。
	GetCanUseLightning(ctx context.Context) (canUseLightning bool, ok bool, err error)
	// GetIsMuted 返回是否被禁言。
	GetIsMuted(ctx context.Context) (isMuted bool, ok bool, err error)
	// GetIsWorldBuilder 返回是否为世界建造者。
	GetIsWorldBuilder(ctx context.Context) (isWorldBuilder bool, ok bool, err error)
	// GetHasNoClip 返回是否启用无碰撞。
	GetHasNoClip(ctx context.Context) (hasNoClip bool, ok bool, err error)
	// GetIsPrivilegedBuilder 返回是否为特权建造者。
	GetIsPrivilegedBuilder(ctx context.Context) (isPrivilegedBuilder bool, ok bool, err error)
}

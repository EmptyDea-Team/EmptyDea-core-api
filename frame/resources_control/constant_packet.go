package resources_control

import (
	"context"
	protocol_pb "github.com/EmptyDea-Team/EmptyDea-core-api/pb/minecraft/protocol"
)

// ConstantPacket 规定常量数据包资源客户端的标准能力。
type ConstantPacket interface {
	// AllCreativeContent 返回全部创造物品。
	AllCreativeContent(ctx context.Context) (items []*protocol_pb.CreativeItem, err error)
	// CreativeItemByCNI 按创造物品网络 ID 查询创造物品。
	CreativeItemByCNI(ctx context.Context, creativeNetworkID uint32) (item *protocol_pb.CreativeItem, existed bool, err error)
	// CreativeItemByNI 按物品网络 ID 查询创造物品。
	CreativeItemByNI(ctx context.Context, networkID int32) (items []*protocol_pb.CreativeItem, err error)
	// CreativeItemByName 按名称查询创造物品。
	CreativeItemByName(ctx context.Context, name string) (items []*protocol_pb.CreativeItem, err error)
	// AllAvailableItems 返回全部可用物品。
	AllAvailableItems(ctx context.Context) (items []*protocol_pb.ItemEntry, err error)
	// ItemByNetworkID 按网络 ID 查询物品。
	ItemByNetworkID(ctx context.Context, networkID int32) (item *protocol_pb.ItemEntry, existed bool, err error)
	// ItemByName 按名称查询物品。
	ItemByName(ctx context.Context, name string) (item *protocol_pb.ItemEntry, existed bool, err error)
	// ItemNameByNetworkID 按网络 ID 查询物品名。
	ItemNameByNetworkID(ctx context.Context, networkID int32) (name string, existed bool, err error)
	// AllCommandItems 返回全部可通过命令获取的物品名。
	AllCommandItems(ctx context.Context) (items []string, err error)
	// ItemCanGetByCommand 判断物品是否可通过命令获取。
	ItemCanGetByCommand(ctx context.Context, name string) (canGet bool, err error)
	// TrimRecipeNetworkID 返回纹饰配方网络 ID。
	TrimRecipeNetworkID(ctx context.Context) (networkID uint32, err error)
}

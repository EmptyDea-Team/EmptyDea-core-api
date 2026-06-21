package define

import "context"

// ItemTransition 规定物品转移客户端的标准能力。
type ItemTransition interface {
	// Transition 在指定窗口间转移物品。
	Transition(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot, srcWindowName WindowName, dstWindowName WindowName) (bool, error)
	// TransitionBetweenInventory 在背包内转移物品。
	TransitionBetweenInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionBetweenContainer 在已打开容器内转移物品。
	TransitionBetweenContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionToContainer 从背包转移物品到已打开容器。
	TransitionToContainer(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
	// TransitionToInventory 从已打开容器转移物品到背包。
	TransitionToInventory(ctx context.Context, src []ItemInfoWithSlot, dst []ItemInfoWithSlot) (bool, error)
}

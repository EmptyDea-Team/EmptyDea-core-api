# EmptyDea-core-api

`EmptyDea-core-api` 是 EmptyDeaCore 的 gRPC/API 定义模块，负责存放框架 API、游戏控制 API、Minecraft protocol proto，以及对应生成的 Go `pb` 代码。

这个模块只描述跨进程通信协议，不包含服务端运行逻辑，也不包含原生 Minecraft 数据包转换逻辑。

## 目录结构

```text
define/                 手写的公共 Go 定义，例如 FrameConfig
proto/                  gRPC 和 protocol 的 proto 源文件
proto/game_control/     GameInterface 与 ResourcesControl API
proto/minecraft/        从 mousetunnel 源码生成的 Minecraft protocol proto
pb/                     protoc 生成的 Go 代码
generate.sh             重新生成 proto、pb 和 core 侧 converter 的脚本
```

## 生成

默认从 `/root/Yeah114/mousetunnel/minecraft/protocol` 读取 mousetunnel protocol 源码：

```bash
cd /root/Yeah114/EmptyDea-core-api
./generate.sh
```

也可以显式指定源码路径：

```bash
./generate.sh /path/to/mousetunnel/minecraft/protocol
```

或使用环境变量：

```bash
MOUSE_PROTOCOL_SRC=/path/to/mousetunnel/minecraft/protocol ./generate.sh
```

脚本会同时更新：

- 本模块的 `proto/minecraft/protocol`
- 本模块的 `pb`
- core 仓库中的 `/root/Yeah114/EmptyDea-core/frame/EmptyDeaCore/converter`

如果 core 仓库不在默认 sibling 路径，可以指定：

```bash
EMPTYDEA_CORE_ROOT=/path/to/EmptyDea-core ./generate.sh
```

## 依赖关系

`EmptyDea-core-api` 应保持轻量，只依赖 gRPC/protobuf 运行时。其他模块的职责如下：

- `EmptyDea-core`：实现 gRPC 服务端、原生 game control 适配、protocol converter。
- `EmptyDea-core-client`：封装客户端调用，面向使用者提供更接近源码风格的 API。

## 验证

```bash
cd /root/Yeah114/EmptyDea-core-api
go test ./...
```

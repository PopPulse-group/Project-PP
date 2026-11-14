# PopPulse 最新 go-zero 框架搭建说明

## 一、对 gzd2 的分析与优化

gzd2 的核心路线是正确的：RPC 持有数据库，API 只做门面，两个目录独立 module，8080/8888 分离。需要优化的地方有：

1. 生成代码不能手工伪造或长期保留两份 pb；每次 proto 变更先清理旧生成物，再生成并编译。
2. API 侧应通过 `zrpc.RpcClientConf` 注入 client，不能在全局 `init` 中硬编码 gRPC 地址，否则配置、测试和故障恢复都不可控。
3. RPC 应使用 `zrpc.MustNewServer`，API 应使用 `rest.MustNewServer`，配置文件作为唯一监听入口。
4. 业务代码按 `handler → logic → svc → rpc/model` 分层；生成目录只由 goctl 管理，logic 和 model 手写。
5. proto、api、生成、校验、启动和联调命令应写进项目 README，避免依赖个人记忆。
6. 当前 PopPulse 先用模块化单体的 product 示例验证链路，接入 MySQL 时只改 RPC 的 `internal/model` 和 logic，API 不增加数据库依赖。

## 二、最终职责

```text
HTTP :8888 product-api
  └── zrpc client
       └── gRPC :8080 product
            └── internal/model (后续接 GORM/MySQL)
```

## 三、工具准备

```powershell
go version
```

```powershell
goctl --version
```

```powershell
protoc --version
```

```powershell
Get-Command protoc-gen-go,protoc-gen-go-grpc | Select-Object Name,Source
```

缺少插件时：

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```powershell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 四、RPC 创建与更新

RPC 目录必须使用 `product` 这类不带连字符的名称，因为 goctl 会把目录名写入 proto service 标识符。

```powershell
goctl rpc new product
```

将 proto 的 `go_package` 固定为 `"./product"`，然后生成：

```powershell
goctl rpc protoc product.proto --go_out=./product --go-grpc_out=./product --zrpc_out=.
```

生成后只保留 `product/product/*.pb.go`，删除根目录旧模板 pb 和无关 Ping logic。确认 `product.go`、server、logic import 指向同一份 pb。

RPC 配置必须包含 `Name`、`ListenOn`、`Mode`；未使用服务发现时不要保留 Etcd 配置，避免启动时等待 Etcd。

## 五、API 创建与更新

```powershell
goctl api new product-api
```

编辑 `.api` 后依次执行：

```powershell
goctl api validate -api product.api
```

```powershell
goctl api format -dir .
```

```powershell
goctl api go -api product.api -dir .
```

API 配置只包含 HTTP 监听配置和 RPC client 配置，不放数据库 DSN。路径参数使用 `path:"id"`，请求体使用 `json:"..."`。

## 六、依赖与构建

两个 module 分别执行：

```powershell
go mod tidy
```

```powershell
go build ./...
```

如果网络不可用，先设置企业代理或可访问的 GOPROXY；不要提交临时生成的 `vendor` 和模块缓存。

## 七、启动

先启动 RPC：

```powershell
cd product; go run product.go -f etc/product.yaml
```

再启动 API：

```powershell
cd product-api; go run product.go -f etc/product-api.yaml
```

## 八、API/RPC 联调

```powershell
Test-NetConnection 127.0.0.1 -Port 8080
```

```powershell
Invoke-RestMethod http://localhost:8888/product/1
```

停掉 RPC 后重新请求：

```powershell
$rpcPid=(Get-NetTCPConnection -LocalPort 8080 -State Listen).OwningProcess; Stop-Process -Id $rpcPid -Force
```

若 API 返回 Unavailable/503，说明请求确实依赖 RPC。恢复 RPC 后再请求即可。

## 九、后续接入 PopPulse

第一阶段建议将 product 示例替换为 `user`、`product`、`trade` 三个领域；仍保持 API 不直连数据库。订单、库存、支付状态机放 RPC logic，数据库模型放对应 RPC 的 `internal/model`，跨服务事件使用 Outbox/Kafka，避免把 Redis 当作资金或订单事实来源。

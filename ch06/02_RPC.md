
RPC 架构主要包括 3 部分，分别是服务注册中心、服务提供者和服务消费者，其功能如下:
* 服务注册中心: 负责将本地服务发布成远程服务，管理远程服务，提供给服务消费者使用。
* 服务提供者: 提供服务接口的定义与服务类的实现。
* 服务消费者: 通过远程代理对象调用远程服务。

服务提供者在启动后，会主动向服务注册中心注册机器的 IP 地址、端口号，以及提供的服务列表；服务消费者在启动后，会向服务注册中心获取服务提供方的服务列表。服务注册中心可实现负载均衡和故障切换。

Go 语言提供了 net/rpc 的 RPC 包，它使用 encoding/gob 包中的 Encoder 对象和 Decoder 对象中可以进行 GOB 格式的编码和解码，并且支持 TCP 或 HTTP 数据传输方式。

[简单的 grpc 服务器示例](t/02_gob_rpc_server.go)

[简单的 grpc 客户端示例](t/02_gob_rpc_client.go)

由于其他语言不支持 GOB 格式编/解码方式，所以使用 net/rpc 包实现的 RPC 方法没办法进行跨语言调用。

JSON 编码 RPC 是指，数据编码采用了 JSON 格式的 RPC.

[json数据编码的 grpc 服务器示例](t/02_json_grpc_server.go)

[json数据编码的 grpc 客户端示例](t/02_json_grpc_client.go)

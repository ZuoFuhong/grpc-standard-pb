# gRPC PB 统一仓库

**1. PB 命名约定**

- 文件夹命名：`{服务标识}`
- service命名：`{服务标识}`
- pb 文件名: `{服务标识}.proto`
- package格式: `{服务标识}`
- go_package约定: `github.com/ZuoFuhong/grpc-standard-pb/{服务标识}`

rpc cmd、参数采用驼峰模式，message 内部的字段采用蛇形命名法，请求和响应参数分别用 Req、Rsp 结尾

**2. 编译命令**

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative go_wallet_manage_svr.proto
```

**3. 模块初始化**

```shell
go mod init github.com/grpc-standard-pb/go_wallet_manage_svr
go get -u google.golang.org/grpc
```

## License

This project is licensed under the [Apache 2.0 license](https://github.com/ZuoFuhong/grpc-standard-pb/blob/master/LICENSE).
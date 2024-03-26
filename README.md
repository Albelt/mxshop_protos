## protoc命令解释

```shell
protoc -I /Users/albelt/Gitlab/common/protos -I . \ #依赖的第三方proto路径,待编译的proto路径
    --go_out=xxx \ #protoc-gen-go插件,生成go结构体
    --go-grpc_out=xxx \ #protoc-gen-go-grpc插件,生成grpc方法
    --validate_out="lang=go:albelt/good_srv" \ #protoc-gen-validate插件,生成校验go结构体的方法
    --doc_out=./doc --doc_opt=html,xxx.html \ #protoc-gen-doc插件,生成html文档
    albelt/good_srv/msg.proto albelt/good_srv/svc.proto #proto文件路径
```

protoc插件:
- protobuf校验: https://github.com/bufbuild/protoc-gen-validate
- protobuf生成文档: https://github.com/pseudomuto/protoc-gen-doc


## 自定义proroc插件:

- 原理参考文章: https://taoshu.in/go/create-protoc-plugin.html

### protoc-gen-md插件
```shell
cd protoc-gen-md && go build
cp protoc-gen-md $GOPATH/bin

protoc -I /Users/albelt/Gitlab/common/protos -I . --md_out=./ hello.proto 
```

### protoc-gen-gin插件
```shell
cd protoc-gen-gin && go build
cp protoc-gen-gin $GOPATH/bin

protoc -I /Users/albelt/Gitlab/common/protos -I . --gin_out=./ hello.proto 
```
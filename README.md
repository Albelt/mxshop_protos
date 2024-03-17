protoc命令解释:
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
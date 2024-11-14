# go-getting-started
Learn Golang

```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

if error

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest google.golang.org/protobuf/cmd/protoc-gen-go@latest google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


install buf (https://github.com/bufbuild/buf):

```bash
brew install bufbuild/buf/buf
```

gen with buf

```
buf generate
```


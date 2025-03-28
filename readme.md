## 專案目錄說明 go 1.23
```
project/
├── proto/                  # 放置所有 .proto 文件
│   └── hello.proto         # gRPC 協議定義
├── helloworld/
│   └── proto/                  # 放置 .proto 生成的文件
│       ├── hello.pb.go         # protoc 生成的消息代碼
│       └── hello_grpc.pb.go    # protoc 生成的 gRPC 服務代碼
├── server/                 # 服務端代碼
│   └── main.go
├── client/                 # 客戶端代碼
│   └── main.go
├── go.mod                  # Go 模組文件
└── go.sum
```

## 建立執行檔
```
go build -o server.exe ./server
go build -o client.exe ./client

CGO_ENABLE=0 GOOS=linux go build -o server ./server
CGO_ENABLE=0 GOOS=linux go build -o client.exe ./client
```

## 產生proto檔案
```
go version >=1.22

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

dlv version
go install github.com/go-delve/delve/cmd/dlv@latest

protoc --go_out=. --go-grpc_out=. proto/hello.proto
protoc --go_out=. --go-grpc_out=. proto/mfa.proto
```
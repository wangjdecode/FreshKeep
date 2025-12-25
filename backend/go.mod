module github.com/freshkeep/backend

go 1.21

require (
	github.com/go-kratos/kratos/v2 v2.7.2
	github.com/google/wire v0.5.0
	google.golang.org/protobuf v1.31.0
	github.com/golang/protobuf v1.5.3
	google.golang.org/grpc v1.59.0
	gorm.io/gorm v1.25.5
	gorm.io/driver/postgres v1.5.4
	gorm.io/driver/sqlite v1.5.4
	github.com/gorilla/mux v1.8.1
)

// 告诉 Go 这是一个私有/本地模块，不要尝试从远程获取
// replace 指令指向当前目录（相对于 go.mod 文件）
replace github.com/freshkeep/backend => ./


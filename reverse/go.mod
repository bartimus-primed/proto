module github.com/bartimus-primed/proto/reverse

go 1.17

// protoc --go_out=. --go_opt=paths=source_relative .\ghost.proto
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\ghost.proto
require google.golang.org/grpc v1.40.0

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210909211513-a8c4777a87af // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

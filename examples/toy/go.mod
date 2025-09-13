module github.com/qts0312/ChaosRPC/examples/toy

go 1.24.1

replace google.golang.org/grpc => github.com/qts0312/grpc v0.0.0-20250912120622-4656bf76be87

require (
	github.com/qts0312/ChaosRPC v0.0.0-20250913082618-eabe98b6defe
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.9
)

require (
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
)

module grpc-go-course

replace (
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190422233926-fe54fb35175b
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)

require (
	github.com/golang/protobuf v1.3.1
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
)

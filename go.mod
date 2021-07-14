module github.com/gs2io/gs2-golang-sdk

go 1.15

require (
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	core v0.0.0-00010101000000-000000000000
)

replace core => ./core

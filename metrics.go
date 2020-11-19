package grpc_proxy

type ICounter interface {
	Inc()
	Dec()
}

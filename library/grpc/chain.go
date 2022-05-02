package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

/*
 * 构造多层调用堆栈，把grpc原生支持单个拦截器调整为链式多级拦截器
 */

// unary服务器拦截器
func ChainUnary(interceptors []grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	nums := len(interceptors)

	switch {
	case nums > 1:
		last := nums - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var chain_handler grpc.UnaryHandler = nil
			var cur int = 0
			chain_handler = func(current_ctx context.Context, current_req interface{}) (interface{}, error) {
				if cur == last {
					return handler(current_ctx, current_req)
				}
				cur++
				resp, err := interceptors[cur](current_ctx, current_req, info, chain_handler)
				cur--
				return resp, err
			}
			return interceptors[0](ctx, req, info, chain_handler)
		}
	case 1 == nums:
		return interceptors[0]
	case 0 == nums:
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		}
	}

	return nil
}

// stream服务拦截器
func ChainStream(interceptors []grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	nums := len(interceptors)

	switch {
	case nums > 1:
		last := nums - 1
		return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			var chain_handler grpc.StreamHandler = nil
			var cur int = 0
			chain_handler = func(current_srv interface{}, current_stream grpc.ServerStream) error {
				if cur == last {
					return handler(current_srv, current_stream)
				}
				cur++
				err := interceptors[cur](current_srv, current_stream, info, chain_handler)
				cur--
				return err
			}
			return interceptors[0](srv, stream, info, chain_handler)
		}
	case 1 == nums:
		return interceptors[0]
	case 0 == nums:
		return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, stream)
		}
	}

	return nil
}

// unary客户端拦截器
func ChainUnaryClient(interceptors []grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	switch {
	case n > 1:
		last := n - 1
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var chain_handler grpc.UnaryInvoker
			var cur int = 0

			chain_handler = func(current_ctx context.Context, current_method string, current_req,
				current_reply interface{}, current_conn *grpc.ClientConn, current_opts ...grpc.CallOption) error {
				if cur == last {
					return invoker(current_ctx, current_method, current_req, current_reply,
						current_conn, current_opts...)
				}
				cur++
				err := interceptors[cur](current_ctx, current_method, current_req, current_reply,
					current_conn, chain_handler, current_opts...)
				cur--
				return err
			}

			return interceptors[0](ctx, method, req, reply, cc, chain_handler, opts...)
		}
	case n == 1:
		return interceptors[0]
	case n == 0:
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}
	}

	return nil
}

// stream客户端拦截器
func ChainStreamClient(interceptors []grpc.StreamClientInterceptor) grpc.StreamClientInterceptor {
	n := len(interceptors)

	switch {
	case n > 1:
		last := n - 1
		return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
			streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			var chain_handler grpc.Streamer
			var cur int

			chain_handler = func(current_ctx context.Context, current_desc *grpc.StreamDesc,
				current_conn *grpc.ClientConn, current_method string, current_opt ...grpc.CallOption) (grpc.ClientStream, error) {
				if cur == last {
					return streamer(current_ctx, current_desc, current_conn, current_method, current_opt...)
				}
				cur++
				stream, err := interceptors[cur](current_ctx, current_desc, current_conn, current_method, chain_handler, current_opt...)
				cur--
				return stream, err
			}

			return interceptors[0](ctx, desc, cc, method, chain_handler, opts...)
		}

	case n == 1:
		return interceptors[0]

	case n == 0:
		return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			return streamer(ctx, desc, cc, method, opts...)
		}
	}

	return nil
}

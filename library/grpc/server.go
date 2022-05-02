package grpc

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

var (
	NotInitError error = errors.New("server not init")
)

// 服务实例
type Engine struct {
	*grpc.Server                                // 原始grpc服务实例
	UnaryChains  []grpc.UnaryServerInterceptor  // unary拦截器
	StreamChains []grpc.StreamServerInterceptor // stream拦截器
	opts         []grpc.ServerOption            // 附加选项
}

// 成员变量懒初始化
func (e *Engine) lazyInit() {
	if e.UnaryChains == nil {
		e.UnaryChains = make([]grpc.UnaryServerInterceptor, 0)
	}
	if e.StreamChains == nil {
		e.StreamChains = make([]grpc.StreamServerInterceptor, 0)
	}
}

// 中间件设置。
// 注意：
// 1. 只有在InitServer之前调用才会生效
func (e *Engine) Use(unaryMiddleware grpc.UnaryServerInterceptor, streamMiddleware grpc.StreamServerInterceptor) {
	e.lazyInit()
	if unaryMiddleware != nil {
		e.UnaryChains = append(e.UnaryChains, unaryMiddleware)
	}
	if streamMiddleware != nil {
		e.StreamChains = append(e.StreamChains, streamMiddleware)
	}
}

// 初始化服务端
func (e *Engine) InitServer() error {
	e.opts = append(e.opts, grpc.UnaryInterceptor(ChainUnary(e.UnaryChains)))
	e.opts = append(e.opts, grpc.StreamInterceptor(ChainStream(e.StreamChains)))
	e.Server = grpc.NewServer(e.opts...)

	return nil
}

// 初始化TLS服务端
func (e *Engine) InitTLSServer(certFile string, certKey string) error {
	creds, err := credentials.NewServerTLSFromFile(certFile, certKey)
	if err != nil {
		return errors.Wrap(err, "创建credentials失败")
	}
	e.opts = append(e.opts, grpc.Creds(creds))
	e.opts = append(e.opts, grpc.UnaryInterceptor(ChainUnary(e.UnaryChains)))
	e.opts = append(e.opts, grpc.StreamInterceptor(ChainStream(e.StreamChains)))
	e.Server = grpc.NewServer(e.opts...)
	return nil
}

// 执行运行服务端，addr参数形如：0.0.0.0:8000
func (e *Engine) Run(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("监听地址[%s]失败", addr))
	}
	if err := e.Serve(listen); err != nil {
		return errors.Wrap(err, "启动服务异常")
	}
	return nil
}

// 执行运行TLS服务端，addr参数形如：0.0.0.0:8000
func (e *Engine) RunTLS(addr string, certFile string, certKey string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("监听地址[%s]失败", addr))
	}
	if err := e.Serve(listen); err != nil {
		return errors.Wrap(err, "启动服务异常")
	}
	return nil
}

// 默认的KeepAlive配置
func (e *Engine) UseDefaultKeepaliveOpt() {
	var keepAlivePolicy = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}

	var keepAliveParam = keepalive.ServerParameters{
		MaxConnectionIdle: 15 * time.Minute,
		//MaxConnectionAge:      10 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               1 * time.Second,
	}

	ep := grpc.KeepaliveEnforcementPolicy(keepAlivePolicy)
	kp := grpc.KeepaliveParams(keepAliveParam)

	e.opts = append(e.opts, ep)
	e.opts = append(e.opts, kp)
}

// 默认的服务端
func Default() *Engine {
	e := &Engine{}
	//e.Use(middleware.UnaryServerContextWrapper, nil)
	//e.Use(middleware.UnaryServerErrorHook, middleware.StreamServerErrorHook)
	//e.Use(middleware.UnaryServerRecoverInterceptor(), middleware.StreamServerRecoverInterceptor())
	return e
}

func (e *Engine) GetRawServer() *grpc.Server {
	return e.Server
}

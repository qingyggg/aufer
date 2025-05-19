package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/qingyggg/aufer/biz/rpc"
	mydal "github.com/qingyggg/aufer/cmd/relation/dal"
	relation "github.com/qingyggg/aufer/kitex_gen/cmd/relation/relationhandler"
	"github.com/qingyggg/aufer/pkg/bound"
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/middleware"
	"github.com/qingyggg/aufer/pkg/tracer"
	"github.com/qingyggg/aufer/pkg/utils"
	"net"
)

func init() {
	utils.EnvInit()
	constants.UrlInit()
	rpc.InitRpc() //初始化rpc连接
	mydal.Init()
	tracer.InitJaeger(constants.RelationService)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":18011")
	if err != nil {
		panic(err)
	}
	svr := relation.NewServer(new(RelationHandlerImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationService}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

}

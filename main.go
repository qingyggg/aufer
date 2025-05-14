package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/qingyggg/aufer/biz/dal/minio"
	"github.com/qingyggg/aufer/biz/dal/redis"
	"github.com/qingyggg/aufer/biz/mw/jwt"
	"github.com/qingyggg/aufer/biz/mw/logger"
	"github.com/qingyggg/aufer/biz/rpc"
	_ "github.com/qingyggg/aufer/docs" // 导入swagger文档
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/tracer"
	"github.com/qingyggg/aufer/pkg/utils"
	"os"
)

// @title Aufer API
// @version 1.0
// @description Aufer平台API服务
// @BasePath /
// @schemes https http
func main() {
	validateConfig := GetCustomValidateConfig()

	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts("0.0.0.0:18014"),
		server.WithValidateConfig(validateConfig),
	)
	// default is "debug/pprof"
	pprof.Register(h, "dev/pprof")

	register(h)
	h.Spin()
}

// Set up /src/*name route forwarding to access minio from external network
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxyUrl := "http://" + constants.MinioEndPoint
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(proxyUrl)
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, "minio图片访问==>"+string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func init() {
	utils.EnvInit()
	constants.UrlInit()
	if os.Getenv("FOO_ENV") == "production" {
		logger.InitLogger()
	}
	jwt.Init()
	minio.InitForHertz()
	redis.InitRedisForHertz()
	rpc.InitRpc() //初始化rpc连接
	tracer.InitJaeger(constants.ApiService)
}
func GetCustomValidateConfig() *binding.ValidateConfig {
	//自定义参数校验
	validateConfig := &binding.ValidateConfig{}
	validateConfig.MustRegValidateFunc("password", func(args ...interface{}) error {
		err := utils.ValidatePassword(fmt.Sprint(args...))
		if err != nil {
			return err
		}
		return nil
	})
	return validateConfig
}

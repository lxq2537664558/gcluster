package main

import (
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/codec"
	"github.com/gosrv/gcluster/gbase/gdb/gmongo"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/gcluster/gbase/gmxdriver"
	"github.com/gosrv/gcluster/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/controller"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

// 客户端网络出了
func initBaseNet(builder gioc.IBeanContainerBuilder) {
	// pb消息编解码器，4字节长度 + 2字节proto id + proto
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)
	// 创建网络模块，这里使用了数据自动同步路由器AutoSyncDataRoute
	net := tcpnet.NewTcpNetServer("pcluster.basenet", "",
		encoder, decoder, nil, entity.NewAutoSyncDataRoute())
	builder.AddBean(net)
}

func initClusterMsgCenter(builder gioc.IBeanContainerBuilder) {
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewIdProtobufEncoder(idtype)
	decoder := codec.NewIdProtobufDecoder(idtype)
	builder.AddBean(common.NewClusterMsgCenter(encoder, decoder))
}

func initServices(builder gioc.IBeanContainerBuilder) {
	builder.AddBean(
		// redis 自动配置
		gredis.NewAutoConfigReids("pcluster.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("pcluster.mongo", ""),
		ghttp.NewHttpServer("pcluster.http", nil),
		gmxdriver.NewGMXDriver("/gmx"),
		common.NewGmxAppStats(),
		common.NewClusterNodeMgr(),
		controller.NewControllerLogin(),
		controller.NewControllerLogic(),
		service.NewPlayerMgr(),
		service.NewServiceDataAutoSync(),
		service.NewServiceLogic(),
		service.NewServiceLogin(),
		service.NewServicePlayerMsgQueue(),
	)
}

func initLog(configLoader gioc.IConfigLoader, builder gioc.IBeanContainerBuilder) {
	logroot := &glog.ConfigLogRoot{}
	configLoader.Config().Get("pcluster.log").Scan(logroot)

	logBuilder := glog.NewLogFactoryBuilder()
	logFactory, err := logBuilder.Build(logroot)
	util.VerifyNoError(err)

	// 重定向系统日志
	err = gl.Redirect(logFactory.GetLogger("engine"))
	util.VerifyNoError(err)
	builder.AddBean(logFactory)
}

func main() {
	application := app.NewApplication()
	configLoader := application.InitCli()
	builder := application.InitBuilder()
	// 初始化日志
	initLog(configLoader, builder)

	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	initBaseNet(builder)
	initClusterMsgCenter(builder)
	beanContainer := application.Build(builder)

	application.Start(beanContainer)
}

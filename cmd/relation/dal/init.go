package mydal

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/qingyggg/aufer/biz/dal"
	"github.com/qingyggg/aufer/biz/dal/minio"
	"github.com/qingyggg/aufer/biz/dal/mongo"
	"github.com/qingyggg/aufer/biz/dal/redis"
	"github.com/qingyggg/aufer/biz/model/query"
)

var Qdb *query.Query
var MyMongo *mongo.MyMongo
var MyRedis *redis.MyRedis
var MyMinio *minio.MyMinio

func Init() {
	err := dal.Init()
	if err != nil {
		klog.Fatal("初始化dal发生错误==>", err)
	}
	myDal := dal.MyDal
	Qdb = myDal.Qdb
	MyMongo = myDal.Mongo
	MyRedis = myDal.Mrs
	MyMinio = myDal.Mio
}

package dal

import (
	"github.com/qingyggg/aufer/biz/dal/db"
	"github.com/qingyggg/aufer/biz/dal/minio"
	mymongo "github.com/qingyggg/aufer/biz/dal/mongo"
	"github.com/qingyggg/aufer/biz/dal/redis"
	"github.com/qingyggg/aufer/biz/model/query"
	"github.com/qingyggg/aufer/pkg/errno"
)

type Dal struct {
	Qdb   *query.Query     //db query实体
	Mrs   *redis.MyRedis   //redis实体
	Mongo *mymongo.MyMongo //mongo实体
	Mio   *minio.MyMinio   //minio实体
}

var MyDal *Dal

// Init init dal
func Init() error {
	MyDal = new(Dal)

	err, qdb := db.Init() // mysql init
	if err != nil {
		return errno.ServiceErr.WithMessage("初始化mariadb失败，失败原因：" + err.Error())
	}
	MyDal.Qdb = qdb

	mrs := redis.InitRedis() //redis init
	MyDal.Mrs = mrs

	err, mg := mymongo.Init()
	if err != nil {
		return errno.ServiceErr.WithMessage("初始化mongodb失败，失败原因：" + err.Error())
	}
	MyDal.Mongo = mg

	err, mio := minio.Init()
	if err != nil {
		return errno.ServiceErr.WithMessage("初始化minio失败，失败原因：" + err.Error())
	}
	MyDal.Mio = mio

	return nil
}

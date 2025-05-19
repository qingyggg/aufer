package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/qingyggg/aufer/pkg/constants"
	"time"
)

var (
	ExpireTime = time.Hour * 1
)

type MyRedis struct {
	RdbCollect    *Collect
	RdbFavorite   *Favorite
	RdbComment    *Comment
	RdbView       *View
	RdbUser       *User
	RdbUserOnline *Online
}

func InitRedis() *MyRedis {
	mr := new(MyRedis)

	mr.RdbCollect = &Collect{rsc: redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})}
	mr.RdbFavorite = &Favorite{rsc: redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})}
	mr.RdbComment = &Comment{rsc: redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       2,
	})}
	mr.RdbView = &View{rsc: redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       3,
	})}
	//uid:uHashId,用以减小请求User数据库的次数
	mr.RdbUser = &User{redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       4,
	})}
	//uhashId:1
	mr.RdbUserOnline = &Online{rsc: redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       5,
	})}
	return mr
}

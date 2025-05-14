package redis

import (
	"github.com/go-redis/redis/v7"
	"strconv"
)

const (
	uidSuffix = ":uid"
)

type User struct {
	rsc *redis.Client
}

func strget(c *redis.Client, k string) string {
	// 使用 GET 命令获取键对应的值
	val := c.Get(k).String()
	c.Expire(k, ExpireTime) //重置过期时间
	return val              // 返回获取到的值
}
func strset(c *redis.Client, k string, v string) error {
	tx := c.TxPipeline()
	tx.Set(k, v, ExpireTime)
	_, err := tx.Exec()
	return err
}
func (u *User) UHashIdExist(uid int64) (error, bool) {
	return check(u.rsc, strconv.FormatInt(uid, 16))
}
func (u *User) UHashIdGet(uid int64) string {
	return strget(u.rsc, strconv.FormatInt(uid, 16)+uidSuffix)
}
func (u *User) UHashIdSet(uid int64, uHashId string) error {
	return strset(u.rsc, strconv.FormatInt(uid, 16)+uidSuffix, uHashId)
}

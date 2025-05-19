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

func (u *User) UHashIdExist(uid int64) (error, bool) {
	return check(u.rsc, strconv.FormatInt(uid, 16))
}
func (u *User) UHashIdGet(uid int64) string {
	return strGet(u.rsc, strconv.FormatInt(uid, 16)+uidSuffix, false)
}
func (u *User) UHashIdSet(uid int64, uHashId string) error {
	return strSet(u.rsc, strconv.FormatInt(uid, 16)+uidSuffix, uHashId, false)
}

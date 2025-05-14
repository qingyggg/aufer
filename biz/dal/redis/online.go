package redis

import "github.com/go-redis/redis/v7"

type Online struct {
	rsc *redis.Client
}

// OnlExist 检查用户是否在线
func (o Online) OnlExist(uHashId string) (error, bool) {
	return check(o.rsc, uHashId)
}

// SetOnline 用户在线，设置在线状态
func (o Online) SetOnline(uHashId string, socketID string) error {
	return strSet(o.rsc, uHashId, socketID, true)
}

// GetSocketID 根据用户id获取socketId,用以广播
func (o Online) GetSocketID(uHashId string) string {
	return strGet(o.rsc, uHashId, true)
}

// RemOnline 用户登出，或者下线
func (o Online) RemOnline(uHashId string) error {
	return delKey(o.rsc, uHashId)
}

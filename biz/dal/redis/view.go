package redis

import "github.com/go-redis/redis/v7"

type View struct {
	rsc *redis.Client
}

const ViewCountSuffix = ":view_count"

func (c View) GetViewClient() *redis.Client {
	return c.rsc
}

func (c View) IncrView(aHashId string) error {
	return incr(c.rsc, aHashId+ViewCountSuffix)
}

func (c View) DecrView(aHashId string) error {
	return decr(c.rsc, aHashId+ViewCountSuffix)
}

func (c View) CountView(aHashId string) (error, int64) {
	return nget(c.rsc, aHashId+ViewCountSuffix)
}

func (c View) ViewCtAssign(aHashId string, count int64) error {
	return nset(c.rsc, aHashId+ViewCountSuffix, count)
}

func (c View) CheckViewCt(aHashId string) (error, bool) {
	return check(c.rsc, aHashId+ViewCountSuffix)
}

func (c View) DelViewCt(aHashId string) error {
	return delKey(c.rsc, aHashId+ViewCountSuffix)
}

func (c View) GetViewMap(aHashIds []string) (error, map[string]int64) {
	return ngets(c.rsc, aHashIds, ViewCountSuffix)
}

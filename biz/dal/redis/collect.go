package redis

import "github.com/go-redis/redis/v7"

type Collect struct {
	rsc *redis.Client
}

const CollectCountSuffix = ":collect_count"

func (c Collect) GetCollectClient() *redis.Client {
	return c.rsc
}

func (c Collect) IncrCollect(aHashId string) error {
	return incr(c.rsc, aHashId+CollectCountSuffix)
}

func (c Collect) DecrCollect(aHashId string) error {
	return decr(c.rsc, aHashId+CollectCountSuffix)
}

func (c Collect) CountCollect(aHashId string) (error, int64) {
	return nget(c.rsc, aHashId+CollectCountSuffix)
}

func (c Collect) CollectCtAssign(aHashId string, count int64) error {
	return nset(c.rsc, aHashId+CollectCountSuffix, count)
}

func (c Collect) CheckCollectCt(aHashId string) (error, bool) {
	return check(c.rsc, aHashId+CollectCountSuffix)
}

func (c Collect) DelCollectCt(aHashId string) error {
	return delKey(c.rsc, aHashId+CollectCountSuffix)
}

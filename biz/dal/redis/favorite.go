package redis

import "github.com/go-redis/redis/v7"

const (
	likeSuffix      = ":liked"
	hateSuffix      = ":hated"
	LikeCountSuffix = ":liked_count"
	HateCountSuffix = ":hated_count"
)

type Favorite struct {
	rsc *redis.Client
}

func (f Favorite) GetFavoriteClient() *redis.Client {
	return f.rsc
}

func (f Favorite) Like(uHashId, aHashId string) error {
	return sadd(f.rsc, aHashId+likeSuffix, uHashId)
}

func (f Favorite) CancerLike(uHashId, aHashId string) error {
	return srem(f.rsc, aHashId+likeSuffix, uHashId)
}

func (f Favorite) ExistLike(uHashId, aHashId string) (error, bool) {
	return sexist(f.rsc, aHashId+likeSuffix, uHashId)
}

func (f Favorite) CheckLike(aHashId string) (error, bool) {
	return check(f.rsc, aHashId+likeSuffix)
}

func (f Favorite) Hate(uHashId, aHashId string) error {
	return sadd(f.rsc, aHashId+hateSuffix, uHashId)
}

func (f Favorite) CancerHate(uHashId, aHashId string) error {
	return srem(f.rsc, aHashId+hateSuffix, uHashId)
}

func (f Favorite) ExistHate(uHashId, aHashId string) (error, bool) {
	return sexist(f.rsc, aHashId+hateSuffix, uHashId)
}

func (f Favorite) CheckHate(aHashId string) (error, bool) {
	return check(f.rsc, aHashId+hateSuffix)
}

// IncrLike 文章点赞处理逻辑
func (f Favorite) IncrLike(aHashId string) error {
	return incr(f.rsc, aHashId+LikeCountSuffix)
}

func (f Favorite) IncrHate(aHashId string) error {
	return incr(f.rsc, aHashId+HateCountSuffix)
}

func (f Favorite) DecrLike(aHashId string) error {
	return decr(f.rsc, aHashId+LikeCountSuffix)
}

func (f Favorite) DecrHate(aHashId string) error {
	return decr(f.rsc, aHashId+HateCountSuffix)
}

func (f Favorite) CountLike(aHashId string) (error, int64) {
	return nget(f.rsc, aHashId+LikeCountSuffix)
}

func (f Favorite) CountHate(aHashId string) (error, int64) {
	return nget(f.rsc, aHashId+HateCountSuffix)
}

//从数据库count点赞和踩的数量，并且赋值给redis

func (f Favorite) LikeCtAssign(aHashId string, count int64) error {
	return nset(f.rsc, aHashId+LikeCountSuffix, count)
}

func (f Favorite) HateCtAssign(aHashId string, count int64) error {
	return nset(f.rsc, aHashId+HateCountSuffix, count)
}

func (f Favorite) CheckLikeCt(aHashId string) (error, bool) {
	return check(f.rsc, aHashId+LikeCountSuffix)
}

func (f Favorite) CheckHateCt(aHashId string) (error, bool) {
	return check(f.rsc, aHashId+HateCountSuffix)
}

// TruncateLikeStatus 用户删除文章后，文章点赞依赖删除逻辑：
func (f Favorite) TruncateLikeStatus(aHashId string) error {
	err := delKey(f.rsc, aHashId+likeSuffix)
	if err != nil {
		return err
	}
	err = delKey(f.rsc, aHashId+LikeCountSuffix)
	return err
}

func (f Favorite) TruncateHateStatus(aHashId string) error {
	err := delKey(f.rsc, aHashId+hateSuffix)
	if err != nil {
		return err
	}
	err = delKey(f.rsc, aHashId+HateCountSuffix)
	return err
}

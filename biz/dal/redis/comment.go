package redis

import "github.com/go-redis/redis/v7"

type Comment struct {
	rsc *redis.Client
}

const CommentCountSuffix = ":comment_count"

func (c Comment) GetCommentClient() *redis.Client {
	return c.rsc
}

func (c Comment) IncrComment(aHashId string) error {
	return incr(c.rsc, aHashId+CommentCountSuffix)
}

func (c Comment) DecrComment(aHashId string) error {
	return decr(c.rsc, aHashId+CommentCountSuffix)
}

func (c Comment) CountComment(aHashId string) (error, int64) {
	return nget(c.rsc, aHashId+CommentCountSuffix)
}

func (c Comment) CommentCtAssign(aHashId string, count int64) error {
	return nset(c.rsc, aHashId+CommentCountSuffix, count)
}

func (c Comment) CheckCommentCt(aHashId string) (error, bool) {
	return check(c.rsc, aHashId+CommentCountSuffix)
}

func (c Comment) DelCommentCt(aHashId string) error {
	return delKey(c.rsc, aHashId+CommentCountSuffix)
}

package db

import (
	"context"
	"fmt"
	"github.com/qingyggg/aufer/biz/dal/redis"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetCmtCtByAids 通过文章id获取该文章的评论数
func GetCmtCtByAids(ctx context.Context, aHashIds []string) (error, map[string]int64) {
	// 将所有的aid从Redis里获取，获取不到的aid收集起来，一起从数据库中获取，获取后将这些键值对存储进Redis
	var aidInCache []string // 存在于缓存里的aHashId:count
	var aidInNoCache []string
	var resMap = make(map[string]int64) // 初始化结果map
	var rdComment = mydal.MyRedis.RdbComment
	var cols = mydal.MyMongo.Cols
	// 遍历输入的aHashIds，检查每个ID是否存在于缓存中
	for _, v := range aHashIds {
		err, exist := rdComment.CheckCommentCt(v)
		if err != nil {
			return err, nil
		}
		if exist {
			aidInCache = append(aidInCache, v)
		} else {
			aidInNoCache = append(aidInNoCache, v)
		}
	}

	// 从缓存中获取点赞数的goroutine
	for _, v := range aidInCache {
		err, count := rdComment.CountComment(v)
		if err != nil {
			return err, nil
		}
		resMap[v] = count
	}

	// 查询MongoDB数据库并将结果同步到Redis的goroutine
	if len(aidInNoCache) > 0 {
		pipeline := []bson.M{
			{
				"$match": bson.M{
					"article_id": bson.M{"$in": aidInNoCache}, // 匹配 article_id 在给定列表中的记录
				},
			},
			{
				"$group": bson.M{
					"_id":   "$article_id",     // 按 article_id 分组
					"count": bson.M{"$sum": 1}, // 统计每组的记录数
				},
			},
		}

		// 执行聚合查询
		cursor, err := cols.Comment.Aggregate(ctx, pipeline)
		if err != nil {
			return err, nil
		}
		defer cursor.Close(ctx)

		// 解析聚合结果,缓存进redis
		fpipe := rdComment.GetCommentClient().TxPipeline()
		for cursor.Next(ctx) {
			// 查询Mongo
			item := make(map[string]interface{})
			if err := cursor.Decode(&item); err != nil {
				return err, nil
			}
			fmt.Println(item)
			resMap[item["_id"].(string)] = int64(item["count"].(int32))
			fpipe.Set(item["_id"].(string)+redis.CommentCountSuffix, int64(item["count"].(int32)), redis.ExpireTime)
		}
		// 执行事务
		_, err = fpipe.Exec()
		if err != nil {
			return err, nil
		}
	}

	//这里如果缓存和MongoDB都没有查找到comment count，
	//说明这篇文章并没有被评论过，所以默认赋值为0
	for _, v := range aHashIds {
		if _, exist := resMap[v]; exist == false {
			resMap[v] = 0
		}
	}
	return nil, resMap
}
func GetCmtCountByAid(ctx context.Context, aHashId string) (error, int64) {
	var rdComment = mydal.MyRedis.RdbComment
	var cols = mydal.MyMongo.Cols
	err, exist := rdComment.CheckCommentCt(aHashId)
	if err != nil {
		return err, 0
	}
	if exist {
		return rdComment.CountComment(aHashId)
	} else {
		count, err := cols.Comment.CountDocuments(ctx, &bson.M{"article_id": aHashId})
		if err != nil {
			return err, 0
		}
		return nil, count
	}
}

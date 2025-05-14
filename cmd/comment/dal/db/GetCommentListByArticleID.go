package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetCommentListByArticleID 通过文章id获取所有的一级评论
func GetCommentListByArticleID(ctx context.Context, aHashId string) ([]*CommentItem, error) {
	// 查询所有顶级评论
	var err error
	var firstList []Comment
	//初始化cmtList
	var finalList []*CommentItem
	var cols = mydal.MyMongo.Cols
	cursor, err := cols.Comment.Find(ctx, bson.M{
		"article_id": aHashId,
		"degree":     1, // 查找顶级评论
	})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &firstList); err != nil {
		return nil, err
	}
	if len(firstList) != 0 {
		//获取所有的hashid
		var hashIDs []string
		for _, v := range firstList {
			hashIDs = append(hashIDs, v.HashID)
		}
		pipeline := []bson.M{
			{
				"$match": bson.M{"ancestor": bson.M{"$in": hashIDs}}, // 匹配祖先ID在给定列表中的记录
			},
			{
				"$group": bson.M{
					"_id":   "$ancestor",       // 按祖先ID分组
					"count": bson.M{"$sum": 1}, // 统计每组的数量
				},
			},
		}

		// 执行聚合查询
		cursor, err = cols.CmtClosure.Aggregate(ctx, pipeline)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		// map[评论id]回复总数
		sumMaps := make(map[string]int64)

		// 解析聚合结果
		for cursor.Next(ctx) {
			var item struct {
				ID    string `bson:"_id"`
				Count int64  `bson:"count"`
			}
			if err := cursor.Decode(&item); err != nil {
				return nil, err
			}
			sumMaps[item.ID] = item.Count // 将结果存入映射
		}
		//处理结果
		for _, v := range firstList {
			finalList = append(finalList, &CommentItem{
				ArticleID:  v.ArticleID,
				Content:    v.Content,
				Degree:     v.Degree,
				UserID:     v.UserID,
				HashID:     v.HashID,
				CreateTime: v.CreateTime,
				ChildNum:   sumMaps[v.HashID] - 1, //减掉1的原因：closure table(has)-->ancestor:hashid ,descendant:hashid
			})
		}
	}
	return finalList, nil
}

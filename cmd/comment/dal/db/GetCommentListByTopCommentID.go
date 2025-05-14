package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetCommentListByTopCommentID(ctx context.Context, aHashId string, cHashId string) (cmtList []*CommentItem, err error) {
	var cols = mydal.MyMongo.Cols
	//通过closure获取所有要请求的评论ID
	cursor, err := cols.CmtClosure.Find(ctx, &bson.M{"ancestor": cHashId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	//将结果推入cids数组
	var cids []string
	for cursor.Next(ctx) {
		item := new(CommentClosure)
		if err := cursor.Decode(item); err != nil {
			return nil, err
		}
		cids = append(cids, item.DescendantID)
	}
	//根据cids去请求评论列表
	var firstList []*Comment
	cursor, err = cols.Comment.Find(ctx, &bson.M{"hash_id": bson.M{"$in": cids}, "article_id": aHashId})
	if err != nil {
		return nil, err
	}
	//map[comment id]user id
	uMaps := make(map[string]string)
	for cursor.Next(ctx) {
		item := new(Comment)
		if err := cursor.Decode(item); err != nil {
			return nil, err
		}
		uMaps[item.HashID] = item.UserID
		if item.HashID != cHashId {
			firstList = append(firstList, item)
		}
	}
	for _, v := range firstList {
		cmtList = append(cmtList, &CommentItem{
			ArticleID:  v.ArticleID,
			Content:    v.Content,
			Degree:     v.Degree,
			UserID:     v.UserID,
			HashID:     v.HashID,
			ParentUID:  uMaps[v.ParentID],
			CreateTime: v.CreateTime,
		})
	}
	return cmtList, nil
}

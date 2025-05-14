package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetCommentByCmtID GetCommentByTopCmtID 根据comment hashid 获取评论
func GetCommentByCmtID(ctx context.Context, cHashId string) (cmt *Comment, err error) {
	cmt = new(Comment)
	//挑选某一个顶级评论
	err = mydal.MyMongo.Cols.Comment.FindOne(ctx, bson.M{
		"hash_id": cHashId,
	}).Decode(cmt)
	if err != nil {
		return nil, err
	}
	return cmt, nil
}

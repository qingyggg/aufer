package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"github.com/qingyggg/aufer/pkg/errno"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// AddNewComment add a comment
func AddNewComment(ctx context.Context, comment *Comment) error {
	var err error
	var cols = mydal.MyMongo.Cols
	//初始化创建时间
	comment.CreateTime = bson.NewDateTimeFromTime(time.Now())
	//顶级评论
	if comment.Degree == 1 {
		comment.ParentID = comment.HashID
		rootCmtL := comment
		rootClosure := &CommentClosure{AncestorID: comment.HashID, DescendantID: comment.HashID, Depth: 0, ArticleID: comment.ArticleID}
		_, err = cols.Comment.InsertOne(ctx, rootCmtL)
		_, err = cols.CmtClosure.InsertOne(ctx, rootClosure)
		if err != nil {
			return err
		}
	} else if comment.Degree == 2 { //子集评论
		//这里判断parent id 是否存在
		count, err := cols.Comment.CountDocuments(ctx, bson.M{"hash_id": comment.ParentID})
		if err != nil {
			return err
		}
		if count == 0 {
			return errno.ServiceErr.WithMessage("该评论的parentID不存在")
		}
		_, err = cols.Comment.InsertOne(ctx, comment) //插入评论
		if err != nil {
			return err
		}
		cursor, err := cols.CmtClosure.Find(ctx, bson.M{"descendant": comment.ParentID})
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		// 遍历祖先评论并生成新的闭包表记录
		var closures []interface{}
		for cursor.Next(ctx) {
			closure := new(CommentClosure)
			if err := cursor.Decode(&closure); err != nil {
				return err
			}
			newClosure := CommentClosure{
				DescendantID: comment.HashID,
				AncestorID:   closure.AncestorID,
				Depth:        closure.Depth + 1,
				ArticleID:    comment.ArticleID,
			}
			closures = append(closures, newClosure)
		}
		//增添自己的闭包表
		closures = append(closures, CommentClosure{
			AncestorID:   comment.HashID,
			Depth:        0,
			DescendantID: comment.HashID,
			ArticleID:    comment.ArticleID,
		})
		// 插入所有闭包表记录
		if len(closures) > 0 {
			_, err = cols.CmtClosure.InsertMany(ctx, closures)
			if err != nil {
				return err
			}
		}
	}
	return err
}

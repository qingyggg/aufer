package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"sync"
)

func DelCommentByHashID(ctx context.Context, cHashId string, aHashId string) error {
	var childIds []string //该评论的子评论数组
	var wg sync.WaitGroup
	var cols = mydal.MyMongo.Cols
	cfrdb := mydal.MyRedis.RdbComment
	errChan := make(chan error, 4) // 处理错误的channel
	defer close(errChan)
	//查找该评论以及该评论的子评论的id
	cursor, err := cols.CmtClosure.Find(ctx, bson.M{"ancestor": cHashId})
	if err != nil {
		return err
	}
	curClo := new(CommentClosure)
	for cursor.Next(ctx) {
		err = cursor.Decode(curClo)
		if err != nil {
			return err
		}
		childIds = append(childIds, curClo.DescendantID)
	}
	//获取评论数量在删除之前
	cmtCount, err := cols.Comment.CountDocuments(ctx, bson.M{"article_id": aHashId})
	if err != nil {
		return err
	}
	wg.Add(4)
	//删掉评论以及子评论
	go func() {
		defer wg.Done()
		_, err := cols.Comment.DeleteMany(ctx, bson.M{"hash_id": bson.M{"$in": childIds}})
		if err != nil {
			errChan <- err
			return
		}
	}()
	// 删除closure依赖
	go func() {
		defer wg.Done()
		_, err := cols.Comment.DeleteMany(ctx, bson.M{
			"$or": []bson.M{
				{"ancestor": bson.M{"$in": childIds}},
				{"descendant": bson.M{"$in": childIds}},
			},
		})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//删除掉该评论的favorite以及该评论的子评论的favorite
	go func() {
		defer wg.Done()
		err := CmtFavoriteFlushByCmtId(cHashId)
		if err != nil {
			errChan <- err
			return
		}
	}()
	//redis comment count减小
	go func() {
		defer wg.Done()
		err := cfrdb.CommentCtAssign(aHashId, cmtCount-int64(len(childIds)))
		if err != nil {
			errChan <- err
			return
		}
	}()
	wg.Wait()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

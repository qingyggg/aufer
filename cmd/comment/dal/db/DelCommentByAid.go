package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"sync"
)

// DelCommentByAid 删除文章前，调用该函数
func DelCommentByAid(ctx context.Context, aHashId string) error {
	var wg sync.WaitGroup
	cmtRdb := mydal.MyRedis.RdbComment
	cols := mydal.MyMongo.Cols
	errChan := make(chan error, 4)
	wg.Add(4)
	defer close(errChan)
	//根据文章id删除评论
	go func() {
		defer wg.Done()
		_, err := cols.Comment.DeleteMany(ctx, bson.M{"article_id": aHashId})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//根据文章id删除评论closure
	go func() {
		defer wg.Done()
		_, err := cols.CmtClosure.DeleteMany(ctx, bson.M{"article_id": aHashId})
		if err != nil {
			errChan <- err
			return
		}
	}()
	//清除redis缓存
	go func() {
		defer wg.Done()
		err, exist := cmtRdb.CheckCommentCt(aHashId)
		if err != nil {
			errChan <- err
			return
		}
		if exist {
			err := cmtRdb.DelCommentCt(aHashId)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		//清除掉cmt favorite
		err := CmtFavoriteFlushByArticleId(aHashId)
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

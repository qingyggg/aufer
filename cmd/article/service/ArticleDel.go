package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/article"
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/biz/rpc"
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	comment_rpc "github.com/qingyggg/aufer/cmd/comment/rpc"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"sync"
)

func (s *ArticleService) ArticleDelete(req *article.DelRequest) (err error) {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	var wg sync.WaitGroup
	wg.Add(5)
	errChan := make(chan error, 5)
	go func() {
		defer wg.Done()
		//1.删除文章
		err = db.DeleteArticle(&orm_gen.Article{UserID: utils.ConvertStringHashToByte(req.Uid), HashID: utils.ConvertStringHashToByte(req.Aid)})
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		//2.删除与文章相关的评论
		err := comment_rpc.CommentDelAction(rpc.Clients.CommentClient, s.ctx, &comment.DelRequest{Aid: req.Aid, Uid: req.Uid})
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		//3.删除article favorite依赖以及article favorite count
		err := db.AFavoriteDeleteByAid(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		err = db.ViewCountDel(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
	}()
	//清理掉collect
	go func() {
		defer wg.Done()
		err := db.ACollectDelByAid(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/article"
	"github.com/qingyggg/aufer/biz/rpc"
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/cmd/article/pack"
	comment_rpc "github.com/qingyggg/aufer/cmd/comment/rpc"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/pkg/errno"
	"sync"
)

func (s *ArticleService) ArticleDetail(req *article.DetailRequest) (*article.Article, error) {
	aA := new(article.Article)
	aA.Info = new(article.ArticleInfo)
	//1.get article data from db
	if err := s.publishDetail(req.Aid, aA); err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errChan = make(chan error, 6)
	wg.Add(6)
	//1.get author info
	go func() {
		defer wg.Done()
		//获取作者信息
		curUser, err := user_rpc.QueryUser(rpc.Clients.UserClient, s.ctx, req.Uid, req.MyUid)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Author = curUser //作者信息
		mu.Unlock()
	}()
	//2. like count
	go func() {
		defer wg.Done()
		err, ct := db.AFavoriteCtGet(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Info.LikeCount = ct
		mu.Unlock()
	}()
	//3.comment count
	go func() {
		defer wg.Done()
		ct, err := comment_rpc.CommentCount(rpc.Clients.CommentClient, s.ctx, req.Aid)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Info.CommentCount = ct
		mu.Unlock()
	}()
	//collect count
	go func() {
		defer wg.Done()
		err, count := db.ACollectCtGet(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Info.CollectCount = count
		mu.Unlock()
	}()
	//view count
	go func() {
		defer wg.Done()
		err := db.SyncViewCountInCache(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
		err, count := db.ViewCountGet(req.Aid)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Info.ViewedCount = count
		mu.Unlock()
	}()
	//if token exists:
	//is favorite,is collect
	go func() {
		defer wg.Done()
		if req.MyUid != "guest" {
			var err error
			var colEx bool
			var faSig int32
			var faEx bool
			err, colEx = db.ACollectExist(req.Aid, req.Uid)
			if err != nil {
				errChan <- err
				return
			}
			err, faSig = db.AFavoriteExist(req.Aid, req.Uid)
			if err != nil {
				errChan <- err
				return
			}
			if faSig == 1 {
				faEx = true
			} else {
				faEx = false
			}
			aA.Info.IsFavorite = faEx
			aA.Info.IsCollect = colEx
		} else {
			aA.Info.IsFavorite = false
			aA.Info.IsCollect = false
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	return aA, nil
}

func (s *ArticleService) publishDetail(aHashId string, article *article.Article) error {
	//检查文章是否存在
	exist, err := db.CheckArticleExistByHashId(aHashId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	//获取文章信息
	aInfo, aContent, err := db.TakeArticle(aHashId)
	if err != nil {
		return err
	}
	article.Content = aContent
	pack.BaseInfoAssign(aInfo, article.Info)
	return nil
}

package service

import (
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/pkg/errno"
)

func (c *FavoriteService) ArticleFavoriteAction(req *article.FavoriteRequest) error {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	err, exSignal := db.AFavoriteExist(req.Aid, req.Uid)
	if err != nil {
		return err
	}
	if exSignal != 3 { //1 or 2
		if exSignal == req.Type {
			return nil
		}
		err = db.AFavoriteStatusFlush(req.Aid, req.Uid, exSignal)
		if err != nil {
			return err
		}
	}
	if req.Type == 1 || req.Type == 2 {
		err = db.AFavorite(req.Aid, req.Uid, req.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

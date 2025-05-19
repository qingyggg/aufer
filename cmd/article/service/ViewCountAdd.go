package service

import (
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/pkg/errno"
)

func (s *ArticleService) AddViewCount(req *article.Aid) error {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	err = db.SyncViewCountInCache(req.Aid)
	if err != nil {
		return err
	}
	err = db.ViewCountIncr(req.Aid)
	if err != nil {
		return err
	}
	return nil
}

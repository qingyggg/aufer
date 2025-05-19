package service

import (
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/pkg/errno"
)

func (c *CollectService) ACollectAction(req *article.CollectRequest) error {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	err, exist = db.ACollectExist(req.Aid, req.Uid)
	if err != nil {
		return err
	}
	if exist {
		if req.Type == 1 {
			return errno.CollectAlreadyExistErr
		} else {
			return db.ACollectDel(req.Aid, req.Uid)
		}

	} else {
		if req.Type == 2 {
			return errno.CollectIsNotExistErr
		} else {
			return db.ACollectAdd(req.Aid, req.Uid, "default")
		}
	}
}

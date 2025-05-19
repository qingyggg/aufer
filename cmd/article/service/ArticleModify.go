package service

import (
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"time"
)

func (s *ArticleService) ArticleModify(req *article.ModRequest) (err error) {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}

	_, err = db.ModifyArticle(&orm_gen.Article{
		UserID:       utils.ConvertStringHashToByte(req.Uid),
		HashID:       utils.ConvertStringHashToByte(req.Aid),
		Title:        req.Title,
		Note:         req.Note,
		CoverURL:     utils.UrlConvertReverse(req.CoverUrl),
		LastModified: time.Now(),
	}, req.Content)
	return err
}

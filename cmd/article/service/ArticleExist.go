package service

import (
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
)

func (s *ArticleService) ArticleExist(req *article.Aid) (bool, error) {
	exist, err := db.CheckArticleExistByHashId(req.Aid)
	if err != nil {
		return false, err
	}
	return exist, nil
}

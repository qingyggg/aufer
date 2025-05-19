package service

import (
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/utils"
	"time"
)

func (s *ArticleService) ArticleCreate(req *article.PublishRequest) (err error, aHashId string) {
	var coverUrl string
	if req.CoverUrl == "" {
		coverUrl = constants.DefaultBackground
	} else {
		coverUrl = utils.UrlConvertReverse(req.CoverUrl)
	}
	aHashId = utils.GetSHA256String(time.Now().String() + req.Uid)
	//1.数据库创建记录
	err = db.CreateArticle(&orm_gen.Article{
		UserID:       utils.ConvertStringHashToByte(req.Uid),
		Title:        req.Title,
		Note:         req.Note,
		CoverURL:     coverUrl,
		PublishTime:  time.Now(),
		LastModified: time.Now(),
		HashID:       utils.ConvertStringHashToByte(aHashId),
	}, req.Content)
	if err != nil {
		return err, ""
	}
	//2.初始化view count
	err = db.ViewCountInit(aHashId)
	if err != nil {
		return err, ""
	}
	return nil, aHashId
}

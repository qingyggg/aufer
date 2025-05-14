package pack

import (
	"github.com/qingyggg/aufer/biz/model/cmd/article"
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/pkg/utils"
)

func BaseInfoAssign(aInfo *orm_gen.Article, resInfo *article.ArticleInfo) {
	resInfo.Title = aInfo.Title
	resInfo.Note = aInfo.Note
	resInfo.CoverUrl = aInfo.CoverURL
	resInfo.CreatedDate = aInfo.PublishTime.Format("2006-01-02")
	resInfo.LastModifiedDate = aInfo.LastModified.Format("2006-01-02")
	resInfo.Aid = utils.ConvertByteHashToString(aInfo.HashID)
}

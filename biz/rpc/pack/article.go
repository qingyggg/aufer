package rpcpack

import (
	"github.com/qingyggg/aufer/biz/model/http/basic/article"
	articlecmd "github.com/qingyggg/aufer/kitex_gen/cmd/article"
)

func ConvertArticleInfo(a *articlecmd.ArticleInfo) *article.ArticleInfo {
	return &article.ArticleInfo{
		LikeCount:        a.LikeCount,
		CommentCount:     a.CommentCount,
		IsFavorite:       a.IsFavorite,
		CollectCount:     a.CollectCount,
		ViewedCount:      a.ViewedCount,
		IsCollect:        a.IsCollect,
		Note:             a.Note,
		Title:            a.Title,
		CoverUrl:         a.CoverUrl,
		Aid:              a.Aid,
		CreatedDate:      a.CreatedDate,
		LastModifiedDate: a.LastModifiedDate,
	}
}
func ConvertArticle(a *articlecmd.Article) *article.Article {
	return &article.Article{
		Info:    ConvertArticleInfo(a.Info),
		Author:  ConvertUser(a.Author),
		Content: a.Content,
	}
}
func ConvertArticleCard(a *articlecmd.ArticleCard) *article.ArticleCard {
	return &article.ArticleCard{
		Info:   ConvertArticleInfo(a.Info),
		Author: ConvertUserBase(a.Author),
	}
}
func ConvertArticleCards(as []*articlecmd.ArticleCard) (cards []*article.ArticleCard) {
	for _, v := range as {
		cards = append(cards, ConvertArticleCard(v))
	}
	return
}

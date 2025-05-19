package rpc

import (
	"context"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article"
	"github.com/qingyggg/aufer/kitex_gen/cmd/article/articlehandler"

	rpc_util "github.com/qingyggg/aufer/cmd"
)

func ArticleExist(ac articlehandler.Client, ctx context.Context, aid string) (bool, error) {
	res, err := ac.ArticleExist(ctx, &article.Aid{Aid: aid})
	if err != nil {
		return false, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return false, err
	}
	return res.Exist, nil
}
func ArticleFavorite(ac articlehandler.Client, ctx context.Context, req *article.FavoriteRequest) error {
	res, err := ac.ArticleFavorite(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func ArticleCollect(ac articlehandler.Client, ctx context.Context, req *article.CollectRequest) error {
	res, err := ac.ArticleCollect(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func ArticleViewCountAdd(ac articlehandler.Client, ctx context.Context, aid string) error {
	res, err := ac.ArticleViewCountAdd(ctx, &article.Aid{Aid: aid})
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func ArticleDetail(ac articlehandler.Client, ctx context.Context, req *article.DetailRequest) (*article.Article, error) {
	res, err := ac.ArticleDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.Article, nil
}
func ArticleList(ac articlehandler.Client, ctx context.Context, req *article.CardsRequest) ([]*article.ArticleCard, error) {
	res, err := ac.ArticleList(ctx, req)
	if err != nil {
		return nil, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
func ArticleDelAction(ac articlehandler.Client, ctx context.Context, req *article.DelRequest) error {
	res, err := ac.ArticleDelAction(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func ArticleModifyAction(ac articlehandler.Client, ctx context.Context, req *article.ModRequest) (string, error) {
	res, err := ac.ArticleModifyAction(ctx, req)
	if err != nil {
		return "", err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return "", err
	}
	return res.Aid, nil
}
func PublishAction(ac articlehandler.Client, ctx context.Context, req *article.PublishRequest) (string, error) {
	res, err := ac.PublishAction(ctx, req)
	if err != nil {
		return "", err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return "", err
	}
	return res.Aid, nil
}

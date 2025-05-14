package main

import (
	"context"
	"github.com/qingyggg/aufer/biz/model/cmd/article"
	"github.com/qingyggg/aufer/biz/model/cmd/common"
	rpcutil "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/cmd/article/service"
	"github.com/qingyggg/aufer/pkg/errno"
)

// ArticleHandlerImpl implements the last service interface defined in the IDL.
type ArticleHandlerImpl struct{}

// PublishAction implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) PublishAction(ctx context.Context, req *article.PublishRequest) (resp *article.PubOrModActionResponse, err error) {
	err, aid := service.NewArticleService(ctx).ArticleCreate(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Aid = aid
	return resp, nil
}

// ArticleModifyAction implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleModifyAction(ctx context.Context, req *article.ModRequest) (resp *article.PubOrModActionResponse, err error) {
	err = service.NewArticleService(ctx).ArticleModify(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// ArticleDelAction implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleDelAction(ctx context.Context, req *article.DelRequest) (resp *common.BaseResponse, err error) {
	err = service.NewArticleService(ctx).ArticleDelete(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// ArticleList implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleList(ctx context.Context, req *article.CardsRequest) (resp *article.CardsResponse, err error) {
	list, err := service.NewArticleService(ctx).ArticleList(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.List = list
	return resp, nil
}

// ArticleDetail implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleDetail(ctx context.Context, req *article.DetailRequest) (resp *article.ArticleResponse, err error) {
	detail, err := service.NewArticleService(ctx).ArticleDetail(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Article = detail
	return resp, nil
}

// ArticleViewCountAdd implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleViewCountAdd(ctx context.Context, req *article.Aid) (resp *common.BaseResponse, err error) {
	err = service.NewArticleService(ctx).AddViewCount(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// ArticleCollect implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleCollect(ctx context.Context, req *article.CollectRequest) (resp *common.BaseResponse, err error) {
	err = service.NewCollectService(ctx).ACollectAction(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// ArticleFavorite implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleFavorite(ctx context.Context, req *article.FavoriteRequest) (resp *common.BaseResponse, err error) {
	err = service.NewFavoriteService(ctx).ArticleFavoriteAction(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// ArticleExist implements the ArticleHandlerImpl interface.
func (s *ArticleHandlerImpl) ArticleExist(ctx context.Context, req *article.Aid) (resp *article.ArticleExistResponse, err error) {
	exist, err := service.NewArticleService(ctx).ArticleExist(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Exist = exist
	return resp, nil
}

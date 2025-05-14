package service

import "context"

type ArticleService struct {
	ctx context.Context
}

func NewArticleService(ctx context.Context) *ArticleService {
	return &ArticleService{
		ctx: ctx,
	}
}

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

type CollectService struct {
	ctx context.Context
}

func NewCollectService(ctx context.Context) *CollectService {
	return &CollectService{
		ctx: ctx,
	}
}

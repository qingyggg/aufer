package service

import "context"

type CmtFavoriteService struct {
	ctx context.Context
}

// NewCmtFavoriteService create comment service
func NewCmtFavoriteService(ctx context.Context) *CmtFavoriteService {
	return &CmtFavoriteService{ctx: ctx}
}

type CommentService struct {
	ctx context.Context
}

// NewCommentService create comment service
func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

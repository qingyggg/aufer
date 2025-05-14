package main

import (
	"context"
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/biz/model/cmd/common"
	rpcutil "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/cmd/comment/service"
	"github.com/qingyggg/aufer/pkg/errno"
)

// CommentHandlerImpl implements the last service interface defined in the IDL.
type CommentHandlerImpl struct{}

// CommentAction implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentAction(ctx context.Context, req *comment.CmtRequest) (resp *comment.CmtActionResponse, err error) {
	resp = new(comment.CmtActionResponse)
	err, cid := service.NewCommentService(ctx).AddNewCmt(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Cid = cid
	return resp, nil
}

// CommentDelAction implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentDelAction(ctx context.Context, req *comment.DelRequest) (resp *common.BaseResponse, err error) {
	resp = new(common.BaseResponse)
	err = service.NewCommentService(ctx).DelCmt(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// CommentList implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentList(ctx context.Context, req *comment.CardsRequest) (resp *comment.CardsResponse, err error) {
	resp = new(comment.CardsResponse)
	err, list := service.NewCommentService(ctx).GetCmtList(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.List = list
	return resp, nil
}

// CommentFavorite implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentFavorite(ctx context.Context, req *comment.FavoriteRequest) (resp *common.BaseResponse, err error) {
	resp = new(common.BaseResponse)
	err = service.NewCmtFavoriteService(ctx).CmtFavoriteAction(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// CommentCount implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentCount(ctx context.Context, req *comment.CmtCtRequest) (resp *comment.CmtCtResponse, err error) {
	resp = new(comment.CmtCtResponse)
	ct, err := service.NewCommentService(ctx).GetCmtCt(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Count = ct
	return resp, nil
}

// CommentCounts implements the CommentHandlerImpl interface.
func (s *CommentHandlerImpl) CommentCounts(ctx context.Context, req *comment.CmtsCtRequest) (resp *comment.CmtCtMapResponse, err error) {
	resp = new(comment.CmtCtMapResponse)
	ctMap, err := service.NewCommentService(ctx).GetCmtCtMap(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.CtMap = ctMap
	return resp, nil
}

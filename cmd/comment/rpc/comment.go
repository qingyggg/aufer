package rpc

import (
	"context"
	rpc_util "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/kitex_gen/cmd/comment"
	"github.com/qingyggg/aufer/kitex_gen/cmd/comment/commenthandler"
)

func CommentAction(cc commenthandler.Client, ctx context.Context, req *comment.CmtRequest) (string, error) {
	res, err := cc.CommentAction(ctx, req)
	if err != nil {
		return "", err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return "", err
	}
	return res.Cid, nil
}
func CommentDelAction(cc commenthandler.Client, ctx context.Context, req *comment.DelRequest) error {
	res, err := cc.CommentDelAction(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func CommentList(cc commenthandler.Client, ctx context.Context, req *comment.CardsRequest) ([]*comment.Comment, error) {
	res, err := cc.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
func CommentCount(cc commenthandler.Client, ctx context.Context, aid string) (int64, error) {
	res, err := cc.CommentCount(ctx, &comment.CmtCtRequest{Aid: aid})
	if err != nil {
		return 0, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return 0, err
	}
	return res.Count, nil
}
func CommentCounts(cc commenthandler.Client, ctx context.Context, aids []string) (map[string]int64, error) {
	res, err := cc.CommentCounts(ctx, &comment.CmtsCtRequest{Aids: aids})
	if err != nil {
		return nil, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.CtMap, nil
}
func CommentFavorite(cc commenthandler.Client, ctx context.Context, req *comment.FavoriteRequest) error {
	res, err := cc.CommentFavorite(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}

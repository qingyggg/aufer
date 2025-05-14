package comment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	comment_cmd "github.com/qingyggg/aufer/biz/model/cmd/comment"
	common "github.com/qingyggg/aufer/biz/model/cmd/common"
	comment "github.com/qingyggg/aufer/biz/model/http/interact/comment"
	favorite "github.com/qingyggg/aufer/biz/model/http/interact/favorite"
	"github.com/qingyggg/aufer/biz/rpc"
	comment_rpc "github.com/qingyggg/aufer/cmd/comment/rpc"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
)

// CommentAction 创建新评论
// @Summary 创建新评论
// @Description 为特定文章或帖子创建新评论
// @Tags 评论
// @Accept json
// @Produce json
// @Param request body comment.CmtRequest true "评论创建信息"
// @Success 200 {object} comment_cmd.CmtActionResponse "成功响应，包含评论ID"
// @Failure 400 {object} common.BaseResponse "请求错误"
// @Failure 401 {object} common.BaseResponse "未授权"
// @Failure 500 {object} common.BaseResponse "服务器内部错误"
// @Router /publish/comment [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(comment.CmtRequest)
	resp := new(comment_cmd.CmtActionResponse)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	cid, err := comment_rpc.CommentAction(rpc.Clients.CommentClient, ctx, &comment_cmd.CmtRequest{
		Pid:     req.Pid,
		Aid:     req.Aid,
		Content: req.Content,
		Degree:  req.Degree,
		Uid:     utils.GetUid(c, ctx),
	})
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	resp.Base = utils.BuildBaseResp(nil)
	resp.Cid = cid
	c.JSON(consts.StatusOK, resp)
}

// CommentDelAction 删除评论
// @Summary 删除评论
// @Description 通过ID删除已存在的评论
// @Tags 评论
// @Accept json
// @Produce json
// @Param request body comment.DelRequest true "评论删除信息"
// @Success 200 {object} common.BaseResponse "成功响应"
// @Failure 400 {object} common.BaseResponse "请求错误"
// @Failure 401 {object} common.BaseResponse "未授权"
// @Failure 403 {object} common.BaseResponse "禁止访问"
// @Failure 404 {object} common.BaseResponse "评论未找到"
// @Failure 500 {object} common.BaseResponse "服务器内部错误"
// @Router /publish/comment [DELETE]
func CommentDelAction(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(comment.DelRequest)
	resp := new(common.BaseResponse)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	err = comment_rpc.CommentDelAction(rpc.Clients.CommentClient, ctx, &comment_cmd.DelRequest{
		Cid:  req.Cid,
		Aid:  req.Aid,
		Type: req.Type,
		Uid:  utils.GetUid(c, ctx),
	})
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	resp = utils.BaseResp(errno.Success)
	c.JSON(consts.StatusOK, resp)
}

// CommentList 获取评论列表
// @Summary 获取评论列表
// @Description 返回特定文章或帖子的评论列表
// @Tags 评论
// @Accept json
// @Produce json
// @Param request body comment.CardsRequest true "评论列表参数"
// @Success 200 {object} comment_cmd.CardsResponse "成功响应，包含评论列表"
// @Failure 400 {object} common.BaseResponse "请求错误"
// @Failure 401 {object} common.BaseResponse "未授权"
// @Failure 404 {object} common.BaseResponse "评论未找到"
// @Failure 500 {object} common.BaseResponse "服务器内部错误"
// @Router /publish/comment/list [POST]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(comment.CardsRequest)
	resp := new(comment_cmd.CardsResponse)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	cmts, err := comment_rpc.CommentList(rpc.Clients.CommentClient, ctx, &comment_cmd.CardsRequest{
		Aid:    req.Aid,
		Cid:    req.Cid,
		Degree: req.Degree,
		MyUid:  req.MyUid,
	})
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	resp.Base = utils.BuildBaseResp(nil)
	resp.List = cmts
	c.JSON(consts.StatusOK, resp)
}

// CommentFavorite 点赞评论
// @Summary 点赞评论
// @Description 点赞评论
// @Tags 评论
// @Accept json
// @Produce json
// @Param request body favorite.FavoriteRequest true "评论收藏信息"
// @Success 200 {object} common.BaseResponse "成功响应"
// @Failure 400 {object} common.BaseResponse "请求错误"
// @Failure 401 {object} common.BaseResponse "未授权"
// @Failure 404 {object} common.BaseResponse "评论未找到"
// @Failure 500 {object} common.BaseResponse "服务器内部错误"
// @Router /publish/comment/favorite [POST]
func CommentFavorite(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(favorite.FavoriteRequest)
	resp := new(common.BaseResponse)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	err = comment_rpc.CommentFavorite(rpc.Clients.CommentClient, ctx, &comment_cmd.FavoriteRequest{
		Cid:  req.Cid,
		Type: req.Type,
		Uid:  utils.GetUid(c, ctx),
		Aid:  req.Aid,
	})
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	resp = utils.BaseResp(errno.Success)
	c.JSON(consts.StatusOK, resp)
}

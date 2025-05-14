package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	"github.com/qingyggg/aufer/pkg/errno"
)

func (c *CommentService) DelCmt(req *comment.DelRequest) error {
	exist, err := db.CheckCmtExistById(c.ctx, req.Cid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistErr
	}
	cmt, err := db.GetCommentByCmtID(c.ctx, req.Cid)
	if err != nil {
		return err
	}
	if cmt.UserID != req.Uid {
		return errno.ServiceErr.WithMessage("当前用户没有权利删除该评论")
	}
	err = db.DelCommentByHashID(c.ctx, req.Cid, req.Aid)
	return err
}

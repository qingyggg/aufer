package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/biz/rpc"
	article_rpc "github.com/qingyggg/aufer/cmd/article/rpc"
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"math/rand"
	"strconv"
	"time"
)

func (c *CommentService) AddNewCmt(req *comment.CmtRequest) (error, string) {
	if req.Content == "" {
		return errno.ParamErr.WithMessage("评论内容不可以为空"), ""
	}
	//检查文章是否存在
	aExist, err := article_rpc.ArticleExist(rpc.Clients.ArticleClient, c.ctx, req.Aid)

	if err != nil {
		return err, ""
	}
	if !aExist {
		return errno.ArticleIsNotExistErr, "" //文章不存在
	}
	//检查父评论是否存在
	if req.Degree == 2 {
		exist, err := db.CheckCmtExistById(c.ctx, req.Pid)
		if err != nil {
			return err, ""
		}
		if !exist {
			return errno.CommentIsNotExistErr.WithMessage("您回复的评论不存在"), ""
		}
	}
	cHashId := utils.GetSHA256String(req.Aid + req.Uid + strconv.FormatInt(int64(rand.Int()), 10) + time.Now().String())
	err = db.AddNewComment(c.ctx, &db.Comment{
		ArticleID: req.Aid,
		UserID:    req.Uid,
		Content:   req.Content,
		ParentID:  req.Pid,
		Degree:    int8(req.Degree),
		HashID:    cHashId,
	})
	if err != nil {
		return err, ""
	}
	return nil, cHashId
}

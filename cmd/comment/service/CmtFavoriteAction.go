package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	"github.com/qingyggg/aufer/pkg/errno"
)

func (c *CmtFavoriteService) CmtFavoriteAction(req *comment.FavoriteRequest) error {
	exist, err := db.CheckCmtExistById(c.ctx, req.Cid)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistErr
	}
	err, exist = db.CmtFavorieExist(req.Uid, req.Cid)
	if err != nil {
		return err
	}
	if exist {
		err, status := db.CmtFavoriteStatus(req.Uid, req.Cid)
		if err != nil {
			return err
		}
		if status == req.Type {
			return nil
		} else {
			err = db.CmtFavoriteStatusFlush(req.Uid, req.Cid)
			if err != nil {
				return err
			}
		}
	}
	if req.Type == 1 || req.Type == 2 {
		err = db.CmtFavoriteAction(req.Uid, req.Cid, req.Aid, req.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

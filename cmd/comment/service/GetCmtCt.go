package service

import (
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/comment"
)

func (c *CommentService) GetCmtCt(req *comment.CmtCtRequest) (int64, error) {
	err, ct := db.GetCmtCountByAid(c.ctx, req.Aid)
	if err != nil {
		return 0, err
	}
	return ct, nil
}

package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/comment"
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
)

func (c *CommentService) GetCmtCtMap(req *comment.CmtsCtRequest) (map[string]int64, error) {
	err, ctMap := db.GetCmtCtByAids(c.ctx, req.Aids)
	if err != nil {
		return nil, err
	}
	return ctMap, nil
}

package service

import (
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/comment"
)

func (c *CommentService) GetCmtCtMap(req *comment.CmtsCtRequest) (map[string]int64, error) {
	err, ctMap := db.GetCmtCtByAids(c.ctx, req.Aids)
	if err != nil {
		return nil, err
	}
	return ctMap, nil
}

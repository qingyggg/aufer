package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/relation"
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
)

func (s *RelationService) QueryIsFollow(req *relation.IsFollowRequest) (bool, error) {
	is, err := db.QueryFollowExist(req.Uid, req.MyUid)
	if err != nil {
		return false, err
	}
	return is, nil
}

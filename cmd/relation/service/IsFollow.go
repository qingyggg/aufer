package service

import (
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
)

func (s *RelationService) QueryIsFollow(req *relation.IsFollowRequest) (bool, error) {
	is, err := db.QueryFollowExist(req.Uid, req.MyUid)
	if err != nil {
		return false, err
	}
	return is, nil
}

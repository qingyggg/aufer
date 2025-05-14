package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/relation"
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
)

func (s *RelationService) QueryFollowInfo(req *relation.UserFollowInfoRequest) (int64, int64, error) {
	fct, err := db.GetFollowCount(req.Uid)
	if err != nil {
		return 0, 0, err
	}
	fcter, err := db.GetFollowerCount(req.Uid)
	if err != nil {
		return 0, 0, err
	}
	return fct, fcter, nil
}

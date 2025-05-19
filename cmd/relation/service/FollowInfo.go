package service

import (
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
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

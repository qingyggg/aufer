package service

import (
	"github.com/qingyggg/aufer/biz/rpc"
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
	userrpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func (s *RelationService) GetRelationList(req *relation.ListRequest) ([]*user.UserBase, error) {
	var err error
	uid := req.Uid
	var uids []string
	var ubList []*user.UserBase
	if req.Type == FOLLOWING {
		uids, err = db.GetFollowIdList(uid)
	} else {
		uids, err = db.GetFollowerIdList(uid)
	}
	if err != nil {
		return nil, err
	}
	ubMap, err := userrpc.QueryUserBases(rpc.Clients.UserClient, s.ctx, uids)
	if err != nil {
		return nil, err
	}
	for _, v := range ubMap {
		ubList = append(ubList, v)
	}
	return ubList, nil
}

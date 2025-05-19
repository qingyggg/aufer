package service

import (
	"context"
	"github.com/qingyggg/aufer/biz/rpc"
	relation_rpc "github.com/qingyggg/aufer/cmd/relation/rpc"
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/cmd/user/pack"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func (s *UserService) User(req *user.UserRequest) (*user.User, error) {
	//1.query user,my
	//2.query user ,guest
	//3.my
	var reqType uint8
	queryUid := req.Uid
	myUid := req.MyUid
	if queryUid != "current" {
		if myUid == "guest" {
			reqType = 2
		} else {
			reqType = 1
		}
	} else {
		reqType = 3
	}

	return getUser(s.ctx, queryUid, myUid, reqType)
}

func getUser(ctx context.Context, queryUid string, myUid string, reqType uint8) (*user.User, error) {
	var err error
	var base *user.UserBase
	var info = new(user.UserInfo)
	var qid string
	if reqType == 3 {
		qid = myUid
	} else {
		qid = queryUid
	}
	ubo, err := db.QueryUserByHashId(qid)
	if err != nil {
		return nil, err
	}
	base = pack.UserAssign(ubo)
	WorkCount, err := db.GetWorkCount(qid)
	if err != nil {
		return nil, err
	} else {
		info.WorkCount = WorkCount
	}
	fCt, ferCt, err := relation_rpc.UserFollowInfo(rpc.Clients.RelationClient, ctx, queryUid)
	if err != nil {
		return nil, err
	}
	info.FollowCount = int64(fCt)
	info.FollowerCount = int64(ferCt)
	if reqType == 1 {
		isFollow, err := relation_rpc.IsFollow(rpc.Clients.RelationClient, ctx, &relation.IsFollowRequest{Uid: queryUid, MyUid: myUid})
		if err != nil {
			return nil, err
		} else {
			info.IsFollow = isFollow
		}
	} else {
		info.IsFollow = false
	}

	return &user.User{Info: info, Base: base}, nil
}

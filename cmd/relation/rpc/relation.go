package rpc

import (
	"context"
	rpc_util "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation/relationhandler"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func Follow(rc relationhandler.Client, ctx context.Context, req *relation.FollowRequest) error {
	res, err := rc.Follow(ctx, req)
	if err != nil {
		return err
	}
	err = rpc_util.CheckBaseResp(res)
	if err != nil {
		return err
	}
	return nil
}
func FollowList(rc relationhandler.Client, ctx context.Context, req *relation.ListRequest) ([]*user.UserBase, error) {
	res, err := rc.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
func IsFollow(rc relationhandler.Client, ctx context.Context, req *relation.IsFollowRequest) (bool, error) {
	res, err := rc.IsFollow(ctx, req)
	if err != nil {
		return false, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return false, err
	}
	return res.IsFollow, nil
}
func UserFollowInfo(rc relationhandler.Client, ctx context.Context, uid string) (followingCt uint64, followerCt uint64, err error) {
	res, err := rc.UserFollowInfo(ctx, &relation.UserFollowInfoRequest{Uid: uid})
	if err != nil {
		return 0, 0, err
	}
	err = rpc_util.CheckBaseResp(res.Base)
	if err != nil {
		return 0, 0, err
	}
	return res.FollowingCount, res.FollowerCount, nil
}

package main

import (
	"context"
	rpcutil "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/cmd/relation/service"
	"github.com/qingyggg/aufer/kitex_gen/cmd/common"
	"github.com/qingyggg/aufer/kitex_gen/cmd/relation"
	"github.com/qingyggg/aufer/pkg/errno"
)

// RelationHandlerImpl implements the last service interface defined in the IDL.
type RelationHandlerImpl struct{}

// Follow implements the RelationHandlerImpl interface.
func (s *RelationHandlerImpl) Follow(ctx context.Context, req *relation.FollowRequest) (resp *common.BaseResponse, err error) {
	err = service.NewRelationService(ctx).FollowAction(req)
	if err != nil {
		resp = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp = rpcutil.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowList implements the RelationHandlerImpl interface.
func (s *RelationHandlerImpl) FollowList(ctx context.Context, req *relation.ListRequest) (resp *relation.ListResponse, err error) {
	resp = new(relation.ListResponse)
	ubases, err := service.NewRelationService(ctx).GetRelationList(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.List = ubases
	return resp, nil
}

// IsFollow implements the RelationHandlerImpl interface.
func (s *RelationHandlerImpl) IsFollow(ctx context.Context, req *relation.IsFollowRequest) (resp *relation.IsFollowResponse, err error) {
	resp = new(relation.IsFollowResponse)
	is, err := service.NewRelationService(ctx).QueryIsFollow(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.IsFollow = is
	return resp, nil
}

// UserFollowInfo implements the RelationHandlerImpl interface.
func (s *RelationHandlerImpl) UserFollowInfo(ctx context.Context, req *relation.UserFollowInfoRequest) (resp *relation.UserFollowInfoResponse, err error) {
	resp = new(relation.UserFollowInfoResponse)
	fct, fcter, err := service.NewRelationService(ctx).QueryFollowInfo(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.FollowingCount = uint64(fct)
	resp.FollowerCount = uint64(fcter)
	return resp, nil
}

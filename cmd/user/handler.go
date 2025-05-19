package main

import (
	"context"
	rpcutil "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/cmd/user/service"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
	"github.com/qingyggg/aufer/pkg/errno"
)

// UserHandlerImpl implements the last service interface defined in the IDL.
type UserHandlerImpl struct{}

// User implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) User(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp = new(user.UserResponse)
	u, err := service.NewUserService(ctx).User(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.User = u
	return resp, nil
}

// UserRegister implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserRegister(ctx context.Context, req *user.LoginOrRegRequest) (resp *user.UserActionResponse, err error) {
	resp = new(user.UserActionResponse)
	puid, uid, err := service.NewUserService(ctx).UserRegister(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Uid = uid
	resp.Puid = puid
	return resp, nil
}

// UserLogin implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserLogin(ctx context.Context, req *user.LoginOrRegRequest) (resp *user.UserActionResponse, err error) {
	resp = new(user.UserActionResponse)
	puid, uid, err := service.NewUserService(ctx).UserLogin(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Uid = uid
	resp.Puid = puid
	return resp, nil
}

// UserPwdModify implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserPwdModify(ctx context.Context, req *user.PwdModRequest) (resp *user.UserActionResponse, err error) {
	resp = new(user.UserActionResponse)
	puid, uid, err := service.NewUserService(ctx).PwdModify(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Uid = uid
	resp.Puid = puid
	return resp, nil
}

// UserProfileModify implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserProfileModify(ctx context.Context, req *user.ProfileModRequest) (resp *user.UserActionResponse, err error) {
	resp = new(user.UserActionResponse)
	err, uid := service.NewUserService(ctx).UserProfileModify(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Uid = uid
	return resp, nil
}

// QueryUserBases implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) QueryUserBases(ctx context.Context, req *user.UserBasesRequest) (resp *user.UserBasesResponse, err error) {
	resp = new(user.UserBasesResponse)
	userMap, err := service.NewUserService(ctx).QueryUserBases(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.UserMap = userMap
	return resp, nil
}

// QueryUserBase implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) QueryUserBase(ctx context.Context, req *user.UserBaseRequest) (resp *user.UserBaseResponse, err error) {
	resp = new(user.UserBaseResponse)
	u, err := service.NewUserService(ctx).QueryUserBase(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.User = u
	return resp, nil
}

// UserExist implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserExist(ctx context.Context, req *user.UserExistRequest) (resp *user.UserExistResponse, err error) {
	resp = new(user.UserExistResponse)
	ex, err := service.NewUserService(ctx).QueryUserExist(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.Exist = ex
	return resp, nil
}

// QueryUserBaseByPuid implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) QueryUserBaseByPuid(ctx context.Context, req *user.UserBaseRequestByPuid) (resp *user.UserBaseResponse, err error) {
	resp = new(user.UserBaseResponse)
	u, err := service.NewUserService(ctx).QueryUserBaseByPuid(req)
	if err != nil {
		resp.Base = rpcutil.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = rpcutil.BuildBaseResp(errno.Success)
	resp.User = u
	return resp, nil
}

package rpc

import (
	"context"
	rpcutil "github.com/qingyggg/aufer/cmd"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user/userhandler"
)

func CheckUserExist(uc userhandler.Client, ctx context.Context, uid string) (bool, error) {
	//1.check user
	res, err := uc.UserExist(ctx, &user.UserExistRequest{Uid: uid})
	if err != nil {
		return false, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return false, err
	}
	if res.Exist != true {
		return false, nil
	}
	return true, nil
}

func QueryUserBase(uc userhandler.Client, ctx context.Context, uid string) (*user.UserBase, error) {
	res, err := uc.QueryUserBase(ctx, &user.UserBaseRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}
func QueryUserBaseByPuid(uc userhandler.Client, ctx context.Context, puid int64) (*user.UserBase, error) {
	res, err := uc.QueryUserBaseByPuid(ctx, &user.UserBaseRequestByPuid{Puid: puid})
	if err != nil {
		return nil, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}
func QueryUserBases(uc userhandler.Client, ctx context.Context, uids []string) (map[string]*user.UserBase, error) {
	res, err := uc.QueryUserBases(ctx, &user.UserBasesRequest{Uids: uids})
	if err != nil {
		return nil, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.UserMap, nil
}

func QueryUser(uc userhandler.Client, ctx context.Context, uid string, myUid string) (*user.User, error) {
	res, err := uc.User(ctx, &user.UserRequest{Uid: uid, MyUid: myUid})
	if err != nil {
		return nil, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}

func UserRegister(uc userhandler.Client, ctx context.Context, email string, pwd string) (string, int64, error) {
	res, err := uc.UserRegister(ctx, &user.LoginOrRegRequest{Email: email, Password: pwd})
	if err != nil {
		return "", 0, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return "", 0, err
	}
	return res.Uid, res.Puid, nil
}
func UserLogin(uc userhandler.Client, ctx context.Context, email string, pwd string) (string, int64, error) {
	res, err := uc.UserLogin(ctx, &user.LoginOrRegRequest{Email: email, Password: pwd})
	if err != nil {
		return "", 0, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return "", 0, err
	}
	return res.Uid, res.Puid, nil
}
func UserPwdModify(uc userhandler.Client, ctx context.Context, req *user.PwdModRequest) (string, int64, error) {
	res, err := uc.UserPwdModify(ctx, req)
	if err != nil {
		return "", 0, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return "", 0, err
	}
	return res.Uid, res.Puid, nil
}
func UserProfileModify(uc userhandler.Client, ctx context.Context, req *user.ProfileModRequest) (string, int64, error) {
	res, err := uc.UserProfileModify(ctx, req)
	if err != nil {
		return "", 0, err
	}
	err = rpcutil.CheckBaseResp(res.Base)
	if err != nil {
		return "", 0, err
	}
	return res.Uid, res.Puid, nil
}

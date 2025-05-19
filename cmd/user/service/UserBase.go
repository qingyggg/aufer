package service

import (
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/cmd/user/pack"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func (s *UserService) QueryUserBase(req *user.UserBaseRequest) (*user.UserBase, error) {
	u, err := db.QueryUserByHashId(req.Uid)
	if err != nil {
		return nil, err
	}
	return pack.UserAssign(u), nil
}

func (s *UserService) QueryUserBaseByPuid(req *user.UserBaseRequestByPuid) (*user.UserBase, error) {
	u, err := db.QueryUserById(req.Puid)
	if err != nil {
		return nil, err
	}
	return pack.UserAssign(u), nil
}

// QueryUserBases 文章列表，评论列表上的人物信息
func (s *UserService) QueryUserBases(req *user.UserBasesRequest) (map[string]*user.UserBase, error) {
	baseMap := make(map[string]*user.UserBase)
	uMap, err := db.QueryUserByHashIds(req.Uids)
	if err != nil {
		return nil, err
	}
	for k, v := range uMap {
		baseMap[k] = pack.UserAssign(v)
	}
	return baseMap, nil
}

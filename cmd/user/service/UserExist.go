package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/cmd/user/dal/db"
)

func (s *UserService) QueryUserExist(req *user.UserExistRequest) (bool, error) {
	ex, err := db.CheckUserExistByHashId(req.Uid)
	if err != nil {
		return false, err
	}
	return ex, nil
}

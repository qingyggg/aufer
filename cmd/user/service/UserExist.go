package service

import (
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func (s *UserService) QueryUserExist(req *user.UserExistRequest) (bool, error) {
	ex, err := db.CheckUserExistByHashId(req.Uid)
	if err != nil {
		return false, err
	}
	return ex, nil
}

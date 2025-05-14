package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/cmd/user/dal/db"
)

func (s *UserService) UserLogin(req *user.LoginOrRegRequest) (puid int64, uid string, err error) {
	//puid===>user table primary key
	puid, uid, err = db.VerifyUser(req.Email, req.Password)
	if err != nil {
		return 0, "", err
	}
	return puid, uid, nil
}

package service

import (
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func (s *UserService) UserLogin(req *user.LoginOrRegRequest) (puid int64, uid string, err error) {
	//puid===>user table primary key
	puid, uid, err = db.VerifyUser(req.Email, req.Password)
	if err != nil {
		return 0, "", err
	}
	return puid, uid, nil
}

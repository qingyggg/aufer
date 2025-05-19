package service

import (
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
	"github.com/qingyggg/aufer/pkg/utils"
)

func (s *UserService) PwdModify(req *user.PwdModRequest) (puid int64, uid string, err error) {
	puid, uid, err = db.VerifyUser(req.Email, req.OldPassword) //验证用户是否存在，旧密码是否正确
	if err != nil {
		return 0, "", err
	}
	crptedPwd, err := utils.Crypt(req.NewPassword)
	if err != nil {
		return 0, "", err
	}
	err = db.UserPwdModify(puid, crptedPwd)
	if err != nil {
		return 0, "", err
	}
	return puid, uid, nil
}

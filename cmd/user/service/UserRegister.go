package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"time"
)

func (s *UserService) UserRegister(req *user.LoginOrRegRequest) (int64, string, error) {
	//检测用户是否存在
	isExist, err := db.CheckUserExistByEmail(req.Email)
	uid := utils.GetSHA256String(req.Email + time.Now().String())[:16] //截取哈希值的前16为作为用户的hashId
	if err != nil {
		return 0, "", err
	}
	if isExist {
		return 0, "", errno.UserAlreadyExistErr
	}
	passWord, err := utils.Crypt(req.Password)
	u := &orm_gen.User{
		Email:           req.Email,
		Password:        passWord,
		HashID:          utils.ConvertStringHashToByte(uid),
		Avatar:          constants.DefaultAva,
		BackgroundImage: constants.DefaultBackground,
		Signature:       constants.DefaultSign,
		UserName:        constants.DefaultUserName,
	}
	err = db.CreateUser(u)
	if err != nil {
		return 0, "", err
	}

	return u.ID, uid, nil
}
